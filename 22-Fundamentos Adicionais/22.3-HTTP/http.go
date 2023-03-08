package main

import (
	"log"
	"net/http"
)

// HTTP é um protocolo de comunicação - Base da comunicação WEB
// Cliente (Faz a requisição) - Servidor (Processa Requisição e envia a resposta)
// Request - Response
// Rotas
// URI - Identificador do Recurso
// Metodo - Get, Post, Put, Delete

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func main() {

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
