package base

type IRiloContext interface {
	DrawRectangle()
	GetID() int
	GetZOrder() int
	HitTest(pt Point) bool
}
