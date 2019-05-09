package main

import (
	"context"
	"fmt"
	"github.com/ruohone/go-context-tp/src/main/entity"
	"time"
)

func main() {
	fmt.Println(",ao")

	lc := entity.LoginContext{}
	lc.Id = "1"
	lc.Code = "8888"
	lc.Phone = "131"

	bk := context.Background()
	c1 := context.WithValue(bk, "k", "v")

	Mk(c1)

	m, err := Login(lc)
	fmt.Println(m, err)
}

func Mk(ctx context.Context) {
	fmt.Println(ctx.Value("k"))
}

func Login(ctx entity.LoginContext) (int, error) {
	fmt.Println(ctx)
	AddUser(entity.SaveLoginContext{MyContext: ctx.MyContext, Code: ctx.Code, CreateDate: time.Now()})
	return 1, nil
}

func AddUser(ctx entity.SaveLoginContext) {
	fmt.Println(ctx)
}
