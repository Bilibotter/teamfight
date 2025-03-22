package s14

import (
	"testing"
	C "tft"
)

func TestSyndicate(t *testing.T) {
	C.SortAnalyze()
	C.Level(0)
	twisted3 := C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(3)
	twisted3.Rapidfire(2)
	twisted3.Guinsoo().Jeweled().Guinsoo() // 破防可以换纳什或羊刀，都一样
	twisted3.End(19)
	twisted3.Fight("Twisted[Syn]")

	// 有大头目，红蓝牌取数学期望
	// 攻速可能会降低施法时间
	twisted3 = C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(3)
	twisted3.Rapidfire(2)
	twisted3.Guinsoo().Jeweled().Breaker() // 破防可以换纳什或羊刀，都一样
	twisted3.End(16)
	twisted3.Fight("Twisted[Syn]")

	kogmaw3 := C.Phy(113, 0, 0.7).
		Skill(40, 0).Fury(5, C.TimeGoA, C.AS(50), C.AtkAmp(40*137/100), C.AE(18))
	kogmaw3.BoomBots(4)
	kogmaw3.Rapidfire(2)
	kogmaw3.Infinity().Whisper().GiantSlayer() // 羊刀不如巨杀
	kogmaw3.Fight("Kogmaw[3]")

	kogmaw3 = C.Phy(113, 0, 0.7).
		Skill(40, 0).Fury(5, C.TimeGoA, C.AS(50), C.AtkAmp(40*137/100), C.AE(18))
	kogmaw3.BoomBots(4)
	kogmaw3.Rapidfire(2)
	kogmaw3.Infinity().Whisper().Guinsoo() // 羊刀不如巨杀
	kogmaw3.Fight("Kogmaw[3]*")

	vayne3 := C.Phy(113, 0, 0.7).
		Skill(60, (0.5*113*2+1.5*113)*1.8, 25*1.8).
		Swing(12) //以2.5攻速射3次
	vayne3.Guinsoo().Infinity().GiantSlayer()
	vayne3.Slayer(2)
	vayne3.Animal(3)
	vayne3.Fight("Vayne[3]")

	draven3 := C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Guinsoo().Whisper().Infinity() // 破防更好
	draven3.Fight("Draven1")

	draven3 = C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Breaker().Whisper().Infinity() // 破防更好
	draven3.Fight("Draven2")

	draven3 = C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Shojin().Whisper().Infinity() // 不要青龙刀
	draven3.Fight("Draven3")

	C.FormatResult()
}
