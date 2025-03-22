package s14

import (
	"testing"
	C "tft"
)

func TestVeigarPlus(t *testing.T) {
	C.SortAnalyze()
	// 有大头目，真伤算180%的魔法伤害
	veigar3 := C.Mag(79, 0, 0.7).
		Skill(40, (540+150)*(0.67+0.33*1.8)).
		Swing(4) // 06:47:70 06:48:40
	veigar3.Cyberboss(2)
	veigar3.Techie(2)
	veigar3.Strategist0(3)
	veigar3.Shojin().Nashor().Jeweled() // 蓝buff也行
	veigar3.Fight("Veigar")

	leblanc3 := C.Mag(79, 0, 0.75).
		Skill(50, 5*145).
		Stack(C.CastStk, 145).
		Swing(10) // 06:59:30 07:00:30
	leblanc3.Cypher(3)
	leblanc3.Strategist(3)
	leblanc3.Shojin().Nashor().Jeweled() // 法爆也行
	leblanc3.Fight("Leblanc")
	C.FormatResult()
}
