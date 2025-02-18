package tft

import "sort"

var results []result

var disableOutput = false

type result struct {
	output string
	dps    int
}

func SortAnalyze() {
	disableOutput = true
	outputLevel = 0
}

func FormatResult() {
	sort.Slice(results, func(i, j int) bool {
		return results[i].dps > results[j].dps
	})
	for _, res := range results {
		print(res.output)
	}
}
