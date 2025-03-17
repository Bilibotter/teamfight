package s14

import (
	"testing"
	C "tft"
)

func Test3Cost(t *testing.T) {
	C.SortAnalyze()
	jinx2 := C.Phy(83, 0, 0.7).
		Skill(50, 1.4*83*5).
		Stack(C.CastStk, 1.4*83).
		Swing(20)
	jinx2.Fight("Jinx[2]")

	jinx3 := C.Phy(124, 0, 0.7).
		Skill(50, 1.5*124*5).
		Stack(C.CastStk, 1.5*124).
		Swing(20)
	jinx3.Fight("Jinx[3]")

	// 刺客
	Rengar2 := C.Phy(98, 20, 0.8).
		Skill(60, 1.8*98*2+2.5*98)
	Rengar2.Fight("Rengar[2]")

	Rengar3 := C.Phy(146, 20, 0.8).
		Skill(60, 1.8*146*2+2.5*146)
	Rengar3.Fight("Rengar[3]")

	// 战士
	fiddkesticks2 := C.Mag(60, 0, 0.7).
		Skill(50, 360*3*3).
		Swing(30)
	fiddkesticks2.Fight("Fiddkesticks[2]")

	fiddkesticks3 := C.Mag(90, 0, 0.7).
		Skill(50, 575*3*3).
		Swing(30)
	fiddkesticks3.Fight("Fiddkesticks[3]")

	varus2 := C.Mag(53, 15, 0.7).
		Skill(75, 300+150*3)
	varus2.Fight("Varus[2]")

	varus3 := C.Mag(79, 15, 0.7).
		Skill(75, 480+240*3)
	varus3.Fight("Varus[3]")

	senna2 := C.Phy(90, 0, 0.6).
		Skill(100, 4*90*2, 30*2).
		Swing(10)
	senna2.Add(C.AtkAmp(33))
	senna2.Fight("Senna[2]")

	senna3 := C.Phy(135, 0, 0.6).
		Skill(100, 4.2*135*2, 45*2).
		Swing(10) // 06:11:70 06:12:70
	senna3.Add(C.AtkAmp(33))
	senna3.Fight("Senna[3]")

	// 未计算A.M.P回蓝
	// 标记会加伤害，后摇未知
	yunnmi2 := C.Mag(53, 0, 0.75).
		Skill(30, 150*1.5).
		Swing(5)
	yunnmi2.Fight("Yunnmi[2]")

	yunnmi3 := C.Mag(79, 0, 0.75).
		Skill(30, 240*1.5).
		Swing(5)
	yunnmi3.Fight("Yunnmi[3]")

	// 有折返伤害
	draven2 := C.Phy(78, 30, 0.75).
		Skill(120, 2*3.3*78*(1+0.8+0.64), 2*45*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven2.Fight("Draven[2]")

	draven3 := C.Phy(117, 30, 0.75).
		Skill(120, 2*3.3*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Fight("Draven[3]")

	// 还减魔抗
	elise2 := C.Mag(60, 0, 0.75).
		Skill(55, 81*12).
		Swing(9) // 11:04:80 11:05:70
	elise2.Fight("Elise[2]")

	elise3 := C.Mag(90, 0, 0.75).
		Skill(55, 140*12).
		Swing(9) // 11:04:80 11:05:70
	elise3.Fight("Elise[3]")

	moderkaiser := C.Mag(75, 30, 0.6).
		Skill(80, 300*2).Fury(2, C.TimeGoA)
	moderkaiser.Fight("Moderkaiser")
	C.FormatResult()
}

func TestSingle3Cost(t *testing.T) {
	C.Level(3)
	//
	jinx2 := C.Phy(83, 0, 0.7).
		Skill(50, 1.4*83*5).
		Stack(C.CastStk, 1.4*83).
		Swing(20)
	jinx2.Fight("Jinx[2]")

	jinx3 := C.Phy(124, 0, 0.7).
		Skill(50, 1.5*124*5).
		Stack(C.CastStk, 1.5*124).
		Swing(20)
	jinx3.Fight("Jinx[3]")
}
