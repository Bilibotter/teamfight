package s14

import (
	"testing"
	C "tft"
)

func Test2Cost(t *testing.T) {
	C.SortAnalyze()
	//graves2 := C.Phy(87, 0, 0.6).
	//	Skill(20, 0.4*87*5, 9*5).Swing(15)
	//graves2.Add(C.AMP(75))
	//graves2.Fight("Graves")

	// 考虑加入逛街的影响
	graves3 := C.Phy(131, 0, 0.6).
		Skill(20, 0.45*131*5*2, 16*5).Swing(15)
	graves3.Add(C.AtkAmp(75))
	graves3.GoldenOx(2)
	graves3.Executioner(4)
	graves3.Justice().Whisper().Sword()
	graves3.Fight("Graves")

	//nafaari2 := C.Phy(75, 0, 0.7).
	//	Skill(50, 65*2.65+65*0.65*3+65*1.65*(1+C.Overload(amp))).
	//	Swing(10)
	//nafaari2.Fight("Nafaari[2]")
	amp := 4
	nafaari3 := C.Phy(113, 0, 0.8).
		Skill(50, 113*2.65+113*0.65*3+113*1.65*(1+C.Overload(amp))).
		Swing(10)
	nafaari3.Beaten()
	nafaari3.Sword().Sterak().Titan()
	nafaari3.Fight("Nafaari")

	//twisted2 := C.Mag(53, 0, 0.75).
	//twisted2 := C.Mag(53, 10, 0.75).
	//	Skill(70, 240+420*1.4).
	//	Swing(10) // 可以改成8
	//twisted2.StackPassive(C.AttackA, C.AP(3))
	//twisted2.Fight("Twisted[2]")

	// 有大头目，红蓝牌取数学期望
	// 攻速可能会降低施法时间
	twisted3 := C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(5)
	twisted3.Rapidfire(2)
	twisted3.Guinsoo().Jeweled().Breaker() // 破防可以换纳什或羊刀，都一样
	twisted3.Fight("Twisted[Syn]")

	twisted3 = C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(3)
	twisted3.Rapidfire(4) // 4迅捷射手也厉害
	twisted3.Guinsoo().Jeweled().Breaker()
	twisted3.Fight("Twisted[Rap]")

	//leblanc2 := C.Mag(53, 0, 0.75).
	//	Skill(50, 5*95).
	//	Stack(C.CastStk, 95).
	//	Swing(10)
	//leblanc2.Fight("Leblanc[2]")

	leblanc3 := C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10) // 06:59:30 07:00:30
	leblanc3.Strategist(4)
	leblanc3.Shojin().Nashor().Staff() // 法爆也行
	leblanc3.Fight("Leblanc")

	//vayne2 := C.Phy(53, 0, 0.7).
	//	Skill(60, (0.5*53*2+1.5*53)*1.6, 15*1.6).
	//	Swing(12) //以2.5攻速射3次
	//vayne2.Fight("Vayne[2]")

	vayne3 := C.Phy(113, 0, 0.7).
		Skill(60, (0.55*113*2+1.65*113)*1.8, 25*1.8).
		Swing(12) //以2.5攻速射3次
	vayne3.Guinsoo().Infinity().GiantSlayer()
	vayne3.Slayer(6)
	vayne3.Fight("Vayne[3]")

	// 有大头目，真伤算160%的魔法伤害
	veigar3 := C.Mag(79, 0, 0.7).
		Skill(40, (540+150)*(0.67+0.33*1.6)).
		Swing(4) // 06:47:70 06:48:40
	veigar3.Cyberboss(2)
	veigar3.Techie(4)
	veigar3.Shojin().Nashor().Jeweled() // 蓝buff也行
	veigar3.Fight("Veigar")

	//// 触发大头目
	//darius2 := C.Phy(98, 30, 0.65).
	//	Skill(80, 2*98*3)
	//darius2.Fight("Darius[2]")
	//
	//darius3 := C.Phy(146, 30, 0.65).
	//	Skill(80, 2*146*3)
	//darius3.Fight("Darius[3]")
	//
	//// 双版本
	//jhin2 := C.Phy(66, 14, 0.74).
	//	Skill(84, 1.44*66*3+3.84*66)
	//jhin2.Fight("Jhin[2]")
	//
	//jhin3 := C.Phy(102, 14, 0.74).
	//	Skill(84, 1.44*102*3+3.84*102)
	//jhin3.Fight("Jhin[3]")
	//
	//// 带治疗
	//ekko2 := C.Mag(68, 30, 0.7).
	//	Skill(80, 270)
	//ekko2.Fight("Ekko[2]")
	//
	//ekko3 := C.Mag(101, 30, 0.7).
	//	Skill(80, 410)
	//ekko3.Fight("Ekko[3]")

	C.FormatResult()
}

func TestSingle2Cost(t *testing.T) {
	C.Level(3)
	twisted3 := C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(5)
	twisted3.Rapidfire(2)
	twisted3.Guinsoo().Jeweled().Breaker() // 破防可以换纳什或羊刀，都一样
	twisted3.End(16)
	twisted3.Fight("Twisted[Syn]")

	C.FormatResult()
}

func TestItemsGraves(t *testing.T) {
	C.SortAnalyze()
	graves3 := C.Phy(131, 0, 0.6).
		Skill(20, 0.45*131*5*2, 16*5).Swing(15)
	graves3.Add(C.AtkAmp(75))
	graves3.GoldenOx(2)
	graves3.Executioner(4)
	graves3.Justice().Whisper().Sword()
	graves3.Fight("Graves[3]")

	graves3 = C.Phy(131, 0, 0.6).
		Skill(20, 0.45*131*5*2, 16*5).Swing(15)
	graves3.Add(C.AtkAmp(75))
	graves3.GoldenOx(2)
	graves3.Executioner(4)
	graves3.Justice().Silver().Sword()
	graves3.Fight("Graves[3]1")

	graves3 = C.Phy(131, 0, 0.6).
		Skill(20, 0.45*131*5*2, 16*5).Swing(15)
	graves3.Add(C.AtkAmp(75))
	graves3.GoldenOx(2)
	graves3.Executioner(4)
	graves3.Justice().Breaker().Sword()
	graves3.Fight("Graves[3]11")
	C.FormatResult()
}

func TestItemsTwistedFate(t *testing.T) {
	C.SortAnalyze()
	twisted3 := C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(5)
	twisted3.Rapidfire(2)
	twisted3.Guinsoo().Jeweled().Breaker() // 破防可以换纳什或羊刀，都一样
	twisted3.End(16)
	twisted3.Fight("Twisted[Syn]")

	twisted3 = C.Mag(79, 10, 0.75).
		Skill(70, 375+650*1.4).
		Swing(10) // 可以改成8
	twisted3.StackPassive(C.AttackA, C.AP(3))
	twisted3.Syndicate(3)
	twisted3.Rapidfire(4) // 4迅捷射手也厉害
	twisted3.Guinsoo().Jeweled().Breaker()
	twisted3.End(16)
	twisted3.Fight("Twisted[Rap]")
	//twisted3 := C.Mag(79, 10, 0.75).
	//	Skill(70, 375+650*1.4).
	//	Swing(10) // 可以改成8
	//twisted3.StackPassive(C.AttackA, C.AP(3))
	//twisted3.Syndicate(5)
	//twisted3.Rapidfire(2)
	//twisted3.Guinsoo().Jeweled().Breaker()
	//twisted3.Fight("Twisted[3]")
	//
	//twisted3 = C.Mag(79, 10, 0.75).
	//	Skill(70, 375+650*1.4).
	//	Swing(10) // 可以改成8
	//twisted3.StackPassive(C.AttackA, C.AP(3))
	//twisted3.Syndicate(3)
	//twisted3.Rapidfire(4)
	//twisted3.Guinsoo().Jeweled().Breaker()
	//twisted3.Fight("Twisted[3]*")
	//
	//twisted3 = C.Mag(79, 10, 0.75).
	//	Skill(70, 375+650*1.4).
	//	Swing(10) // 可以改成8
	//twisted3.StackPassive(C.AttackA, C.AP(3))
	//twisted3.Syndicate(5)
	//twisted3.Rapidfire(2)
	//twisted3.Guinsoo().Jeweled().DeathCap()
	//twisted3.Fight("Twisted[3]1")
	//
	//twisted3 = C.Mag(79, 10, 0.75).
	//	Skill(70, 375+650*1.4).
	//	Swing(10) // 可以改成8
	//twisted3.StackPassive(C.AttackA, C.AP(3))
	//twisted3.Syndicate(5)
	//twisted3.Rapidfire(2)
	//twisted3.DoubleGuinsoo().Jeweled()
	//twisted3.Fight("Twisted[3]11")
	//
	//twisted3 = C.Mag(79, 10, 0.75).
	//	Skill(70, 375+650*1.4).
	//	Swing(10) // 可以改成8
	//twisted3.StackPassive(C.AttackA, C.AP(3))
	//twisted3.Syndicate(5)
	//twisted3.Rapidfire(2)
	//twisted3.Guinsoo().Nashor().Jeweled()
	//twisted3.Fight("Twisted[3]111")
	C.FormatResult()
}

func TestLeblanc(t *testing.T) {
	C.SortAnalyze()
	leblanc3 := C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10)
	leblanc3.Strategist(4)
	leblanc3.Shojin().Nashor().Jeweled()
	leblanc3.Fight("Leblanc[3]")

	leblanc3 = C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10)
	leblanc3.Strategist(4)
	leblanc3.Blue().Nashor().Jeweled()
	leblanc3.Fight("Leblanc[3]1")

	leblanc3 = C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10)
	leblanc3.Strategist(4)
	leblanc3.Blue().Nashor().Staff()
	leblanc3.Fight("Leblanc[3]11")

	leblanc3 = C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10)
	leblanc3.Strategist(4)
	leblanc3.Shojin().Nashor().Staff()
	leblanc3.Fight("Leblanc[3]11*")

	leblanc3 = C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10)
	leblanc3.Strategist(4)
	leblanc3.Shojin().Jeweled().Staff()
	leblanc3.Fight("Leblanc[3]111")
	C.FormatResult()
}

func TestVeigar(t *testing.T) {
	C.SortAnalyze()
	veigar3 := C.Mag(79, 0, 0.7).
		Skill(40, (540+150)*(0.67+0.33*1.6)).
		Swing(4) // 06:47:70 06:48:40
	veigar3.Cyberboss(2)
	veigar3.Techie(4)
	veigar3.Blue().Nashor().Jeweled()
	veigar3.Fight("Veigar[3]")

	veigar3 = C.Mag(79, 0, 0.7).
		Skill(40, (540+150)*(0.67+0.33*1.6)).
		Swing(4) // 06:47:70 06:48:40
	veigar3.Cyberboss(2)
	veigar3.Techie(4)
	veigar3.Shojin().Nashor().Jeweled()
	veigar3.Fight("Veigar[3]1")

	veigar3 = C.Mag(79, 0, 0.7).
		Skill(40, (540+150)*(0.67+0.33*1.6)).
		Swing(4) // 06:47:70 06:48:40
	veigar3.Cyberboss(2)
	veigar3.Techie(4)
	veigar3.Blue().GiantSlayer().Jeweled()
	veigar3.Fight("Veigar[3]11")
	C.FormatResult()
}
