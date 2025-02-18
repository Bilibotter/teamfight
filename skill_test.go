package tft

import "testing"

func TestNormalCast(t *testing.T) {
	reset()
	Level(3)
	g := Phy(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 1200 || avg != 150 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestNormalCastApAmp(t *testing.T) {
	reset()
	Level(0)
	g := Mag(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 1200 || avg != 150 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	g = Mag(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.skillCrit = true
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 165*8 || avg != 165 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	g = Mag(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.castAmp = 200
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 300*8 || avg != 300 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	g = Mag(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.ap = 200
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 250*8 || avg != 250 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	g = Mag(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.amp = 200
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 4400*2 || avg != 110*2 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 300*8 || avg != 300 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestNormalCastAd(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.ap = 200
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 200*8 || avg != 200 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	g = Phy(100, 0, 2)
	g.Skill(50, 100, 50)
	g.skill.wait = 0
	g.ad = 200
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 4400*2 || avg != 110*2 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 250*8 || avg != 250 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestStackCastSwing(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Skill(40, 10)
	g.Stack(CastStk, 100)
	g.Swing(10)
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 16*110 || avg != 110 || len(record) != 16 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 640 || avg != 160 || len(record) != 4 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 2)
	g.Skill(40, 10)
	g.Stack(CastStk, 100)
	g.Swing(5)
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 32*110 || avg != 110 || len(record) != 32 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 2880 || avg != 360 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestStackAmp(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Skill(40, 10)
	g.Stack(CastStk, 100)
	g.Swing(10)
	g.amp = 200
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 16*110*2 || avg != 110*2 || len(record) != 16 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 640*2 || avg != 160*2 || len(record) != 4 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 2)
	g.Skill(40, 10)
	g.Stack(CastStk, 100)
	g.Swing(5)
	g.castAmp = 200
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 32*110 || avg != 110 || len(record) != 32 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 2880*2 || avg != 360*2 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 2)
	g.Skill(40, 10)
	g.Stack(CastStk, 100)
	g.Swing(5)
	g.ad = 200
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 32*110*2 || avg != 110*2 || len(record) != 32 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 2880*2 || avg != 360*2 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 2)
	g.Skill(40, 10)
	g.Stack(CastStk, 100)
	g.Swing(5)
	g.skillCrit = true
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 32*110 || avg != 110 || len(record) != 32 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 288*11 || avg != 36*11 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestFury(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Skill(40, 0)
	g.Fury(100, TimeGoA, AD(100))
	g.Swing(0)
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 3960 || avg != 198 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 1 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 2)
	g.Skill(100, 0)
	g.Fury(5, TimeGoA, AD(100))
	g.Swing(0)
	g.endTime = 11
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	// 不包含右边界
	if all != 3410 || avg != 155 || len(record) != 22 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 1 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestInspire(t *testing.T) {
	reset()
	Level(0)
	g := Phy(100, 0, 1)
	g.Skill(40, 0)
	g.Inspire(100, TimeGoA, AD(100))
	g.Swing(0)
	g.Fight0()
	all, avg, record := g.getAtkRecord()
	if all != 3960 || avg != 198 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 5 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 2)
	g.Skill(50, 0)
	g.Inspire(60, TimeGoA, AD(100))
	g.Swing(0)
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 8250 || avg != 206 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 8 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	Level(0)
	g = Phy(100, 0, 1)
	g.Skill(50, 0)
	g.Fury(5, AttackA, AD(100))
	g.Swing(0)
	g.Fight0()
	all, avg, record = g.getAtkRecord()
	if all != 165*20 || avg != 165 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 2 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}
