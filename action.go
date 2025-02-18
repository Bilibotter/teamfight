package tft

type action int

const (
	TimeGoA     action = iota // 每秒触发一次
	AttackA                   // 每次攻击触发
	BeforeCastA               // 每次施法前触发
	AfterCastA                // 每次施法后触发
)

type event struct {
	action
	num int
}

func newE(act action, num int) event {
	return event{
		action: act,
		num:    num,
	}
}

func (e event) match(act action) bool {
	return e.action == act
}

func (g *ground) process(e event) {
	for _, handler := range g.handlers {
		handler.process(e)
	}
}
