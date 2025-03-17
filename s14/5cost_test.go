package s14

import (
	"testing"
	C "tft"
)

// 棋子与星级关系为1.8
const fac = 1.8

func Test5Cost(t *testing.T) {
	C.SortAnalyze()
	urgot := C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125)).
		Swing(0)
	urgot.Fight("Urgot")

	// 加蓝的被动待做
	kobuko := C.Mag(30, 175, 0.8)
	kobuko.Skill(150, 300+20*5*3, 20*12*3).Swing(15)
	kobuko.Add(C.AE(135))
	kobuko.BuffPassive(C.TimeGoA, 0, 8, C.AS(100))
	kobuko.Beaten()
	kobuko.Fight("Kobuko")

	// 未计算距离衰减
	reduce := 0.866
	samira := C.Phy(90, 20, 0.8).
		Skill(100, 90*25*0.8+90*1.8*5*reduce)
	samira.Fight("Samira")

	// 双版本
	aurora := C.Mag(75, 20, 0.8).
		Skill(80, 900+125*3)
	aurora.Fight("Aurora")

	renekton := C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(100, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.TimeGoA, C.AtkAmp(50))
	renekton.Beaten()
	renekton.Fight("Renekton")

	viego := C.Mag(90, 20, 0.9).
		Skill(80, 360+180*2)
	viego.Beaten()
	viego.Fight("Viego")

	garen := C.Phy(120, 30, 0.85).
		Skill(100, 6.5*120+2.6*120*3).Fury(3, C.TimeGoA)
	garen.Beaten()
	garen.Fight("Garen")
	C.FormatResult()
}

func TestSingle5Cost(t *testing.T) {
	C.Fac()
	C.Level(3)
}
