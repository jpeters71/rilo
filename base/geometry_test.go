package base

import "testing"

// Test for point distance calcs.
func TestPointDistance(t *testing.T) {
	var tests = []struct {
		pt1  Point
		pt2  Point
		want int
	}{
		{Point{55, 28}, Point{90, 123}, 101},
		{Point{200, 45}, Point{5, 175}, 234},
		{Point{153, 167}, Point{29, 150}, 125},
	}

	for _, test := range tests {
		if got := test.pt1.Distance(&test.pt2); got != test.want {
			t.Errorf("Distance(%v, %v) = %v", test.pt1, test.pt2, got)
		}
	}
}

// Tests the rectangle area function
func TestRectangleArea(t *testing.T) {
	var tests = []struct {
		rc   Rectangle
		want int
	}{
		{Rectangle{Point{5, 16}, Point{78, 224}}, 15184},
		{Rectangle{Point{45, 23}, Point{5, 450}}, 17080},
		{Rectangle{Point{978, 59}, Point{542, 982734}}, 428446300},
	}

	for _, test := range tests {
		if got := test.rc.Area(); got != test.want {
			t.Errorf("Rectangle.Area(%v) = %v", test.rc, got)
		}
	}
}

// Tests rectangle hit tests
func TestRectangleHitTest(t *testing.T) {
	var tests = []struct {
		rc   Rectangle
		pt   Point
		want bool
	}{
		{Rectangle{Point{5, 16}, Point{78, 224}}, Point{66, 200}, true},
		{Rectangle{Point{5, 16}, Point{78, 224}}, Point{66, 225}, false},
		{Rectangle{Point{100, 100}, Point{10, 10}}, Point{11, 11}, true},
		{Rectangle{Point{100, 100}, Point{10, 10}}, Point{10, 10}, true},
		{Rectangle{Point{100, 100}, Point{10, 10}}, Point{9, 10}, false},
	}

	for _, test := range tests {
		if got := test.rc.HitTest(test.pt); got != test.want {
			t.Errorf("Rectangle.HitTest(%v, %v) = %v", test.rc, test.pt, got)
		}
	}
}

// Tests the area of a given polygon.  Not trival...should come up with other test cases eventually.
func TestPolygonArea(t *testing.T) {
	var tests = []struct {
		poly Polygon
		want float64
	}{
		{Polygon{[]Point{{1, 1}, {4, 1}, {4, 3}}}, 3.0},
		{Polygon{[]Point{{4, 6}, {4, -4}, {8, -4}, {8, -8}, {-4, -8}, {-4, 6}}}, 128.0},
	}

	for _, test := range tests {
		if got := test.poly.Area(); got != test.want {
			t.Errorf("Polygon.Area(%v) = %v", test.poly, got)
		}
	}
}

//
func TestPolygonHitTest(t *testing.T) {
	var tests = []struct {
		poly Polygon
		pt   Point
		want bool
	}{
		{Polygon{[]Point{{1, 1}, {4, 1}, {4, 3}}}, Point{3, 2}, true},
		{Polygon{[]Point{{1, 1}, {4, 1}, {4, 3}}}, Point{1, 1}, true},
		{Polygon{[]Point{{1, 1}, {4, 1}, {4, 3}}}, Point{5, 2}, false},
	}

	for _, test := range tests {
		if got := test.poly.HitTest(test.pt); got != test.want {
			t.Errorf("Polygon.HitTest(%v, %v) = %v", test.poly, test.pt, got)
		}
	}
}
