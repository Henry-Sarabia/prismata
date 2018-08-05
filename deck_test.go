package main

import (
	"reflect"
	"testing"
)

func TestAdvancedSet(t *testing.T) {
	var cases = []struct {
		name string
		d    Deck
		exp  []string
	}{
		{
			"Pass: deck 1",
			Deck{
				Randomizer: [][]string{
					[]string{
						"Xaetron",
						"Corpus",
						"Thermite Core",
						"Drake",
						"Thorium Dynamo",
						"Shadowfang",
						"The Wincer",
						"Nivo Charge",
					},
				},
			},
			[]string{
				"Xaetron",
				"Corpus",
				"Thermite Core",
				"Drake",
				"Thorium Dynamo",
				"Shadowfang",
				"The Wincer",
				"Nivo Charge",
			},
		},
		{
			"Pass: deck 2",
			Deck{
				Randomizer: [][]string{
					[]string{
						"Synthesizer",
						"Valkyrion",
						"Blood Phage",
						"Defense Grid",
						"Infusion Grid",
						"Aegis",
						"Thorium Dynamo",
						"Nivo Charge",
						"Redeemer",
					},
				},
			},
			[]string{
				"Synthesizer",
				"Valkyrion",
				"Blood Phage",
				"Defense Grid",
				"Infusion Grid",
				"Aegis",
				"Thorium Dynamo",
				"Nivo Charge",
				"Redeemer",
			},
		},
		{
			"Pass: deck 3",
			Deck{
				Randomizer: [][]string{
					[]string{
						"Odin",
						"Mobile Animus",
						"Barrier",
						"Perforator",
						"Hannibull",
					},
				},
			},
			[]string{
				"Odin",
				"Mobile Animus",
				"Barrier",
				"Perforator",
				"Hannibull",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			adv := tt.d.AdvancedSet()
			if !reflect.DeepEqual(adv, tt.exp) {
				t.Errorf("got: <%v>, want: <%v>", adv, tt.exp)
			}
		})
	}
}
