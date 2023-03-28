package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD  - Creat - Read - Update - Delete
	// Creat - POST / Read - GET / UPDATE - PUT / Delete - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods("GET")
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods("DELETE")

	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
