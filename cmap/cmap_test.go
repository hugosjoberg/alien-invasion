package cmap

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {

	tests := []struct {
		a        []string
		cityName string
		want     CityNode
	}{
		{[]string{"Foo", "north=Bar", "west=Baz", "south=Qu-ux"}, "Foo", CityNode{North: "Bar", West: "Baz", South: "Qu-ux"}},
		{[]string{"Foo", "", "south=Qu-ux"}, "Foo", CityNode{South: "Qu-ux"}},
		{[]string{"Foo"}, "Foo", CityNode{}},
	}
	for _, tt := range tests {

		testname := fmt.Sprintf("%s", tt.a)
		t.Run(testname, func(t *testing.T) {
			cityName, ans := parseLine(tt.a)
			if cityName != tt.cityName {
				t.Errorf("got %s, want %s", cityName, tt.cityName)
			}
			if ans.North != tt.want.North {
				t.Errorf("got %s, want %s", ans.North, tt.want.North)
			}
			if ans.South != tt.want.South {
				t.Errorf("got %s, want %s", ans.South, tt.want.South)
			}
			if ans.East != tt.want.East {
				t.Errorf("got %s, want %s", ans.East, tt.want.East)
			}
			if ans.West != tt.want.West {
				t.Errorf("got %s, want %s", ans.West, tt.want.West)
			}
		})
	}
}
