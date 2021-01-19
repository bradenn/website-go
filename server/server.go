package server

import (
	"fmt"
	"os"
)

func Serve() {
	r := NewRouter()
	err := r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}
}
