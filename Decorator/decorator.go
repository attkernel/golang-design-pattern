package Decorator

type Component interface {
	Cal() int
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Cal() int {
	return 0
}

type AddWarp struct {
	Component
	Num int
}

type MulWarp struct {
	Component
	Num int
}

func (a *AddWarp) Cal() int {
	return a.Component.Cal() + a.Num
}

func (m *MulWarp) Cal() int {
	return m.Component.Cal() * m.Num
}

func NewAddWarp(c Component, n int) Component {
	return &AddWarp{
		Component: c,
		Num:       n,
	}
}

func NewMulWarp(c Component, n int) Component {
	return &MulWarp{
		Component: c,
		Num:       n,
	}
}
