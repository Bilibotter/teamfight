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
	g.Fight0()
	if g.atkTimes != g.endTime*2 {
		t.Errorf("atkTimes = %d, want 20", g.atkTimes)
	}
	if g.castTimes != 0 {
		t.Errorf("castTimes = %d, want 0", g.castTimes)
	}

	reset()
	g = newGround()
	g.baseSpeed = 0.5
	g.Fight0()
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
	g.passive(p)
	g.Fight0()
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
	g.passive(p)
	g.Fight0()
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
	g.passive(p)
	g.Fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
	if atomic.LoadInt64(&current) != 4 {
		t.Errorf("current = %d, want 20", atomic.LoadInt64(&current))
	}
}

func TestBuff1_1(t *testing.T) {
	reset()
	Level(0)
	p := buffPassiveT(TimeGoA, 3, AP(10))
	p.freq = 5
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 2
	g.passive(p)
	g.Fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
	if atomic.LoadInt64(&current) != 4 {
		t.Errorf("current = %d, want 20", atomic.LoadInt64(&current))
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
	g.passive(p)
	g.Fight0()
	if g.ability() != 110 {
		t.Errorf("ability = %d, want 110", g.ability())
	}
}

func TestAttr0_1(t *testing.T) {
	reset()
	Level(0)
	p := attrPassive(TimeGoA, AP(10))
	p.left = 5
	p.right = 10
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.passive(p)
	g.Fight0()
	if g.ability() != 100 {
		t.Errorf("ability = %d, want 100", g.ability())
	}
}

func TestStack0_0(t *testing.T) {
	reset()
	Level(0)
	p := stackPassive(TimeGoA, 1, AP(10))
	g := newGround()
	g.endTime = 20
	g.baseSpeed = 0.5
	g.passive(p)
	g.Fight0()
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
	g.passive(p)
	g.Fight0()
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
	g.passive(p)
	g.Fight0()
	if p.overlay.count != 10 {
		t.Errorf("count = %d, want 10", p.overlay.count)
	}
	if g.ability() != 200 {
		t.Errorf("ability = %d, want 200", g.ability())
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
