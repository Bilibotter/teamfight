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
	C.FormatResult()
	t.Logf("Finish")
}

func TestSingle13(t *testing.T) {
	C.Level(3)
	heimerdinger := C.Mag(60, 0).
		Skill(40, 81*5).Stack(C.CastStk, 81).
		Staff().Blue()
	heimerdinger.Add(C.ManaAmp(50)) // 4先知
	heimerdinger.Add(C.AP(20))      // 科技枪
	heimerdinger.Fight("黑默丁格")

	heimerdinger0 := C.Mag(60, 0).
		Skill(40, 81*5).Stack(C.CastStk, 81).
		Shojin().Nashor().Jeweled() // 只有青龙刀的素材
	heimerdinger0.Add(C.ManaAmp(50)) // 4先知
	heimerdinger0.Fight("黑默丁格0")

	t.Logf("Finish")
}
