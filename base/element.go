package base

// Element is the base entity for all of Rilo.  Each entity must know how to draw itself
// as well as be able to perform hit tests, etc.
type Element struct {
	id     int
	zOrder int
	show   bool
}

// IElement is the interface that enables Rilo to draw and hit test items.
type IElement interface {
	Draw()
	GetID() int
	GetZOrder() int
	HitTest(pt Point) bool
}

// GetID returns the element's id.
func (e *Element) GetID() int {
	return e.id
}

// Draw draws the element.  The base element implementation is a no-op.
func (e *Element) Draw() {
	// No-op by default
}

// GetZOrder returns the zordering value for this element.
func (e *Element) GetZOrder() int {
	return e.zOrder
}

// HitTest returns wheth
func (e *Element) HitTest(pt Point) bool {
	// Default implementation returns false
	return false
}
