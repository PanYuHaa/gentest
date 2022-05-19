package main

import (
	"context"
	"gentest/biz"
)

func main()  {
	ctx := context.Background()

	biz.Create(ctx)
}
