package mypackage

type MyInterface interface {
	MyMethod() int
}

type MyComponent struct {
	i MyInterface
}

func (c *MyComponent) MyMethod() int {
	return c.i.MyMethod()
}

func NewMyComponent(i MyInterface) *MyComponent {
	return &MyComponent{i: i}
}
