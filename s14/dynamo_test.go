package s14

import (
	"testing"
	C "tft"
)

func TestDynamo(t *testing.T) {
	C.SortAnalyze()
	// 双版本
	aurora := C.Mag(75, 20, 0.8).
		Skill(80, 900+175*(6+4+3)/3).
		Swing(20) // 13:70 15:70
	aurora.Dynamo(4)
	aurora.Animal(3)
	aurora.Jeweled().Staff().Breaker()
	aurora.Fight("Aurora")

	// 存在大头目效果
	fortune := C.Phy(75, 50, 0.75).
		Skill(150, 0.6*75*(1+4*0.75)*11, 12*(1+4*0.75)*11).
		Swing(20)
	fortune.Syndicate(3)
	fortune.Dynamo(4)
	fortune.Breaker().Infinity().Whisper()
	fortune.Fight("Fortune")

	C.FormatResult()
}
