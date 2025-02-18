package tft

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
	asAmp     int  // 攻击次数增幅,影响羊刀叠加
	amp       int  // 伤害增幅
	atkAmp    int  // 普攻增幅
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
	a.asAmp += src.asAmp * src.factor / 100       // 攻击次数增幅
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

func (a *attrs) fac() int {
	return a.factor
}

func (a *attrs) apAE() bool {
	return a.aeType == 0
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
