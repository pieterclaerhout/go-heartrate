package main

import (
	"fmt"
	"math"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

type HRZone struct {
	MinPCT int
	MaxPCT int
	Name   string
}

func (hrZone *HRZone) ToHeartRate(rest int, max int) (int, int) {
	lower := hrZone.calcHeartRate(hrZone.MinPCT, rest, max)
	upper := hrZone.calcHeartRate(hrZone.MaxPCT, rest, max)
	return lower, upper
}

func (hrZone *HRZone) calcHeartRate(pct int, rest int, max int) int {
	pctAsFloat := float64(pct) / 100.0
	restAsFloat := float64(rest)
	diffAsFloat := float64(max - rest)
	return int(math.Round((pctAsFloat * diffAsFloat) + restAsFloat))
}

func printZones(restHR int, maxHR int, zones []*HRZone) {

	tw := table.NewWriter()
	tw.SetStyle(table.StyleBold)
	tw.AppendHeader(table.Row{"type", "% min", "% max", "hr min", "hr max"})
	tw.SetAlign([]text.Align{
		text.AlignLeft,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
	})
	for _, zone := range zones {
		lowerHR, upperHR := zone.ToHeartRate(restHR, maxHR)
		tw.AppendRow(table.Row{
			zone.Name,
			fmt.Sprintf("%d%%", zone.MinPCT),
			fmt.Sprintf("%d%%", zone.MaxPCT),
			lowerHR,
			upperHR,
		})
	}
	fmt.Printf("%s\n", tw.Render())
}

func main() {

	restHR := 45
	maxHR := 194

	fmt.Printf("Rest HR: %d\n Max HR: %d\n", restHR, maxHR)

	fmt.Println("\nSysteem Wellens")
	printZones(restHR, maxHR, []*HRZone{
		{45, 60, "Recuperatietraining"},
		{61, 64, "Long Slow Distance"},
		{65, 70, "Extensieve Uithouding"},
		{71, 78, "Tempoduurtraining"},
		{79, 84, "Bloktrainging"},
		{85, 89, "Extensieve Intervaltraining"},
		{90, 100, "Weerstandstraining"},
	})

	fmt.Println("\nhttp://doefiets.nl/training/basisprincipes/hartslagzones-karvonen")
	printZones(restHR, maxHR, []*HRZone{
		{45, 60, "Hersteltraining (H)"},
		{60, 70, "Rustige duurtraining (D1)"},
		{70, 80, "Intensieve duurtraining (D2)"},
		{80, 90, "MLSS-training (D3)"},
		{90, 100, "Weerstandstraining (W)"},
	})

	fmt.Println("\nhttps://www.3athlon.be/2005/09/29/formule-van-karvonen/ (fietsen)")
	printZones(restHR, maxHR, []*HRZone{
		{55, 60, "recuperatie training"},
		{60, 65, "LSD training (Long Slow Distance)"},
		{65, 70, "extensieve uithouding"},
		{70, 80, "intensieve uithouding"},
		{80, 90, "tempo-interval"},
		{90, 95, "Weerstand"},
	})

	fmt.Println("\nhttps://www.3athlon.be/2005/09/29/formule-van-karvonen/ (lopen)")
	printZones(restHR, maxHR, []*HRZone{
		{55, 65, "recuperatie training"},
		{65, 70, "LSD training (Long Slow Distance)"},
		{70, 75, "extensieve uithouding"},
		{75, 85, "intensieve uithouding"},
		{85, 90, "tempo-interval"},
		{90, 95, "Weerstand"},
	})

}
