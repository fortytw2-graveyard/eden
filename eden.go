package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fortytw2/eden/web"
	_ "github.com/joho/godotenv/autoload"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", web.Homepage)

	log.Println("eden: now listening on port", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), httpLogger(router))
	if err != nil {
		panic(err)
	}
}

// cleanly log all HTTP requests
func httpLogger(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		router.ServeHTTP(w, req)
		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)
		log.Println(req.Method, req.URL, elapsedTime)
	})
}
