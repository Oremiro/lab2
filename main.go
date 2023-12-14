package main

import (
	"fmt"
	"hl_lab2/pkg/async"
	"time"
)

type User struct {
	ID   int
	Name string
}

func main() {
	fut := new(async.Future[User])

	go func() {
		time.Sleep(200 * time.Millisecond)
		user := User{ID: 1, Name: "Yarick"}

		async.ResolveFuture(fut, user, nil)
	}()

	user, err := fut.Value()

	fmt.Println(user, err)
}
