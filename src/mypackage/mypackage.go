package mypackage

import "context"

//go:generate mockery --name MyInterface --outpkg mypackage_mocks --with-expecter=true
type MyInterface interface {
	MyMethod(context.Context, string) (int, error)
}

type MyComponent struct {
	i MyInterface
}

func (c *MyComponent) MyMethod(ctx context.Context, s string) (int, error) {
	return c.i.MyMethod(ctx, s)
}

func NewMyComponent(i MyInterface) *MyComponent {
	return &MyComponent{i: i}
}
