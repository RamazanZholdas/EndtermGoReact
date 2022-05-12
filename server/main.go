package main

import (
	"fmt"
	"net/http"
	"react-golang-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server running on port 9000...")

	http.ListenAndServe(":9000", r)
}
