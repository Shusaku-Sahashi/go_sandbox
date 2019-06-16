package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("Jhone", "a1234")
}

func ProcessRequest(userId, authToken string) {
	ctx := context.WithValue(context.Background(), "UserID", userId)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleRequest(ctx)
}

func HandleRequest(ctx context.Context) {
	fmt.Println(
		ctx.Value("UserID"),
		ctx.Value("authToken"),
	)
}
