/*
メインパッケージ
*/
package main

import (
	db "github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure"
	router "github.com/Tomoya185-miyawaki/attend-log-gin/router"
)

func main() {
	db := db.Init()
	defer db.Close()
	r := router.Bind()
	r.Run(":3000")
}
