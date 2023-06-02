package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// configurando aplicación
	var app application

	// leer desde la terminal

	// conectar a la base de datos
	app.Domain = "example.com"
	log.Println("Iniciando app en el puerto:", port)
	http.HandleFunc("/", app.Home)

	// iniciar el servidor web
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}