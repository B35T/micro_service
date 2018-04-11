package router

import (
	"net/http"
	"micro_service/controllers"
)

func Start() {
	http.HandleFunc("/register", controllers.Register)
	println("server running to port 3000")
	http.ListenAndServe(":3000", nil)
}
