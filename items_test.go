package tft

import "testing"

func TestShojin(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Skill(30, 1000)
	g.Swing(0)
	g.Shojin()
	g.mana = 0
	g.fight0()

	all, avg, record := g.getAtkRecord()
	if all != 126*20 || avg != 126 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 1150*10 || avg != 1150 || len(record) != 10 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestGuinsoo(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Guinsoo()
	g.as = 100
	g.fight0()
	times := g.bonusAS() / 5
	all, avg, record := g.getAtkRecord()
	if all != 3630 || avg != 110 || len(record) != times {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestInfinity(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Skill(40, 1000)
	g.Swing(0)
	g.infinity()
	g.mana = 0
	g.ad = 100
	g.fight0()

	all, avg, record := g.getAtkRecord()
	if all != 124*20 || avg != 124 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 1240*5 || avg != 1240 || len(record) != 5 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestStaff(t *testing.T) {
	reset()
	Level(0)
	g := Mag(100, 0, 1)
	g.Skill(50, 1000)
	g.Swing(0)
	g.Staff()
	g.mana = 0
	g.ap = 100
	g.fight0()

	all, avg, record := g.getAtkRecord()
	if all != 110*20 || avg != 110 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 7000 || avg != 1750 || len(record) != 4 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestBlue(t *testing.T) {
	reset()
	Level(0)
	g := Mag(100, 0, 1)
	g.Skill(50, 1000)
	g.Swing(0)
	g.blue()
	g.mana = 10
	g.ap = 100
	g.ad = 100
	g.fight0()

	all, avg, record := g.getAtkRecord()
	if all != 2280 || avg != 114 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 5200 || avg != 1040 || len(record) != 5 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestSilver(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Silver()
	g.mana = 0
	g.ad = 100
	g.fight0()

	_, avg, _ := g.getAtkRecord()
	if avg != 118 {
		t.Errorf("Wrong.%d", avg)
	}
	if g.bonusAS() != 57 {
		t.Errorf("Wrong.%d", g.bonusAS())
	}
}

func TestTitan(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Titan()
	g.mana = 0
	g.ad = 100
	g.fight0()

	ap, ad := g.ap, g.ad
	for _, a := range g.attributes {
		if a.valid() {
			ap += a.body().ap
			ad += a.body().ad
		}
	}

	if ap != 150 {
		t.Errorf("Wrong.%d", ap)
	}

	if ad != 150 {
		t.Errorf("Wrong.%d", ad)
	}
}
