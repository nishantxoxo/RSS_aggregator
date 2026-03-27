package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("hello")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("port empty")
	}


	router:= chi.NewRouter()
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300},
	),
	)



	serv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}


	log.Printf("Server startiung on Port %v", portString)
 	err :=	serv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	

	// fmt.Println("PORT- ", portString)
}