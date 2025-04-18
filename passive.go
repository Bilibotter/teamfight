package tft

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type stack struct {
	attrs
	count    int
	maxStack int  // 为0表示无限叠加
	reset    bool // 是否在达到最大值时清零
}

/*
Freq: 到达一定频次再执行call。用于周期触发被动与可叠加被动。
once: 1: 表示被动只触发一次，2: 表示已触发
Left：once=1，表示被动触发的阈值; once=0，表示时间段被动的左边界，包含边界值。
Right：表示时间段被动的右边界，不包含边界值。
Trigger：表示触发call的事件类型。

	频次+可叠加被动支持触发类型：TimeGo、Cast、Attack
	时间段被动支持触发类型：TimeGo
*/
type passive struct {
	attrs
	name    string
	g       *ground
	overlay *stack
	Trigger action
	Freq    int             // 匹配事件全部事件，每多少次事件触发一次
	Count   int             // 与freq结合使用。
	Left    int             // 匹配事件timeGoA，界定的数值/左边界值
	Right   int             // 匹配事件timeGoA，右边界值，不包含右边界
	call    func(g *ground) // 匹配到事件时调用
	once    int             // once:0表示无限次,1表示只触发一次,2表示已触发
}

func minionPassive(trigger action, remain int, a ...*attrs) *passive {
	p := &passive{
		Trigger: trigger,
		call:    addMinion(remain, TimeGoA, a...),
	}
	return p
}

func buffPassive(trigger action, remain int, a ...*attrs) *passive {
	p := &passive{
		Trigger: trigger,
		call:    addBuff(remain, a...),
	}
	return p
}

func stackPassive(trigger action, freq int, a ...*attrs) *passive {
	return stackPassive0(trigger, 0, freq, a...)
}

func stackPassive0(trigger action, limit, freq int, a ...*attrs) *passive {
	attr := attrs{}
	for _, add := range a {
		attr.Add(add)
	}
	p := &passive{
		attrs:   attrs{}, // 修复叠加被动天生自带一层
		overlay: &stack{attrs: attr, maxStack: limit},
		Trigger: trigger,
		Freq:    freq,
	}
	return p
}

func attrPassive(trigger action, a ...*attrs) *passive {
	attr := attrs{}
	for _, add := range a {
		attr.Add(add)
	}
	p := &passive{
		attrs:   attr,
		Trigger: trigger,
	}
	return p
}

func (g *ground) StackPassive(trigger action, a ...*attrs) *passive {
	p := stackPassive(trigger, 1, a...)
	g.addPassive(p)
	return p
}

func (g *ground) StackPassive0(trigger action, limit, freq int, a ...*attrs) *passive {
	p := stackPassive0(trigger, limit, freq, a...)
	g.addPassive(p)
	return p
}

// 只触发一次的被动
func (g *ground) OncePassive(trigger action, freq int, a ...*attrs) *passive {
	// 蓝发只会触发一次有时限的加速，因此用buff而不是attr实现
	p := buffPassive(trigger, 30, a...)
	p.Freq = freq
	p.once = 1
	g.addPassive(p)
	return p
}

// 类似无人机
func (g *ground) DmgPassive(trigger action, freq, dmg int) *passive {
	p := &passive{
		Trigger: trigger,
		Freq:    freq,
		call: func(g *ground) {
			g.recordDmg(float64(dmg), fromPassive)
		},
	}
	g.addPassive(p)
	return p
}

func (g *ground) BuffPassive(trigger action, freq, remain int, a ...*attrs) *passive {
	p := buffPassive(trigger, remain, a...)
	p.Freq = freq
	g.addPassive(p)
	return p
}

// 泽丽召唤镜像，月男召唤炮台
func (g *ground) Minion(trigger action, freq, remain int, a ...*attrs) *passive {
	p := minionPassive(trigger, remain, a...)
	p.Freq = freq
	g.addPassive(p)
	return p
}

func (p *passive) key() string {
	segment1, _ := json.Marshal(p.attrs)
	segment2, _ := json.Marshal(p.Trigger)
	segment3, _ := json.Marshal((p.Freq*1 + 1) * (p.Left*10 + 10) * (p.Right*100 + 100))
	// 创建 SHA256 哈希对象
	h := sha256.New()
	// 写入 JSON 数据
	h.Write(segment1)
	h.Write(segment2)
	h.Write(segment3)
	// 获取哈希值
	return fmt.Sprintf("%x", h.Sum(nil))
}

// valid的被动的attr才会应用于champion
func (p *passive) valid() bool {
	if p.Trigger == TimeGoA && p.Left+p.Right > 0 {
		return p.g.now >= p.Left && p.g.now < p.Right
	}
	return p.once == 0
}

// 附加的passive每个事件都会调用handle
func (p *passive) process(e event) {
	if p.once > 1 {
		return
	}

	if !e.match(p.Trigger) {
		return
	}

	if p.Freq == 0 {
		if p.tryCall() && outputLevel >= 3 {
			fmt.Printf("%4.1f秒:常规被动触发\n", p.g.current())
		}
		return
	}

	p.Count++
	if p.Count%p.Freq != 0 {
		return
	}

	if p.once == 1 {
		if p.tryCall() && outputLevel >= 3 {
			fmt.Printf("%4.1f秒:单次被动触发\n", p.g.current())
		}
		p.once = 2
		return
	}

	if p.Freq == 0 {
		if p.tryCall() && outputLevel >= 3 {
			fmt.Printf("%4.1f秒:常规被动触发\n", p.g.current())
		}
		return
	}

	if p.overlay != nil {
		if p.overlay.maxStack == 0 || p.overlay.count < p.overlay.maxStack {
			p.overlay.count++
			p.overlay.attrs.factor = p.overlay.count * 100
			p.attrs = attrs{}
			p.attrs.Add(&p.overlay.attrs)
			if outputLevel >= 3 {
				fmt.Printf("%4.1f秒:第%d次叠加被动\n", p.g.current(), p.overlay.count)
			}
		}
		if p.overlay.reset && p.overlay.count == p.overlay.maxStack {
			if outputLevel >= 3 {
				fmt.Printf("%4.1f秒:可叠加被动重置\n", p.g.current())
			}
			p.overlay.count = 0
			p.attrs = attrs{}
			p.attrs.factor = -1
		}
		return
	}

	if p.tryCall() && outputLevel >= 3 {
		fmt.Printf("%4.1f秒:频率被动触发\n", p.g.current())
	}
}

func (p *passive) body() *attrs {
	return &p.attrs
}

func (p *passive) tryCall() bool {
	if p.call != nil {
		p.call(p.g)
	}
	return p.call != nil
}
