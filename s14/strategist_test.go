package s14

import (
	"testing"
	C "tft"
)

func TestStrategist(t *testing.T) {
	C.SortAnalyze()
	ziggs := C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360)/4+150*0.8).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(5)
	ziggs.Shojin().Nashor().Jeweled() // 大天使、羊刀也行
	ziggs.Fight("Ziggs")

	urgot := C.Phy(90, 0, 0.8).
		Skill(50, 0).
		Fury(5, C.TimeGoA, C.AtkAmp(200), C.AS(125), C.AE(13)).
		Swing(0)
	urgot.DeathBlade().Whisper().Guinsoo() // 破防可以换羊刀或轻语 泰坦可能也厉害
	urgot.Executioner(2)
	//urgot.BoomBots(2)
	urgot.Strategist0(5)
	urgot.Fight("Urgot")

	brand := C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(3)
	brand.Techie(2)
	brand.Strategist(5)
	brand.Shojin().Nashor().Jeweled() // 纳什也可以，帽子/巨杀差很多
	brand.Fight("Brand+")

	brand = C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(3)
	brand.Techie(2)
	brand.Strategist0(5)
	brand.Shojin().Nashor().Jeweled() // 纳什也可以，帽子/巨杀差很多
	brand.Fight("Brand")

	aurora := C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Strategist0(5)
	aurora.Shojin().Nashor().Jeweled()
	aurora.Fight("Aurora")

	aurora = C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Strategist(5)
	aurora.Shojin().Nashor().Jeweled()
	aurora.Fight("Aurora+")

	// 火球数受A.M.P影响
	// 提伯斯普攻伤害大约为坠落伤害的2/3
	amp := 3
	annie := C.Mag(45, 0, 0.75).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 405).Loop()
	annie.Minion(C.BeforeCastA, 4, 8, C.AE(405*2/3))
	annie.Blue().Nashor().Jeweled()
	annie.Strategist(5)
	annie.Fight("Annie+")

	annie = C.Mag(45, 0, 0.75).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 405).Loop()
	annie.Minion(C.BeforeCastA, 4, 8, C.AE(405*2/3))
	annie.Blue().Nashor().Jeweled()
	annie.Strategist0(5)
	annie.Fight("Annie")

	//reduce := (2 + 0.8*3) / 5
	//samira := C.Phy(90, 20, 0.8).
	//	Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
	//	Swing(25)
	//samira.Strategist0(5)
	//samira.Night().Infinity().Shojin()
	//samira.End(17)
	//samira.Fight("Samira")
	//
	//samira = C.Phy(90, 20, 0.8).
	//	Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
	//	Swing(25)
	//samira.Strategist(5)
	//samira.Night().Infinity().Shojin()
	//samira.End(17)
	//samira.Fight("Samira+")
	reduce := (2 + 0.8*3) / 5
	samira := C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
		Swing(25)
	samira.Strategist0(5)
	samira.Guinsoo().Infinity().Shojin()
	samira.End(17)
	samira.Fight("Samira")

	samira = C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
		Swing(25)
	samira.Strategist(5)
	samira.Guinsoo().Infinity().Shojin()
	samira.End(17)
	samira.Fight("Samira+")

	samira = C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
		Swing(25)
	samira.Strategist0(5)
	samira.Shojin().Infinity().Shojin()
	samira.End(17)
	samira.Fight("Samira*")

	samira = C.Phy(90, 20, 0.8).
		Skill(100, 0.8*90*(25+5*C.Overload(amp))*reduce+90*1.8*5).
		Swing(25)
	samira.Strategist(5)
	samira.Shojin().Infinity().Shojin()
	samira.End(17)
	samira.Fight("Samira*+")
	C.FormatResult()
}

func TestTestStrategistRun(t *testing.T) {
	C.SortAnalyze()
	// 留存时间受AP影响，33AP加一秒(0.03AP/秒)
	// 分身应该不能叠羊刀和触发电容
	zeri := C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	zeri.Minion(C.BeforeCastA, 1, 5, C.AtkAmp(50))
	zeri.Guinsoo().Whisper().Infinity()
	zeri.Rapidfire(2)
	zeri.Strategist0(5)
	zeri.Fight("Zeri")

	aphelios := C.Phy(95, 0, 0.75).
		Skill(60, 2.5*3*95, 30*3).
		Fury(6, C.AttackA, C.AtkAmp(4*7+3*7))
	aphelios.Strategist0(5)
	aphelios.Marksman(2)
	aphelios.Guinsoo().Whisper().Infinity()
	aphelios.Fight("Aphelios")

	xayah := C.Phy(83, 25, 0.8).
		Skill(75, 1.6*83*6, 30*6).
		Swing(16) // 01:21:40 01:23:00
	xayah.Shojin().Infinity().Whisper() // 优先青龙刀，羊刀也还行
	xayah.Strategist0(5)
	xayah.Marksman(2)
	xayah.Fight("Xayah")

	zed := C.Phy(75, 0, 0.85).
		Skill(40, 2.5*75*1.5+1.5*75*1, 5, 90).
		Stack(C.CastStk, 1.5*75*1.5).
		Swing(15) // 03:40:50 03:42:00
	zed.Beaten()
	zed.Guinsoo().Whisper().Infinity()
	zed.Slayer(2)
	zed.Strategist0(5)
	zed.Fight("Zed")

	zed = C.Phy(75, 0, 0.85).
		Skill(40, 2.5*75*1.5+1.5*75*1, 5, 90).
		Stack(C.CastStk, 1.5*75*1.5).
		Swing(15) // 03:40:50 03:42:00
	zed.Beaten()
	zed.Guinsoo().Whisper().Infinity()
	zed.Cypher(3)
	zed.Strategist0(5)
	zed.Fight("Zed*")

	fortune := C.Phy(75, 50, 0.75).
		Skill(150, 0.6*75*(1+3*0.75)*8, 12*(1+4*0.75)*8).
		Swing(20)
	fortune.Strategist0(5)
	fortune.Shojin().Infinity().Whisper()
	fortune.Fight("Fortune")

	draven2 := C.Phy(78, 30, 0.75).
		Skill(120, 2*3.3*78*(1+0.8+0.64), 2*45*(1+0.8+0.64)).
		Swing(12) // 03:40:80 03:42:00
	draven2.Strategist0(5)
	draven2.Rapidfire(2)
	draven2.Guinsoo().Whisper().Infinity()
	draven2.Fight("Draven")
	C.FormatResult()
}
