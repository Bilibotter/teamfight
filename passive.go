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
freq: 到达一定频次再执行call。用于周期触发被动与可叠加被动。
once: 1: 表示被动只触发一次，2: 表示已触发
left：once=1，表示被动触发的阈值; once=0，表示时间段被动的左边界，包含边界值。
right：表示时间段被动的右边界，不包含边界值。
trigger：表示触发call的事件类型。

	频次+可叠加被动支持触发类型：TimeGo、Cast、Attack
	时间段被动支持触发类型：TimeGo
*/
type passive struct {
	attrs
	name    string
	g       *ground
	overlay *stack
	trigger action
	freq    int             // 匹配事件全部事件，每多少次事件触发一次
	left    int             // 匹配事件timeGoA，界定的数值/左边界值
	right   int             // 匹配事件timeGoA，右边界值，不包含右边界
	call    func(g *ground) // 匹配到事件时调用
	once    int             // once:0表示无限次,1表示只触发一次,2表示已触发
}

func buffPassive(trigger action, remain int, a ...*attrs) *passive {
	p := &passive{
		trigger: trigger,
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
		attrs:   attr,
		overlay: &stack{attrs: attr, maxStack: limit},
		trigger: trigger,
		freq:    freq,
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
		trigger: trigger,
	}
	return p
}

func (p *passive) key() string {
	segment1, _ := json.Marshal(p.attrs)
	segment2, _ := json.Marshal(p.trigger)
	segment3, _ := json.Marshal((p.freq*1 + 1) * (p.left*10 + 10) * (p.right*100 + 100))
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
	if p.trigger == TimeGoA && p.left+p.right > 0 {
		return p.g.now >= p.left && p.g.now < p.right
	}
	return p.once == 0
}

// 附加的passive每个事件都会调用handle
func (p *passive) process(e event) {
	if p.once > 1 {
		return
	}
	if p.once == 1 && e.match(p.trigger) {
		if p.left == 0 {
			panic("未设定once的临界值")
		}
		p.left--
		if p.left == 0 {
			if p.tryCall() && outputLevel >= 3 {
				fmt.Printf("%d秒:单次被动触发\n", p.g.now)
			}
			p.once++
			p.left-- // 使p.left永远不为0
		}
		return
	}

	if !e.match(p.trigger) {
		return
	}

	if p.freq == 0 {
		if p.tryCall() && outputLevel >= 3 {
			fmt.Printf("%d秒:常规被动触发\n", p.g.now)
		}
		return
	}

	times := 0
	switch p.trigger {
	case BeforeCastA, AfterCastA:
		times = p.g.castTimes
	case TimeGoA:
		times = p.g.now
	case AttackA:
		times = p.g.atkTimes
	default:
		panic("未知频率事件")
	}

	enough := times > 0 && times%p.freq == 0

	if p.overlay != nil {
		if !enough {
			return
		}
		if p.overlay.maxStack == 0 || p.overlay.count < p.overlay.maxStack {
			p.overlay.count++
			p.overlay.attrs.factor = p.overlay.count * 100
			p.attrs = attrs{}
			p.attrs.Add(&p.overlay.attrs)
			if outputLevel >= 3 {
				fmt.Printf("%4d秒:第%d次叠加被动\n", p.g.now, p.overlay.count)
			}
		}
		if p.overlay.reset && p.overlay.count == p.overlay.maxStack {
			if outputLevel >= 3 {
				fmt.Printf("%d秒:可叠加被动重置\n", p.g.now)
			}
			p.overlay.count = 0
			p.attrs = attrs{}
			p.attrs.factor = -1
		}
		return
	}

	if enough {
		if p.tryCall() && outputLevel >= 3 {
			fmt.Printf("%d秒:频率被动触发\n", p.g.now)
		}
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
