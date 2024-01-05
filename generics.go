package main

type Constraint interface {
	Something
}

type Something struct{}

type API[T Constraint] interface {
	Foo(t T) error
}

type Client[T Constraint] struct {
	a string
	b string
}

func NewClient[T Constraint](opts ...func(*Client[T])) API[T] {
	c := Client[T]{}

	for _, o := range opts {
		o(&c)
	}

	return c
}

func (c Client[T]) Foo(t T) error {
	return nil
}

func WithA[T Constraint](a string) func(*Client[T]) {
	return func(c *Client[T]) {
		c.a = a
	}
}

func WithB[T Constraint](b string) func(*Client[T]) {
	return func(c *Client[T]) {
		c.a = b
	}
}
