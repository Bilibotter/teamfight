package s14

import (
	"testing"
	C "tft"
)

func TestDivinicorps(t *testing.T) {
	C.SortAnalyze()
	zed := C.Phy(82, 0, 0.95).
		Skill(50, 2.7*82*1.5+2.0*82*1.5, 60, 50).
		//Stack(C.CastStk, 2.0*82*1.5).
		Swing(12) // 03:40:50 03:42:00
	zed.Beaten()
	zed.Slayer(4)
	zed.GoldenOx(2)
	zed.Infinity().Silver().Justice()
	zed.Fight("Zed[S]")

	zed = C.Phy(82, 0, 0.95).
		Skill(50, 2.7*82*1.5+2.0*82*1.5, 60, 50).
		//Stack(C.CastStk, 2.0*82*1.5).
		Swing(12) // 03:40:50 03:42:00
	zed.Beaten()
	zed.Slayer(4)
	zed.GoldenOx(2)
	zed.Infinity().Breaker().Justice()
	zed.Fight("Zed[Brk]")

	zed = C.Phy(90, 0, 0.95).
		Skill(50, 2.25*90*1.5+2.0*90*1.5, 60, 50).
		//Stack(C.CastStk, 2.0*90*1.5).
		Swing(15) // 03:40:50 03:42:00
	zed.Beaten()
	zed.Slayer(4)
	zed.GoldenOx(2)
	zed.Infinity().Silver().Justice()
	zed.Fight("Zed[S]*")

	renekton := C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	eat := 98.0 * 0.33 / 105.0 * 100.0 // 2星诺手
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	renekton.Silver().Titan0().Sword()
	renekton.Fight("Renekton[D]")

	renekton = C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	renekton.Titan().Silver().Sword()
	renekton.Fight("Renekton[D]*")

	renekton = C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(2, C.CritRate(7), C.AS(5))
	renekton.Silver().Titan0().Sword()
	renekton.Executioner(3)
	renekton.Fight("Renekton[E]")

	renekton = C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(2, C.CritRate(7), C.AS(5))
	renekton.Titan().Silver().Sword()
	renekton.Executioner(3)
	renekton.Fight("Renekton[E]*")

	vex := C.Mag(45, 0, 0.8).
		Skill(30, 150*2+300+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Executioner(3)
	vex.Blue().Breaker().Staff()
	vex.Fight("Vex[E]")

	vex = C.Mag(45, 0, 0.8).
		Skill(30, 150*2+300+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	vex.Blue().Breaker().Jeweled()
	vex.Fight("Vex[D]")

	aurora := C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	aurora.Shojin().Breaker().Jeweled()
	aurora.Fight("Aurora[D]")

	aurora = C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	aurora.Shojin().Staff().Jeweled()
	aurora.Fight("Aurora[D]*")

	aurora = C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	aurora.Shojin().Nashor().Jeweled()
	aurora.Fight("Aurora[D]**")

	viego := C.Mag(90, 20, 0.9).
		Skill(65, 345+180*(5+3+3)/3).
		Swing(11) // 17:21:70 17:22:80
	viego.Beaten()
	viego.Jeweled().Sword().Justice()
	viego.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	viego.Fight("Viego")

	garen := C.Phy(120, 30, 0.85).
		Skill(100, 6.5*120+2.6*120*2).Fury(2, C.TimeGoA).
		Skill(70, 6.5*120+2.6*120*2).Fury(2, C.TimeGoA)
	garen.Beaten()
	garen.Sword().Sterak().Titan()
	garen.Divinicorps(6, C.CritRate(7), C.AS(5), C.AD(9), C.AP(9))
	garen.Fight("Garen")

	C.FormatResult()
}
