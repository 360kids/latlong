/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package latlong

import "testing"

func TestLookupLatLong(t *testing.T) {
	cases := []struct {
		lat, long float64
		want      string
	}{
		{37.7833, -122.4167, "America/Los_Angeles"},
		{33.924565, -118.431329, "America/Los_Angeles"},
		{24.414390, 121.319733, "Asia/Taipei"},
		{39.980125, 116.483461, "Asia/Shanghai"},
		{33.944655, -103.052396, "America/Denver"},
		{33.900382, -102.911706, "America/Chicago"},
		{51.6891362, 7.7865954, "Europe/Berlin"},
		{59.9182647,10.747593, "Europe/Oslo"},
	}
	for _, tt := range cases {
		if got := LookupZoneName(tt.lat, tt.long); got != tt.want {
			t.Errorf("LookupZoneName(%v, %v) = %q; want %q", tt.lat, tt.long, got, tt.want)
		}
	}
}

var testAllPixels func(t *testing.T)

func TestAllPixels(t *testing.T) {
	if testAllPixels == nil {
		t.Skip("exhaustive pixel test disabled without --tags=latlong_gen (requires extra deps)")
	}
	testAllPixels(t)
}

func TestNewTileKey(t *testing.T) {
	cases := []struct {
		size, x, y int
	}{
		{0, 1<<14 - 1, 1<<14 - 1},
		{0, 1<<14 - 1, 0},
		{0, 0, 1<<14 - 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 2, 3},
		{2, 3, 1},
		{3, 3, 3},
	}
	for i, tt := range cases {
		tk := newTileKey(byte(tt.size), uint16(tt.x), uint16(tt.y))
		if tk.size() != byte(tt.size) {
			t.Errorf("%d. size = %d; want %d", i, tk.size(), tt.size)
		}
		if tk.x() != uint16(tt.x) {
			t.Errorf("%d. x = %d; want %d", i, tk.x(), tt.x)
		}
		if tk.y() != uint16(tt.y) {
			t.Errorf("%d. y = %d; want %d", i, tk.y(), tt.y)
		}
	}
}
