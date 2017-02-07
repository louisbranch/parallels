package camera

import (
	"testing"

	"github.com/luizbranco/parallels/math"
)

func TestCameraClip(t *testing.T) {
	testCases := []struct {
		cam    Camera
		width  int
		height int
		start  int
		w      int
		h      int
	}{
		{},
		{
			cam: Camera{
				Rect: math.Rect{
					X: 0, Y: 0, W: 640, H: 480,
				},
				TileSize: 200,
			},
			width:  6,
			height: 4,
			w:      4,
			h:      3,
			start:  0,
		},
		{
			cam: Camera{
				Rect: math.Rect{
					X: 0, Y: 0, W: 640, H: 480,
				},
				TileSize: 50,
			},
			width:  6,
			height: 4,
			w:      6,
			h:      4,
			start:  0,
		},
		{
			cam: Camera{
				Rect: math.Rect{
					X: 250, Y: 250, W: 1024, H: 768,
				},
				TileSize: 200,
			},
			width:  13,
			height: 10,
			w:      6,
			h:      4,
			start:  14,
		},
	}

	for _, tc := range testCases {
		start, w, h := tc.cam.Clip(tc.width, tc.height)

		if tc.start != start {
			t.Errorf("expect start = %d, got %d", tc.start, start)
		}
		if tc.w != w {
			t.Errorf("expect width = %d, got %d", tc.w, w)
		}
		if tc.h != h {
			t.Errorf("expect height = %d, got %d", tc.h, h)
		}
	}
}
