package router

import (
	"net/http"
	"micro_service/controllers"
)

func Start() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/test", controllers.TestConnect)
	http.HandleFunc("/", controllers.JsonAPI)
	println("server running to port 3000")
	http.ListenAndServe(":3000", nil)
}
