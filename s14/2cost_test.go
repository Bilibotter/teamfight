package s14

import (
	"testing"
	C "tft"
)

func Test2Cost(t *testing.T) {
	C.SortAnalyze()
	// 考虑过载加成
	amp := 0.0
	nafaari2 := C.Phy(75, 0, 0.7).
		Skill(50, 65*2.65+65*0.65*3+65*1.65*(1+amp)).
		Swing(10)
	nafaari2.Fight("Nafaari[2]")

	nafaari3 := C.Phy(113, 0, 0.7).
		Skill(50, 113*2.65+113*0.65*3+113*1.65*(1+amp)).
		Swing(10)
	nafaari3.Fight("Nafaari[3]")

	// 有大头目，红蓝牌取数学期望
	// 攻速可能会降低施法时间
	twisted2 := C.Mag(53, 10, 0.75).
		Skill(70, 240+420*1.4).
		Swing(10) // 可以改成8
	twisted2.StackPassive(C.AttackA, C.AP(3))
	twisted2.Fight("Twisted[2]")

	twisted3 := C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Fight("Twisted[3]")

	leblanc2 := C.Mag(53, 0, 0.75).
		Skill(50, 5*95).
		Stack(C.CastStk, 95).
		Swing(10)
	leblanc2.Fight("Leblanc[2]")

	leblanc3 := C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10)
	leblanc3.Fight("Leblanc[3]")

	vayne2 := C.Phy(53, 0, 0.7).
		Skill(60, (0.5*53*2+1.5*53)*1.6, 15*1.6).
		Swing(12) //以2.5攻速射3次
	vayne2.Fight("Vayne[2]")

	vayne3 := C.Phy(79, 0, 0.7).
		Skill(60, (0.5*79*2+1.5*79)*1.6, 25*1.6).
		Swing(12) //以2.5攻速射3次
	vayne3.Fight("Vayne[3]")

	// 有大头目，真伤算160%的魔法伤害
	veigar2 := C.Mag(53, 0, 0.7).
		Skill(40, (400+120)*(1+0.2*1.6))
	veigar2.Fight("Veigar[2]")

	veigar3 := C.Mag(79, 0, 0.7).
		Skill(40, (540+150)*(1+0.33*1.6))
	veigar3.Fight("Veigar[3]")

	// 触发大头目
	darius2 := C.Phy(98, 30, 0.65).
		Skill(80, 2*98*3)
	darius2.Fight("Darius[2]")

	darius3 := C.Phy(146, 30, 0.65).
		Skill(80, 2*146*3)
	darius3.Fight("Darius[3]")

	// 双版本
	jhin2 := C.Phy(66, 14, 0.74).
		Skill(84, 1.44*66*3+3.84*66)
	jhin2.Fight("Jhin[2]")

	jhin3 := C.Phy(99, 14, 0.74).
		Skill(84, 1.44*99*3+3.84*99)
	jhin3.Fight("Jhin[3]")

	// 带治疗
	ekko2 := C.Mag(68, 30, 0.7).
		Skill(80, 270)
	ekko2.Fight("Ekko[2]")

	ekko3 := C.Mag(101, 30, 0.7).
		Skill(80, 410)
	ekko3.Fight("Ekko[3]")

	// 考虑加入逛街的影响
	graves2 := C.Phy(87, 0, 0.6).
		Skill(20, 0.4*87*5, 9*5).Swing(15)
	graves2.Add(C.AMP(75))
	graves2.Fight("Graves[2]")

	graves3 := C.Phy(131, 0, 0.6).
		Skill(20, 0.45*131*5, 16*5).Swing(15)
	graves3.Add(C.AMP(75))
	graves3.Fight("Graves[3]")
	C.FormatResult()
}

func TestSingle2Cost(t *testing.T) {
	C.Level(3)

}
