package s14

import (
	"testing"
	C "tft"
)

func Test1Cost(t *testing.T) {
	C.SortAnalyze()
	// morgana垃圾中的垃圾
	amp := 1.0
	// 考虑超载
	nidalee2 := C.Mag(60, 20, 0.75).
		Skill(60, 315+120*(2+amp))
	nidalee2.Jeweled().Justice().Sword()
	nidalee2.Fight("Nidalee[2]")

	nidalee3 := C.Mag(90, 20, 0.75).
		Skill(60, 475+180*(2+amp))
	nidalee3.Fight("Nidalee[3]")

	seraphine2 := C.Mag(53, 10, 0.7).
		Skill(60, 330*(1+0.6+0.36))
	seraphine2.Animal(3)
	seraphine2.Techie(2)
	seraphine2.Shojin().Nashor().Jeweled()
	seraphine2.Fight("Seraphine[2]")

	seraphine3 := C.Mag(79, 10, 0.7).
		Skill(60, 500*(1+0.6+0.36))
	seraphine3.Fight("Seraphine[3]")

	// 考虑成长攻击力
	// 考虑大头目和专属强化
	grow := 40
	shaco2 := C.Phy(83, 0, 0.8).
		Skill(50, 3.4*83+0.5*83)
	shaco2.Syndicate(3)
	shaco2.Slayer(2)
	shaco2.Add(C.AD(grow))
	shaco2.Infinity().Whisper().Guinsoo()
	shaco2.Fight("Shaco[2]")

	grow = 8
	shaco2 = C.Phy(83, 0, 0.8).
		Skill(50, 3.4*83+0.5*83)
	shaco2.Syndicate(3)
	shaco2.Slayer(2)
	shaco2.Add(C.AD(grow))
	shaco2.Infinity().Whisper().Guinsoo()
	shaco2.Fight("Shaco[2]0")

	shaco3 := C.Phy(124, 0, 0.8).
		Skill(50, 3.4*124+0.5*124)
	shaco3.Add(C.AD(grow))
	shaco3.Syndicate(3)
	shaco3.Slayer(4)
	shaco3.Infinity().Sword().Silver()
	shaco3.Fight("Shaco[3]")

	// 考虑机甲伤害
	kogmaw2 := C.Phy(75, 0, 0.7).
		Skill(40, 0).Fury(5, C.TimeGoA, C.AS(40), C.AtkAmp(40*137/100), C.AE(12))
	kogmaw2.BoomBots(2)
	kogmaw2.Rapidfire(2)
	kogmaw2.Infinity().Whisper().Guinsoo() // 羊刀不如巨杀
	kogmaw2.Fight("Kogmaw[2]")

	kogmaw3 := C.Phy(113, 0, 0.7).
		Skill(40, 0).Fury(5, C.TimeGoA, C.AS(50), C.AtkAmp(40*137/100), C.AE(18))
	kogmaw3.BoomBots(2)
	kogmaw3.Rapidfire(2)
	kogmaw3.Infinity().Whisper().GiantSlayer() // 羊刀不如巨杀
	kogmaw3.Fight("Kogmaw[3]")

	zyra2 := C.Mag(45, 10, 0.75).
		Skill(60, 375+180)
	zyra2.StreetDemon(3)
	zyra2.Techie(2)
	zyra2.Nashor().Shojin().Jeweled()
	zyra2.Fight("Zyra[2]")

	zyra3 := C.Mag(68, 10, 0.75).
		Skill(60, 585+270)
	zyra3.Fight("Zyra[3]")

	kindred2 := C.Phy(68, 0, 0.7).
		Skill(50, 5.2*68, 30).Swing(10)
	kindred2.Marksman(2)
	kindred2.Shojin().Infinity().Whisper()
	kindred2.Fight("Kindred[2]M")

	kindred2 = C.Phy(68, 0, 0.7).
		Skill(50, 5.2*68, 30).Swing(10)
	kindred2.Rapidfire(2)
	kindred2.Shojin().Infinity().Whisper()
	kindred2.Fight("Kindred[2]R")

	kindred3 := C.Phy(101, 0, 0.7).
		Skill(50, 5.2*101, 45)
	kindred3.Marksman(2)
	kindred3.Rapidfire(2)
	kindred3.Shojin().Infinity().Whisper()
	kindred3.Fight("Kindred[3]")
	C.FormatResult()
}

func TestSingle1Cost(t *testing.T) {
	C.Level(0)
	kogmaw3 := C.Phy(113, 0, 0.7).
		Skill(40, 0).Fury(5, C.TimeGoA, C.AS(50), C.AtkAmp(40*137/100), C.AE(18))
	kogmaw3.BoomBots(2)
	kogmaw3.Rapidfire(2)
	kogmaw3.Infinity().Whisper().GiantSlayer() // 羊刀不如巨杀
	kogmaw3.Fight("Kogmaw[3]")

	kogmaw3 = C.Phy(113, 0, 0.7).
		Skill(40, 0).Fury(5, C.TimeGoA, C.AS(50), C.AtkAmp(40*137/100), C.AE(18))
	kogmaw3.BoomBots(2)
	kogmaw3.Rapidfire(2)
	kogmaw3.Infinity().Whisper().Guinsoo() // 羊刀不如巨杀
	kogmaw3.Fight("Kogmaw[3]*")
}
