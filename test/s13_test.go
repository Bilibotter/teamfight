package test

import (
	"testing"
	C "tft"
)

func Test13(t *testing.T) {
	C.SortAnalyze()
	heimerdinger := C.Mag(60, 0).
		Skill(40, 81*5).Stack(C.CastStk, 81).
		Staff().Blue()
	heimerdinger.Add(C.ManaAmp(50)) // 4先知
	heimerdinger.Add(C.AP(20))      // 科技枪
	heimerdinger.Fight("黑默丁格")

	heimerdinger0 := C.Mag(60, 0).
		Skill(40, 81*5).Stack(C.CastStk, 81).
		Invoker()
	heimerdinger0.Add(C.ManaAmp(50)) // 4先知
	heimerdinger0.Fight("黑默丁格0")

	camille := C.Phy(117, 0).
		Skill(25, 3*117, 120).
		Swing(10)
	camille.Add(C.Critable())
	camille.Add(C.CritRate(40)) // 4伏击
	camille.Add(C.CritAmp(30))
	camille.Add(C.AMP(12))
	camille.Sword().Justice().Titan().Beaten()
	camille.Fight("卡密尔")

	camille0 := C.Phy(117, 0).
		Skill(25, 3*117, 120).
		Swing(10)
	camille0.Add(C.Critable())
	camille0.Add(C.CritRate(55)) // 5伏击
	camille0.Add(C.CritAmp(35))
	camille0.Add(C.AMP(12))
	camille0.Sword().Justice().Titan().Beaten()
	camille0.Fight("卡密尔0")
	C.FormatResult()
	t.Logf("Finish")
}

func TestSingle13(t *testing.T) {
	C.Level(3)
	camille := C.Phy(117, 0).
		Skill(25, 3*117, 120).
		Swing(10)
	camille.Add(C.Critable())
	camille.Add(C.CritRate(55))
	camille.Add(C.CritAmp(35))
	camille.Add(C.AMP(12))
	camille.Sword().Justice().Titan().Beaten()
	camille.Fight("卡密尔")
	//heimerdinger := C.Mag(60, 0).
	//	Skill(40, 81*5).Stack(C.CastStk, 81).
	//	Staff().Blue()
	//heimerdinger.Add(C.ManaAmp(50)) // 4先知
	//heimerdinger.Add(C.AP(20))      // 科技枪
	//heimerdinger.Fight("黑默丁格")
	//
	//heimerdinger0 := C.Mag(60, 0).
	//	Skill(40, 81*5).Stack(C.CastStk, 81).
	//	Shojin().Nashor().Jeweled() // 只有青龙刀的素材
	//heimerdinger0.Add(C.ManaAmp(50)) // 4先知
	//heimerdinger0.Fight("黑默丁格0")

	t.Logf("Finish")
}
