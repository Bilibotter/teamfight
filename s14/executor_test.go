package s14

import (
	"testing"
	C "tft"
)

func TestExecutor(t *testing.T) {
	C.SortAnalyze()

	kobuko := C.Mag(30, 100, 0.8)
	kobuko.Skill(220, 300+20*12*3, 20*5*3).
		Swing(30) // 可能是5秒...
	kobuko.Add(C.AE(135))
	kobuko.KBK()
	kobuko.Cyberboss(2)
	kobuko.Executioner(4)
	kobuko.Breaker().Staff().Sword() // 可能不需要吸血，最多两个大
	kobuko.Beaten()
	kobuko.Fight("Kobuko")

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

	vex := C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten(8) // vex的挨打回蓝可能偏低,同时法力锁时间过长
	vex.Executioner(4)
	vex.Blue().Night().Breaker() // 大天使最优，其次是帽子、法爆
	vex.End(16)
	vex.Fight("Vex")

	urgot := C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.DeathBlade().Whisper().Guinsoo() // 破防可以换羊刀或轻语 泰坦可能也厉害
	urgot.Executioner(4)
	urgot.Fight("Urgot")

	// 考虑加入逛街的影响
	graves3 := C.Phy(124, 0, 0.6).
		Skill(20, 0.45*131*5*2, 16*5).Swing(15)
	graves3.Add(C.AtkAmp(75))
	graves3.GoldenOx(2)
	graves3.Divinicorps(3, C.CritRate(7), C.AD(9))
	graves3.Executioner(5)
	graves3.Silver().Titan().Sword()
	graves3.Fight("Graves")

	C.FormatResult()
}
