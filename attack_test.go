package tft

import "testing"

func TestAttack(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400 || avg != 220 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400 || avg != 110 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}

	reset()
	g = newGround()
	g.baseSpeed = 0.5
	g.baseAtk = 100
	g.Fight0()
	if g.atkTimes != g.endTime/2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record = g.getSecRecord()
	if all != 1100 || avg != 55 || len(record) != 21 {
		t.Errorf("Wrong sec0.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 1100 || avg != 110 || len(record) != 10 {
		t.Errorf("Wrong atk0.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong cast0.%d, %d, %v", all, avg, record)
	}
}

func TestEffect(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.ae = 330
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400*4 || avg != 220*4 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400*4 || avg != 110*4 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestEffectAD(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.ae = 330
	g.aeType = 1
	g.ad = 200
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400*8 || avg != 220*8 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400*8 || avg != 110*8 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestEffectAp(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.ae = 330
	g.ap = 200
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400*7 || avg != 220*7 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400*7 || avg != 110*7 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestEffectAmp(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.ae = 330
	g.amp = 200
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400*8 || avg != 220*8 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400*8 || avg != 110*8 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestEffectCrit(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.ae = 300
	g.skillCrit = true
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400*4 || avg != 220*4 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400*4 || avg != 110*4 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestEffectDoubleExplode(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.baseAtk = 100
	g.ae = 110
	g.asAmp = 200
	g.atkAmp = 200
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
	all, avg, record := g.getSecRecord()
	if all != 4400*6 || avg != 220*6 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getAtkRecord()
	if all != 4400*6 || avg != 110*6 || len(record) != 40 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
	all, avg, record = g.getCastRecord()
	if all != 0 || avg != 0 || len(record) != 0 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}
