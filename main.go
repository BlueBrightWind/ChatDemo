package main

import "ChatDemo/router"

func main() {
	app := router.InitRouter()
	app.Run("127.0.0.1:80")
}
