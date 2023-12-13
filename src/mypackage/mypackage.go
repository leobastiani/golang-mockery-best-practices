package mypackage

import "context"

//go:generate mockery --name MyInterface --outpkg mypackage_mocks --with-expecter=true
type MyInterface[T any] interface {
	MyMethod(context.Context, T) (int, error)
}

type MyComponent struct {
	i MyInterface[string]
}

func (c *MyComponent) MyMethod(ctx context.Context, s string) (int, error) {
	return c.i.MyMethod(ctx, s)
}

func NewMyComponent(i MyInterface[string]) *MyComponent {
	return &MyComponent{i: i}
}
