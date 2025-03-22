package s14

import (
	"testing"
	C "tft"
)

func Test3Cost(t *testing.T) {
	C.SortAnalyze()
	// 有折返伤害
	//draven2 := C.Phy(78, 30, 0.75).
	//	Skill(120, 2*3.3*78*(1+0.8+0.64), 2*45*(1+0.8+0.64)).
	//	Swing(12) // 03:40:80 03:42:00
	//draven2.Fight("Draven[2]")

	draven3 := C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Guinsoo().Whisper().Infinity() // 破防更好
	draven3.Fight("Draven")

	//jinx2 := C.Phy(83, 0, 0.7).
	//	Skill(50, 1.4*83*5).
	//	Stack(C.CastStk, 1.4*83).
	//	Swing(20)
	//jinx2.Fight("Jinx[2]")

	jinx3 := C.Phy(124, 0, 0.7).
		Skill(50, 1.5*124*5, 40).
		Stack(C.CastStk, 1.5*124+40).
		Swing(20)
	//jinx3.StreetDemon(7)
	jinx3.Marksman(2)
	jinx3.StreetDemon(3)
	jinx3.Strategist0(2)
	jinx3.Shojin().Infinity().Breaker() // 羊刀也可以接受
	jinx3.Fight("Jinx")

	//Rengar2 := C.Phy(98, 20, 0.8).
	//	Skill(60, 1.8*98*2+2.5*98).
	//	Swing(14)
	//Rengar2.Fight("Rengar[2]")

	// 刺客
	Rengar3 := C.Phy(146, 20, 0.8).
		Skill(60, 1.8*146*2+2.5*146).
		Swing(14) //05:46:00 05:47:40
	Rengar3.Executioner(4)
	Rengar3.Beaten()
	Rengar3.Justice().Sword().Titan() // 治疗量有AP加成
	Rengar3.Fight("Rengar")

	senna2 := C.Phy(120, 0, 0.6).
		Skill(100, 2.8*120*2, 30*2).
		Swing(10)
	senna2.Add(C.AtkAmp(40))
	senna2.Slayer(6)
	senna2.Guinsoo().Whisper().Infinity()
	senna2.Fight("Senna[2]")

	senna3 := C.Phy(180, 0, 0.6).
		Skill(100, 3*180*2, 45*2).
		Swing(10) // 06:11:70 06:12:70
	senna3.Add(C.AtkAmp(40))
	senna3.Slayer(6)
	senna3.Guinsoo().Whisper().Infinity()
	senna3.Fight("Senna")

	//elise2 := C.Mag(60, 0, 0.75).
	//	Skill(55, 81*8).
	//	Swing(9) // 11:04:80 11:05:70
	//elise2.Fight("Elise[2]")

	// 还减魔抗
	elise3 := C.Mag(90, 0, 0.75).
		Skill(55, 140*8).
		Swing(9) // 11:04:80 11:05:70
	elise3.Dynamo(4)
	elise3.Jeweled().Breaker().Staff() // 破防可以换帽子
	elise3.Fight("Elise")

	//yuumi2 := C.Mag(53, 0, 0.75).
	//	Skill(30, 150*1.5).
	//	Swing(5)
	//yuumi2.Fight("yuumi[2]")

	// 标记会加伤害，后摇未知
	yuumi3 := C.Mag(79, 0, 0.75).
		Skill(30, (240+480)/2).
		Swing(6)
	yuumi3.YuumiAmp(3)
	yuumi3.Strategist(5)              // 超载和战略分析差不多，但战略有前排
	yuumi3.Blue().Jeweled().Breaker() // 破防 帽子 大天使都一样
	yuumi3.End(16)
	yuumi3.Fight("yuumi")

	// 战士
	fiddkesticks2 := C.Mag(60, 0, 0.7).
		Skill(50, 360*3).
		Swing(30)
	fiddkesticks2.Fight("Fiddkesticks[2]")

	fiddkesticks3 := C.Mag(90, 0, 0.7).
		Skill(50, 575*3).
		Swing(30)
	fiddkesticks3.BoomBots(2)
	fiddkesticks3.Techie(2)
	fiddkesticks3.Beaten()
	fiddkesticks3.Jeweled().Justice().Sword()
	fiddkesticks3.Fight("Fiddkesticks[3]")

	//varus2 := C.Mag(53, 15, 0.7).
	//	Skill(75, 300+150*3)
	//varus2.Fight("Varus[2]")

	varus3 := C.Mag(79, 15, 0.7).
		Skill(75, 480+240*3)
	varus3.Executioner(4)
	varus3.Swing(13) // 05:31:10 05:32:40
	varus3.Shojin().Breaker().Staff()
	varus3.Fight("Varus")

	//moderkaiser := C.Mag(75, 30, 0.6).
	//	Skill(80, 300*2).Fury(2, C.TimeGoA)
	//moderkaiser.Fight("Moderkaiser")

	jarven4 := C.Phy(113, 30, 0.65).
		Skill(90, 2*1.2*113)
	jarven4.Slayer(4)
	jarven4.GoldenOx(4)
	jarven4.Fight("Jarven")
	C.FormatResult()
}

func TestSingle3Cost(t *testing.T) {
	C.Level(0)
	draven3 := C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(4)
	draven3.Cypher(3)
	draven3.Breaker().Whisper().Infinity()
	draven3.Fight("draven3.Guinsoo().Whisper().Infinity()")
}

func TestItemsDraven(t *testing.T) {
	C.SortAnalyze()
	draven3 := C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Guinsoo().Whisper().Infinity()
	draven3.Fight("draven3.Guinsoo().Whisper().Infinity()")

	draven3 = C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Shojin().Whisper().Infinity()
	draven3.Fight("draven3.Shojin().Whisper().Infinity()")

	draven3 = C.Phy(117, 30, 0.75).
		Skill(120, 2*3.9*117*(1+0.8+0.64), 2*80*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven3.Rapidfire(2)
	draven3.Cypher(3)
	draven3.Breaker().Whisper().Infinity()
	draven3.Fight("draven3.Breaker().Whisper().Infinity()")
	C.FormatResult()
}

func TestItemsJinx(t *testing.T) {
	C.SortAnalyze()
	jinx3 := C.Phy(124, 0, 0.7).
		Skill(50, 1.5*124*5).
		Stack(C.CastStk, 1.5*124).
		Swing(20)
	jinx3.StreetDemon(7)
	jinx3.Marksman(2)
	jinx3.Shojin().Infinity().Whisper()
	jinx3.Fight("jinx3.Shojin().Infinity().Whisper()")

	jinx3 = C.Phy(124, 0, 0.7).
		Skill(50, 1.5*124*5).
		Stack(C.CastStk, 1.5*124).
		Swing(20)
	jinx3.StreetDemon(7)
	jinx3.Marksman(2)
	jinx3.Guinsoo().Infinity().Whisper()
	jinx3.Fight("jinx3.Guinsoo().Infinity().Whisper()")

	jinx3 = C.Phy(124, 0, 0.7).
		Skill(50, 1.5*124*5).
		Stack(C.CastStk, 1.5*124).
		Swing(20)
	jinx3.StreetDemon(7)
	jinx3.Marksman(2)
	jinx3.Breaker().Infinity().Whisper()
	jinx3.Fight("jinx3.Breaker().Infinity().Whisper()")
	C.FormatResult()
}

func TestItemElise(t *testing.T) {
	C.SortAnalyze()
	elise3 := C.Mag(90, 0, 0.75).
		Skill(55, 140*8).
		Swing(9) // 11:04:80 11:05:70
	elise3.Dynamo(4)
	elise3.Jeweled().Breaker().Staff()
	elise3.Fight("elise3.Jeweled().Breaker().Staff()")

	elise3 = C.Mag(90, 0, 0.75).
		Skill(55, 140*8).
		Swing(9) // 11:04:80 11:05:70
	elise3.Dynamo(4)
	elise3.Jeweled().DeathCap().Staff()
	elise3.Fight("elise3.Jeweled().DeathCap().Staff()")

	elise3 = C.Mag(90, 0, 0.75).
		Skill(55, 140*8).
		Swing(9) // 11:04:80 11:05:70
	elise3.Dynamo(4)
	elise3.Jeweled().Blue().Staff()
	elise3.Fight("elise3.Jeweled().Blue().Staff()")
	C.FormatResult()
}

func TestItemsyuumi(t *testing.T) {
	C.SortAnalyze()
	//yuumi3 := C.Mag(79, 0, 0.75).
	//	Skill(30, (240+480)/2).
	//	Swing(5)
	//yuumi3.yuumiAmp(5)
	//yuumi3.Strategist(3)
	//yuumi3.Blue().Jeweled().Staff()
	//yuumi3.Fight("yuumi[3]")

	yuumi3 := C.Mag(79, 0, 0.75).
		Skill(30, (240+480)/2).
		Swing(5)
	yuumi3.YuumiAmp(3)
	yuumi3.Strategist(5)
	yuumi3.Blue().Jeweled().Staff()
	yuumi3.Fight("yuumi1")

	yuumi3 = C.Mag(79, 0, 0.75).
		Skill(30, (240+480)/2).
		Swing(5)
	yuumi3.YuumiAmp(3)
	yuumi3.Strategist(5)
	yuumi3.Blue().Jeweled().DeathCap()
	yuumi3.Fight("yuumi2")

	yuumi3 = C.Mag(79, 0, 0.75).
		Skill(30, (240+480)/2).
		Swing(5)
	yuumi3.YuumiAmp(3)
	yuumi3.Strategist(5)
	yuumi3.Blue().Jeweled().Breaker()
	yuumi3.Fight("yuumi3")
	C.FormatResult()
}

func TestItemsVarus(t *testing.T) {
	C.SortAnalyze()
	varus3 := C.Mag(79, 15, 0.7).
		Skill(75, 480+240*3)
	varus3.Executioner(4)
	varus3.Swing(13) // 05:31:10 05:32:40
	varus3.Shojin().Breaker().Staff()
	varus3.Fight("varus3.Shojin().Breaker().Staff()")

	varus3 = C.Mag(79, 15, 0.7).
		Skill(75, 480+240*3)
	varus3.Executioner(4)
	varus3.Swing(13) // 05:31:10 05:32:40
	varus3.Shojin().Breaker().Nashor()
	varus3.Fight("arus3.Shojin().Breaker().Nashor()")
	C.FormatResult()
}
