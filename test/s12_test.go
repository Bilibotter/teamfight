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
}
