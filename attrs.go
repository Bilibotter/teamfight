package tft

import (
	"fmt"
	"strings"
)

type processI interface {
	process(event)
}

type attribute interface {
	body() *attrs
	fac() int
	validI
}

type validI interface {
	valid() bool
}

type attrs struct {
	factor    int  // 倍数，默认100
	ad        int  // 攻击力
	ap        int  // 法强
	as        int  // 攻速
	ae        int  // 附带特效伤害
	aeType    int  // 附带特效类型,0为ap加成
	strike    int  // 触发双重打击的概率
	amp       int  // 伤害增幅
	atkAmp    int  // 普攻增幅，金克斯火箭，泽丽镜像，月男炮台，凯隐旋转
	castAmp   int  // 施法增幅
	critRate  int  // 暴击率
	critAmp   int  // 暴击伤害
	skillCrit bool // 是否技能暴击
	hitMana   int  // 攻击回复法力值
	manaSrk   int  // 减少的最大魔法值
	manaAmp   int  // 法力增幅
	manaDec   int  // 减少的最大魔法值
}

func (a *attrs) Add(src *attrs) *attrs {
	if src.factor < 0 {
		return a
	}
	if src.factor == 0 {
		src.factor = 100
	}
	// 倍数默认为100，不影响加成
	// 下面是按每个属性、factor进行加成处理
	a.ad += src.ad * src.factor / 100             // 攻击力
	a.ap += src.ap * src.factor / 100             // 法强
	a.as += src.as * src.factor / 100             // 攻速
	a.ae += src.ae * src.factor / 100             // 附带特效伤害
	a.amp += src.amp * src.factor / 100           // 伤害增幅
	a.atkAmp += src.atkAmp * src.factor / 100     // 普攻增幅
	a.strike += src.strike * src.factor / 100     // 攻击次数增幅
	a.castAmp += src.castAmp * src.factor / 100   // 施法增幅
	a.critRate += src.critRate * src.factor / 100 // 暴击率
	a.critAmp += src.critAmp * src.factor / 100   // 暴击伤害
	a.hitMana += src.hitMana * src.factor / 100   // 攻击回复法力值
	a.manaSrk += src.manaSrk * src.factor / 100   // 减少的最大魔法值 (这里的名称可能有误，可能是 manaDec？)

	a.manaAmp += src.manaAmp * src.factor / 100 // 法力增幅
	a.manaDec += src.manaDec * src.factor / 100 // 减少的最大魔法值

	a.skillCrit = a.skillCrit || src.skillCrit

	return a
}

// kimi生成
func (a *attrs) Signature() string {
	var sb strings.Builder

	// 检查每个字段是否不为空，并写入前缀和数值
	if a.factor != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "factor", a.factor))
	}
	if a.ad != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "ad", a.ad))
	}
	if a.ap != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "ap", a.ap))
	}
	if a.as != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "as", a.as))
	}
	if a.ae != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "ae", a.ae))
	}
	if a.aeType != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "aeType", a.aeType))
	}
	if a.strike != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "strike", a.strike))
	}
	if a.amp != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "amp", a.amp))
	}
	if a.atkAmp != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "atkAmp", a.atkAmp))
	}
	if a.castAmp != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "castAmp", a.castAmp))
	}
	if a.critRate != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "critRate", a.critRate))
	}
	if a.critAmp != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "critAmp", a.critAmp))
	}
	if a.skillCrit {
		sb.WriteString(fmt.Sprintf("%s%d", "skillCrit", 1))
	}
	if a.hitMana != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "hitMana", a.hitMana))
	}
	if a.manaSrk != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "manaSrk", a.manaSrk))
	}
	if a.manaAmp != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "manaAmp", a.manaAmp))
	}
	if a.manaDec != 0 {
		sb.WriteString(fmt.Sprintf("%s%d", "manaDec", a.manaDec))
	}

	return sb.String()
}

func (a *attrs) fac() int {
	return a.factor
}

func (g *champion) apAE() bool {
	return g.dmgType == magHero
}

func AE(i int) *attrs {
	return &attrs{ae: i, factor: 100}
}

func DoubleStrike(i int) *attrs {
	return &attrs{strike: i, factor: 100}
}

func AtkAmp(i int) *attrs {
	return &attrs{atkAmp: i, factor: 100}
}

func CastAmp(i int) *attrs {
	return &attrs{castAmp: i, factor: 100}
}

func ManaAmp(i int) *attrs {
	return &attrs{manaAmp: i, factor: 100}
}

func AD(i int) *attrs {
	return &attrs{ad: i, factor: 100}
}

func AP(i int) *attrs {
	return &attrs{ap: i, factor: 100}
}

func AS(i int) *attrs {
	return &attrs{as: i, factor: 100}
}

func AMP(i int) *attrs {
	return &attrs{amp: i, factor: 100}
}

// 技能暴击
func Critable() *attrs {
	return &attrs{skillCrit: true}
}

func CritRate(i int) *attrs {
	return &attrs{critRate: i, factor: 100}
}

func CritAmp(i int) *attrs {
	return &attrs{critAmp: i, factor: 100}
}
