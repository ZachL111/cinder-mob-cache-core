package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 88, Capacity: 93, Latency: 10, Risk: 14, Weight: 8}
	if got := Score(signal); got != 171 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 80, Capacity: 85, Latency: 16, Risk: 25, Weight: 7}
	if got := Score(signal); got != 41 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 92, Capacity: 101, Latency: 22, Risk: 23, Weight: 5}
	if got := Score(signal); got != 61 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
