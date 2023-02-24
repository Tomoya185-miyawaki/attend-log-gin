package main

import (
	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
)

func main() {
	db := db.Init()
	defer db.Close()
}
