package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articulo", articulosTodos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequest()
}

type Articulo struct {
	Titulo    string `json : "titulo"`
	Desc      string `json : "desc"`
	Contenido string `json:"contenido"`
}

type ArticulosList []Articulo

func articulosTodos(w http.ResponseWriter, r *http.Request) {
	articulos := ArticulosList{
		Articulo{Titulo: "Encendedor", Desc: "medidas 1 x 4cms ,Color Verde", Contenido: "Plastico"},
	}
	fmt.Println("Endpoint Hit: Todos los Articulo ")
	json.NewEncoder(w).Encode(articulos)

}
