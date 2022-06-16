package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	w "weatherBackend/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){
		// Loading the .env in memory
		errEnv := godotenv.Load()
		if errEnv!=nil{
			log.Fatal(errEnv)
		}

    router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/weather/{name}",w.GetWeatherByCityName).Methods("GET")

		port:=os.Getenv("PORT")
		fmt.Println("Server running on port: "+port)
		log.Fatal(http.ListenAndServe(":"+port,router))
}