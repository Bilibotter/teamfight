package tft

import (
	"sync/atomic"
	"testing"
)

var current int64 = 0

func reset() {
	atomic.StoreInt64(&current, 0)
}

func TestAtkTimes(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseSpeed = 2
	g.fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}

	reset()
	g = newGround()
	g.baseSpeed = 0.5
	g.fight0()
	if g.atkTimes != g.endTime/2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}
}

func TestBuff0_0(t *testing.T) {
	reset()
	Level(0)
	p := buffPassiveT(AttackA, 3, AP(10))
	p.freq = 5
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
	if atomic.LoadInt64(&current) != 2 {
		t.Errorf("current = %d, want 2", atomic.LoadInt64(&current))
	}
}

func TestBuff0_1(t *testing.T) {
	reset()
	Level(0)
	p := buffPassiveT(AttackA, 3, AP(10))
	p.freq = 5
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 2
	g.addPassive(p)
	g.fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}

	if atomic.LoadInt64(&current) != 8 {
		t.Errorf("current = %d, want 8", atomic.LoadInt64(&current))
	}
}

func TestBuff1_0(t *testing.T) {
	reset()
	Level(0)
	p := buffPassiveT(TimeGoA, 3, AP(10))
	p.freq = 5
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
	if atomic.LoadInt64(&current) != 4 {
		t.Errorf("current = %d, want 20", atomic.LoadInt64(&current))
	}
}

func TestBuff1_1(t *testing.T) {
	reset()
	Level(3)
	p := buffPassiveT(TimeGoA, 3, AP(10))
	p.freq = 5
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 2
	g.addPassive(p)
	g.fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
	if atomic.LoadInt64(&current) != 4 {
		t.Errorf("current = %d, want 4", atomic.LoadInt64(&current))
	}
}

func TestAttr0_0(t *testing.T) {
	reset()
	Level(0)
	p := attrPassive(TimeGoA, AP(10))
	p.left = 5
	p.right = 30
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
}

func TestAttr0_1(t *testing.T) {
	reset()
	Level(0)
	p := attrPassive(TimeGoA, AD(10))
	p.left = 5
	p.right = 10
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 1.0
	g.baseAtk = 100
	g.addPassive(p)
	g.fight0()
	if g.physics()-1.0 > 1e-6 {
		t.Errorf("ability = %f, want 1.0", g.physics())
	}
	all, avg, record := g.getAtkRecord()
	if all != 2255 || avg != 112 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestStack0_0(t *testing.T) {
	reset()
	Level(0)
	p := stackPassive(TimeGoA, 1, AP(10))
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if p.overlay.count != 20 {
		t.Errorf("count = %d, want 20", p.overlay.count)
	}
	if g.ability() != 300 {
		t.Errorf("ability = %d, want 300", g.ability())
	}
}

func TestStack0_1(t *testing.T) {
	reset()
	Level(0)
	p := stackPassive0(TimeGoA, 3, 1, AP(10))
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if p.overlay.count != 3 {
		t.Errorf("count = %d, want 3", p.overlay.count)
	}
	if g.ability() != 130 {
		t.Errorf("ability = %d, want 130", g.ability())
	}
}

func TestStack0_2(t *testing.T) {
	reset()
	Level(0)
	p := stackPassive(AttackA, 1, AP(10))
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if p.overlay.count != 10 {
		t.Errorf("count = %d, want 10", p.overlay.count)
	}
	if g.ability() != 200 {
		t.Errorf("ability = %d, want 200", g.ability())
	}
}

func TestStack0_3(t *testing.T) {
	reset()
	Level(0)
	p := stackPassive(AttackA, 2, AP(10))
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.addPassive(p)
	g.fight0()
	if p.overlay.count != 5 {
		t.Errorf("count = %d, want 10", p.overlay.count)
	}
	if g.ability() != 150 {
		t.Errorf("ability = %d, want 200", g.ability())
	}
}

func TestOncePassive(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseAtk = 100
	g.endTime = 20
	g.baseSpeed = 0.5
	g.OncePassive(AttackA, 5, AD(100))
	g.fight0()
	if g.physics()-2.0 > 1e-6 {
		t.Errorf("ability = %f, want 2.0", g.physics())
	}
	all, avg, record := g.getAtkRecord()
	if all != 1650 || avg != 165 || len(record) != 10 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestMinionPassive(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.baseAtk = 100
	g.endTime = 20
	g.baseSpeed = 1.0
	g.Minion(BeforeCastA, 1, 6, AtkAmp(30))
	g.Skill(40, 0).Swing(0)
	g.fight0()
	all, avg, record := g.getAtkRecord()
	if all != 2926 || avg != 146 || len(record) != 20 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func TestDmgPassive(t *testing.T) {
	reset()
	Level(0)
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 1.0
	g.DmgPassive(AttackA, 2, 100)
	g.fight0()
	all, avg, record := g.getSecRecord()
	if all != 1000 || avg != 50 || len(record) != 21 {
		t.Errorf("Wrong.%d, %d, %v", all, avg, record)
	}
}

func buffPassiveT(trigger action, remain int, a ...*attrs) *passive {
	f := addBuff(remain, a...)
	call := func(g *ground) {
		f(g)
		atomic.AddInt64(&current, 1)
	}
	p := &passive{
		trigger: trigger,
		call:    call,
	}
	return p
}
