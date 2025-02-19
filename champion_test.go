package tft

import "testing"

func TestSingle(t *testing.T) {
	Level(3)
	//xerath0 := C.Mag(90, 0, 0.8).
	//	Skill(120, 375*10).
	//	Shojin().Staff().Jeweled()
	//xerath0.Fight("泽拉斯0")
	//fiora := C.Phy(113, 70, 1).
	//	Skill(130, 1.8*113*(3+3+2)/3+3.5*113, 60).
	//	Fighter().Beaten(4) // s13才改动法力值
	//fiora.Add(C.AMP(15))
	//fiora.Fight("菲奥娜")
	innerTesting = false
	kassadin := Mag(101, 30, 0.75).
		Skill(70, 255).Fury(1, TimeGoA)
	kassadin.StackPassive(AfterCastA, AE(80))
	kassadin.Guinsoo().Silver().Sword()
	kassadin.strike = 70
	kassadin.endTime = 20
	kassadin.beaten(9)
	kassadin.walk = true
	kassadin.endTime = 18
	kassadin.Fight("卡萨丁")
	innerTesting = true
}
