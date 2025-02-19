package test

import (
	"testing"
	C "tft"
)

func TestMage(t *testing.T) {
	C.SortAnalyze()
	//C.Level(3)
	ks := C.Mag(108, 75, 1.1).
		Skill(150, 8*90).
		Stack(C.AtkStk, 90).
		Guinsoo().Jeweled().Justice()
	ks.Fight("卡莎")

	kassadin := C.Mag(101, 30, 0.75).
		Skill(70, 255).Fury(1, C.TimeGoA)
	kassadin.StackPassive(C.AfterCastA, C.AE(80))
	kassadin.Add(C.DoubleStrike(70))
	kassadin.Guinsoo().Silver().Sword().Beaten()
	kassadin.Walk()
	kassadin.End(18)
	kassadin.Fight("卡萨丁")

	fiora := C.Phy(113, 70, 1).
		Skill(130, 1.8*113*(3+3+2)/3+3.5*113, 60).
		Fighter().Beaten(4) // s13才改动法力值
	fiora.Add(C.AMP(15))
	fiora.Fight("菲奥娜")

	xerath := C.Mag(90, 0, 0.8).
		Skill(120, 375*10).
		Sorcerer()
	xerath.Fight("泽拉斯")

	xerath0 := C.Mag(90, 0, 0.8).
		Skill(120, 375*10).
		Shojin().Staff().Jeweled()
	xerath0.Fight("泽拉斯0")

	//morgana := C.Mag(72, 0, 0.8).
	//	Skill(60, 1360).
	//	Sorcerer()
	//morgana.Add(C.AP(30))
	//morgana.Fight("莫甘娜")

	morgana := C.Mag(72, 0, 0.8).
		Skill(60, 1600).
		Shojin().Jeweled().Book()
	morgana.Add(C.AP(30))
	morgana.Fight("莫甘娜")
	C.FormatResult()
	t.Logf("Finish")
}

func TestSingle(t *testing.T) {
	C.Level(3)
	//xerath0 := C.Mag(90, 0, 0.8).
	//	Skill(120, 375*10).
	//	Shojin().Staff().Jeweled()
	//xerath0.Fight("泽拉斯0")
	//fiora := C.Phy(113, 70, 1).
	//	Skill(130, 1.8*113*(3+3+2)/3+3.5*113, 60).
	//	Fighter().Beaten(4) // s13才改动法力值
	//fiora.Add(C.AMP(15))
	//fiora.Fight("菲奥娜")

	kassadin := C.Mag(101, 30, 0.75).
		Skill(70, 255).Fury(1, C.TimeGoA)
	kassadin.StackPassive(C.AfterCastA, C.AE(80))
	kassadin.Add(C.DoubleStrike(60 * 2))
	kassadin.Guinsoo().Silver().Sword()
	kassadin.Fight("卡萨丁")
}
