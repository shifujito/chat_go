package main

import (
	_ "github.com/lib/pq"
	"github.com/shifujito/chat_go/server/controllers"
)

func main() {
	controllers.Run()
}
