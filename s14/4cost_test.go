package s14

import (
	"testing"
	C "tft"
)

func Test4Cost(t *testing.T) {
	C.SortAnalyze()
	// 双版本
	vex := C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten(8) // vex的挨打回蓝可能偏低,同时法力锁时间过长
	vex.Executioner(4)
	vex.Blue().Night().Breaker() // 大天使最优，其次是帽子、法爆
	vex.End(16)
	vex.Fight("Vex<4>")

	vex = C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten(8) // vex的挨打回蓝可能偏低,同时法力锁时间过长
	vex.Executioner(5)
	vex.Blue().Night().Breaker() // 大天使最优，其次是帽子、法爆
	vex.End(16)
	vex.Fight("Vex<5>")

	// 留存时间受AP影响，33AP加一秒(0.03AP/秒)
	// 分身应该不能叠羊刀和触发电容
	zeri := C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	zeri.Minion(C.BeforeCastA, 1, 5, C.AtkAmp(50))
	zeri.DoubleGuinsoo().DeathBlade()
	zeri.Rapidfire(2)
	zeri.Fight("Zeri")

	// 非时间buff存在bug待修复
	aphelios := C.Phy(95, 0, 0.75).
		Skill(60, 2.2*3*95, 30*3).
		Fury(6, C.AttackA, C.AtkAmp(4*7+3*7))
	aphelios.GoldenOx(2)
	aphelios.Marksman(4)
	aphelios.Guinsoo().Whisper().Infinity()
	aphelios.Fight("Aphelios")

	ziggs := C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360)/4+150*0.8).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(4)
	ziggs.Shojin().Nashor().Jeweled() // 大天使、羊刀也行
	ziggs.Fight("Ziggs<4>")

	ziggs = C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360)/4+150*0.8).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(5)
	ziggs.Shojin().Nashor().Jeweled() // 大天使、羊刀也行
	ziggs.Fight("Ziggs<5>")

	fortune := C.Phy(75, 50, 0.75).
		Skill(150, 0.6*75*(1+4*0.75)*11, 12*(1+4*0.75)*11).
		Swing(20)
	fortune.Syndicate(5)
	fortune.Dynamo(2)
	fortune.Shojin().Infinity().Whisper()
	fortune.Fight("Fortune|5|")

	fortune = C.Phy(75, 50, 0.75).
		Skill(150, 0.6*75*(1+3*0.75)*11, 12*(1+4*0.75)*11).
		Swing(20)
	fortune.Syndicate(5)
	fortune.Dynamo(2)
	fortune.Shojin().Infinity().Whisper()
	fortune.Fight("Fortune|4|")

	fortune = C.Phy(75, 50, 0.75).
		Skill(150, 0.6*75*(1+3*0.75)*11, 12*(1+4*0.75)*11).
		Swing(0)
	fortune.Syndicate(5)
	fortune.Guinsoo().Infinity().Whisper()
	fortune.Fight("Fortune|4|*")

	xayah := C.Phy(90, 25, 0.8).
		Skill(75, 1.6*90*6, 30*6).
		Swing(16) // 01:21:40 01:23:00
	xayah.Shojin().Infinity().Whisper() // 优先青龙刀，羊刀也还行
	xayah.Animal(3).Marksman(4)
	xayah.Fight("Xayah")

	// 叠加效果需要更明确
	zed := C.Phy(75, 0, 0.85).
		Skill(40, 2.25*75*1.5+2.0*75*1, 5, 90).
		Stack(C.CastStk, 2.0*75*1.5).
		Swing(15) // 03:40:50 03:42:00
	zed.Beaten()
	zed.Slayer(6)
	zed.Guinsoo().Silver().Sword()
	zed.Fight("Zed")

	// 火球数受A.M.P影响
	// 提伯斯普攻伤害大约为坠落伤害的2/3
	amp := 4
	annie := C.Mag(45, 0, 0.75).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 405).Loop()
	tibbers := 300.0 * (0.8 * (float64(amp) + 1)) / 1.28
	annie.Minion(C.BeforeCastA, 4, 8, C.AE(int(tibbers)))
	annie.Blue().Nashor().Jeweled()
	annie.GoldenOx(2)
	annie.Fight("Annie<4>")

	amp = 5
	annie = C.Mag(45, 0, 0.75).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 405).Loop()
	tibbers = 300.0 * (0.8 * (float64(amp) + 1)) / 1.28
	annie.Minion(C.BeforeCastA, 4, 8, C.AE(int(tibbers)))
	annie.Blue().Nashor().Jeweled()
	annie.GoldenOx(2)
	annie.Fight("Annie<5>")

	brand := C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(7)
	brand.Techie(2)
	brand.Shojin().Breaker().Jeweled() // 纳什也可以，帽子/巨杀差很多
	brand.Fight("Brand")

	C.FormatResult()
}

func TestSingle4Cost(t *testing.T) {
	C.Level(3)
	// 火球数受A.M.P影响
	// 提伯斯普攻伤害大约为坠落伤害的2/3
	amp := 4
	annie := C.Mag(45, 0, 0.75).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 405).Loop()
	tibbers := 300.0 * (0.8 * (float64(amp) + 1)) / 1.28 / 1.4
	annie.Minion(C.BeforeCastA, 4, 8, C.AE(int(tibbers)))
	annie.Blue().Nashor().Shojin()
	annie.Strategist0(5)
	annie.Fight("Annie<4>")
}

func TestItemsZeri(t *testing.T) {
	C.SortAnalyze()
	// 留存时间受AP影响，33AP加一秒(0.03AP/秒)
	zeri := C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	zeri.Minion(C.BeforeCastA, 1, 5, C.AtkAmp(50))
	zeri.DoubleGuinsoo().DeathBlade()
	zeri.Rapidfire(2)
	zeri.Fight("zeri.DoubleGuinsoo().DeathBlade()")

	zeri = C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	// 破防加法强
	zeri.Minion(C.BeforeCastA, 1, 5+1, C.AtkAmp(50))
	zeri.DoubleGuinsoo().Breaker()
	zeri.Rapidfire(2)
	zeri.Fight("zeri.DoubleGuinsoo().Breaker()")

	zeri = C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	// 破防加法强
	zeri.Minion(C.BeforeCastA, 1, 5+1, C.AtkAmp(50))
	zeri.Guinsoo().Whisper().GiantSlayer()
	zeri.Rapidfire(2)
	zeri.Fight("zeri.Guinsoo().Whisper().GiantSlayer()")

	zeri = C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	// 破防加法强
	zeri.Minion(C.BeforeCastA, 1, 5, C.AtkAmp(50))
	zeri.Guinsoo().Whisper().DeathBlade()
	zeri.Rapidfire(2)
	zeri.Fight("zeri.Guinsoo().Whisper().DeathBlade()")

	zeri = C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	// 破防加法强
	zeri.Minion(C.BeforeCastA, 1, 5, C.AtkAmp(50))
	zeri.Guinsoo().Whisper().Breaker()
	zeri.Rapidfire(2)
	zeri.Fight("zeri.Guinsoo().Whisper().Breaker()")

	zeri = C.Phy(102, 0, 0.75).
		Skill(40, 0).
		Swing(2) // 02:57:00 02:57:20
	// 破防加法强
	zeri.Minion(C.BeforeCastA, 1, 5, C.AtkAmp(50))
	zeri.Guinsoo().Whisper().Infinity()
	zeri.Rapidfire(2)
	zeri.Fight("zeri.Guinsoo().Whisper().Infinity()")
	C.FormatResult()
}

func TestItemsAphelios(t *testing.T) {
	C.SortAnalyze()
	// 非时间buff存在bug待修复
	aphelios := C.Phy(95, 0, 0.75).
		Skill(60, 2.5*3*95, 30*3).
		Fury(6, C.AttackA, C.AtkAmp(4*7+3*7))
	aphelios.GoldenOx(2)
	aphelios.Marksman(2)
	aphelios.Guinsoo().Whisper().Infinity()
	aphelios.Fight("Aphelios")

	// 非时间buff存在bug待修复
	aphelios = C.Phy(95, 0, 0.75).
		Skill(60, 2.5*3*95, 30*3).
		Fury(6, C.AttackA, C.AtkAmp(4*7+3*7))
	aphelios.GoldenOx(2)
	aphelios.Marksman(4)
	aphelios.Guinsoo().Whisper().Infinity()
	aphelios.Fight("Aphelios+")
	C.FormatResult()
}

func TestItemsVex(t *testing.T) {
	C.SortAnalyze()
	// 双版本
	vex := C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten()
	vex.Executioner(4)
	vex.Blue().Night().Staff()
	vex.Fight("vex.Blue().Night().Staff()")

	// 双版本
	vex = C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten()
	vex.Executioner(4)
	vex.Blue().Night().Breaker()
	vex.Fight("vex.Blue().Night().Breaker()")

	// 双版本
	vex = C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten()
	vex.Executioner(4)
	vex.Blue().Night().DeathCap()
	vex.Fight("vex.Blue().Night().DeathCap()")

	// 双版本
	vex = C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten()
	vex.Executioner(4)
	vex.Blue().Night().Justice()
	vex.Fight("vex.Blue().Night().Justice()")

	// 双版本
	vex = C.Mag(45, 0, 0.8).
		Skill(30, 150*2+375+150*2).
		Swing(15) // 02:40:00 02:41:60
	vex.Beaten()
	vex.Executioner(4)
	vex.Blue().Night().Jeweled()
	vex.Fight("vex.Blue().Night().Jeweled()")
	C.FormatResult()
}

func TestItemsZiggs(t *testing.T) {
	C.SortAnalyze()

	ziggs := C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360+150)/4).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(3)
	ziggs.Shojin().Nashor().Jeweled()
	ziggs.Fight("ziggs.Shojin().Nashor().Jeweled()")

	ziggs = C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360+150)/4).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(3)
	ziggs.Shojin().Blue().Jeweled()
	ziggs.Fight("ziggs.Shojin().Blue().Jeweled()")

	ziggs = C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360+150)/4).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(3)
	ziggs.Shojin().Staff().Jeweled()
	ziggs.Fight("ziggs.Shojin().Staff().Jeweled()")

	ziggs = C.Mag(60, 20, 0.75).
		Skill(70, (4+4+3+2)*(360+150)/4).
		Swing(9) // 03:16:70 03:17:60
	ziggs.Cyberboss(2)
	ziggs.Strategist(3)
	ziggs.Shojin().Guinsoo().Jeweled()
	ziggs.Fight("ziggs.Shojin().Guinsoo().Jeweled()")

	C.FormatResult()
}

func TestItemsFortune(t *testing.T) {
	C.SortAnalyze()

	// 存在大头目效果
	fortune := C.Phy(75, 50, 0.75).
		Skill(150, 0.6*75*(1+4*0.75)*11, 12*(1+4*0.75)*11).
		Swing(20)
	fortune.Syndicate(5)
	fortune.Dynamo(2)
	fortune.Shojin().Infinity().Whisper()
	fortune.Fight("Fortune")

	C.FormatResult()
}

func TestItemsAnnie(t *testing.T) {
	C.SortAnalyze()

	// 叠加效果需要更明确
	// 火球数受A.M.P影响
	// 提伯斯普攻伤害大约为坠落伤害的2/3
	amp := 4
	annie := C.Mag(45, 0, 0.75).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 265+2*60+C.Overload(amp)*60).
		Skill(40, 405).Inspire(8, C.TimeGoA, C.AE(405*3/2)).Loop()
	annie.Blue().Nashor().Jeweled()
	annie.GoldenOx(2)
	annie.Fight("Annie")

	C.FormatResult()
}

func TestItemsBrand(t *testing.T) {
	C.SortAnalyze()
	brand := C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(7)
	brand.Techie(2)
	brand.Shojin().Nashor().Jeweled()
	brand.Fight("brand.Shojin().Nashor().Jeweled()")

	brand = C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(7)
	brand.Techie(2)
	brand.Shojin().Breaker().Jeweled()
	brand.Fight("brand.Shojin().Breaker().Jeweled()")

	brand = C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(7)
	brand.Techie(2)
	brand.Shojin().GiantSlayer().Jeweled()
	brand.Fight("brand.Shojin().GiantSlayer().Jeweled()")

	brand = C.Mag(53, 25, 0.75).
		Skill(75, 360*3+120*4).
		Swing(17) // 04:05:90 04:07:60
	brand.StreetDemon(7)
	brand.Techie(2)
	brand.Shojin().DeathCap().Jeweled()
	brand.Fight("brand.Shojin().DeathCap().Jeweled()")
	C.FormatResult()
}
