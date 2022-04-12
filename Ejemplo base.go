//Ejemplo base

package main //instancio el main

import (
	"fmt"
	"ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json "id"`
	Name string `json:"name"`
}

func main() {

	cats, err := GetCategories("MLA")
	if err != nil {
		//validar
	}
	fmt.Println("Las categorias son ......")
}

func GetCategories(siteID string) (Categories, error) { //Dividos la categorias del json

	reponse := http.Get // Hay q completar

	bytes := ioutil.ReadAll(reponse.bytes) //vompletar

	var categories Categories
	err := json.Unmarshall(bytes, &cats)

	return cats, nil

}
