package s14

import (
	"fmt"
	"testing"
	C "tft"
)

// 棋子与星级关系为1.8
const fac = 1.8

func Test5Cost(t *testing.T) {
	C.SortAnalyze()
	urgot := C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.DeathBlade().Whisper().Guinsoo() // 破防可以换羊刀或轻语 泰坦可能也厉害
	urgot.Executioner(5)
	urgot.Fight("Urgot")

	kobuko := C.Mag(30, 100, 0.8)
	kobuko.Skill(220, 300+20*12*3, 20*5*3).
		Swing(30) // 可能是5秒...
	kobuko.Add(C.AE(135))
	kobuko.KBK()
	kobuko.Cyberboss(2)
	kobuko.Jeweled().Staff().Sword() // 可能不需要吸血，最多两个大
	kobuko.Beaten()
	kobuko.Fight("Kobuko")

	amp := 5
	reduce := (2 + 0.8*3) / 5
	samira := C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
		Swing(25)
	samira.Night().Infinity().Shojin()
	samira.End(17)
	samira.Fight("Samira")

	// 双版本
	aurora := C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Dynamo(4) // 阿罗拉在人造人里面是大爹，幻灵里面差一点
	aurora.Jeweled().Staff().Breaker()
	aurora.Fight("Aurora")

	// 普攻倍率不太确定
	renekton := C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	eat := 98.0 * 0.33 / 105.0 * 100.0 // 2星诺手
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(6, C.AP(9), C.AD(8), C.CritRate(8), C.AS(5))
	renekton.Guinsoo().Silver().Sword()
	renekton.Fight("Renekton")

	viego := C.Mag(90, 20, 0.9).
		Skill(80, 360+180*(5+3+3)/3).
		Swing(25) // 01:31:70 01:34:20
	viego.Beaten()
	viego.Jeweled().Sword().Justice()
	viego.GoldenOx(6) //开6福牛也不厉害
	viego.Techie(2)
	viego.Fight("Viego")

	garen := C.Phy(120, 30, 0.85).
		Skill(100, 6.5*120+2.6*120*3).Fury(3, C.TimeGoA)
	garen.Beaten()
	garen.Sword().Sterak().Titan()
	garen.Fight("Garen")
	C.FormatResult()
}

func TestSingle5Cost(t *testing.T) {
	C.Level(3)
	// 普攻倍率不太确定
	renekton := C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	eat := 98.0 * 0.33 / 105.0 * 100.0 // 2星诺手
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(2, C.CritRate(8), C.AS(5))
	renekton.Guinsoo().Silver().Sword()
	renekton.Executioner(4)
	renekton.Fight("Renekton")
}

func TestSamira(t *testing.T) {
	C.SortAnalyze()
	amp := 5
	reduce := (1 + 0.8*3) / 4
	fmt.Printf("%.1f\n", 0.8*90*(25+5*C.Overload(amp))*reduce+90*5)
	samira := C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*5).
		Swing(25)
	samira.Night().Infinity().Justice()
	samira.Fight("Samira")

	samira = C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*5).
		Swing(25)
	samira.Night().Infinity().Shojin()
	samira.Fight("Samira1")

	samira = C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*5).
		Swing(25)
	samira.Night().Infinity().Breaker()
	samira.Fight("Samira2")
	C.FormatResult()
}

func TestRenekton(t *testing.T) {
	C.SortAnalyze()
	// 普攻倍率不太确定
	renekton := C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	eat := 98.0 * 0.33 / 105.0 * 100.0 // 2星诺手
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(6, C.AP(9), C.AD(8), C.CritRate(8), C.AS(5))
	renekton.Guinsoo().Silver().Sword()
	renekton.Fight("Renekton")

	renekton = C.Phy(105, 70, 0.91).
		Skill(150, 3*105, 40).Inspire(50, C.TimeGoA, C.AtkAmp(50)).
		Skill(150, 0).Fury(100, C.AttackA, C.AtkAmp(50)) // reduce方式不同以避免去重
	renekton.Beaten()
	eat = 98.0 * 0.33 / 105.0 * 100.0 // 2星诺手
	renekton.Add(C.AD(int(eat)))
	renekton.Divinicorps(6, C.AP(9), C.AD(8), C.CritRate(8), C.AS(5))
	renekton.Infinity().Silver().Sword()
	renekton.Fight("Renekton")
	C.FormatResult()
}

func TestAurora(t *testing.T) {
	C.SortAnalyze()
	// 双版本
	aurora := C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Dynamo(4)
	aurora.Jeweled().Staff().Breaker()
	aurora.Fight("Aurora")

	aurora = C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Animal(7)
	aurora.Dynamo(2)
	aurora.Jeweled().Staff().Shojin()
	aurora.Fight("Aurora1")

	aurora = C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Animal(7)
	aurora.Dynamo(2)
	aurora.Jeweled().Shojin().Breaker()
	aurora.Fight("Aurora1")
	C.FormatResult()
}

func TestUrgot(t *testing.T) {
	C.SortAnalyze()
	urgot := C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.Breaker().Whisper().Guinsoo()
	urgot.Executioner(5)
	urgot.Fight("Urgot")

	urgot = C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.Breaker().Whisper().Shojin()
	urgot.Executioner(5)
	urgot.Fight("Urgot1")

	urgot = C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.DeathBlade().Whisper().Shojin()
	urgot.Executioner(5)
	urgot.Fight("Urgot2")

	urgot = C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.DeathBlade().Whisper().Guinsoo()
	urgot.Executioner(5)
	urgot.Fight("Urgot3")

	urgot = C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.DeathBlade().Whisper().GiantSlayer()
	urgot.Executioner(5)
	urgot.Fight("Urgot4")
	C.FormatResult()
}
