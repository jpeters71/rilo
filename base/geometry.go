package base

import (
	"bozosonparade/rilo/utils"
	"math"
)

// Point type defines a logic point on the screen.  Used for drawing, hit tests, etc.
type Point struct {
	X int
	Y int
}

// Distance calculates the distance between 2 points.
func (pt *Point) Distance(pt2 *Point) int {
	return utils.Round(math.Sqrt((math.Pow(float64(pt2.X-pt.X), 2) + (math.Pow(float64(pt2.Y-pt.Y), 2)))))
}

// Rectangle type defines a logic geometric rectangle.
type Rectangle struct {
	Pt1 Point
	Pt2 Point
}

// Area returns the area of the rectangle
func (rc *Rectangle) Area() int {
	return utils.Round(math.Abs(float64(rc.Pt2.X-rc.Pt1.X) * float64(rc.Pt2.Y-rc.Pt1.Y)))
}

// HitTest returns true if the point lies within the rectangle, false otherwise.
// It's a little more complicated than you might think since the rectangle might between
// reversed; meaning Pt1 may be greater than Pt2.  Again, I'm throughly annoyed golang
// doesn't really support integer math.
func (rc *Rectangle) HitTest(pt Point) bool {
	return (float64(pt.X) <= math.Max(float64(rc.Pt1.X), float64(rc.Pt2.X))) &&
		(float64(pt.X) >= math.Min(float64(rc.Pt1.X), float64(rc.Pt2.X))) &&
		(float64(pt.Y) <= math.Max(float64(rc.Pt1.Y), float64(rc.Pt2.Y))) &&
		(float64(pt.Y) >= math.Min(float64(rc.Pt1.Y), float64(rc.Pt2.Y)))
}

// Polygon type defines a logical polygon.
type Polygon struct {
	Pts []Point
}

// Area returns the area of the polygon
func (poly *Polygon) Area() float64 {
	// Obtained from http://paulbourke.net/geometry/polygonmesh/source1.c
	// Independently varified.
	var area float64

	for i := range poly.Pts {
		j := (i + 1) % len(poly.Pts)
		area += float64(poly.Pts[i].X) * float64(poly.Pts[j].Y)
		area -= float64(poly.Pts[i].Y) * float64(poly.Pts[j].X)
	}
	area /= 2
	return math.Abs(area)
}

// HitTest checks to see if the specified point lies within polygon.  This
// algorthim
func (poly *Polygon) HitTest(pt Point) bool {
	// The winding number
	var wn int

	// Loop through all the edges of the polygon.
	for i, polyPt := range poly.Pts {
		var nextPt Point

		// Make sure that the next point is always valid.  Once we get to the
		// end of the array, loop back around to the beginning for i+1
		if i == (len(poly.Pts) - 1) {
			nextPt = poly.Pts[0]
		} else {
			nextPt = poly.Pts[i+1]
		}
		// Edge from polyPt to nextPolyPt
		if polyPt.Y <= pt.Y {
			// Start y <= target point
			if nextPt.Y > pt.Y {
				// This represents an upward crossing
				if isLeft(polyPt, nextPt, pt) > 0 {
					// pt is left of edge...we have a valid up intersect
					wn++
				}
			}
		} else {
			// Start y > target point (no test needed)
			if nextPt.Y <= pt.Y {
				// This is a downward crossing
				if isLeft(polyPt, nextPt, pt) < 0 {
					// pt is right of edge...we have a valid down intersect
					wn--
				}
			}
		}
	}
	// The winding is zero if the point lies outside the polygon.
	return (wn != 0)
}

func isLeft(pt0 Point, pt1 Point, pt2 Point) int {
	return ((pt1.X-pt0.X)*(pt2.Y-pt0.Y) - (pt2.X-pt0.X)*(pt1.Y-pt0.Y))
}
