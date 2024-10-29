package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)


type Hello interface {
	SayHello(context.Context, string) (string, error)
}

type hello struct{ weaver.Implements[Hello] }

func (h *hello) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

func main() {
	if err := weaver.Run(context.Background(), func(router *weaver.Router) {
		router.Register(weaver.Registration[Hello]{
			Component: "hello",
			Constructor: func(ctx context.Context) (Hello, error) {
				return &hello{}, nil
			},
		})
	}); err
