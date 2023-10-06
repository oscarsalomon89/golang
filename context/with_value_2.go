package main

import (
	"context"
	"fmt"
)

func enrichContext2(ctx context.Context) context.Context {
	nexCtx := context.WithValue(ctx, "api-key", "my-super-secret-api-key")
	return context.WithValue(nexCtx, "user", "oscar")
}

func doSomethingCools2(ctx context.Context) {
	apiKey := ctx.Value("api-key")
	fmt.Println("ctx.ApiKey: ", apiKey)

	user := ctx.Value("user")
	fmt.Println("ctx.User: ", user)
}

func main88() {
	fmt.Println("--Go Context Example--")
	ctx := context.Background()
	ctx = enrichContext2(ctx)
	doSomethingCools2(ctx)
}
