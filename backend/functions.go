package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func conversor(w http.ResponseWriter, r *http.Request) {
	identifier := mux.Vars(r)
	log.Println(identifier["id"])
	number, err := strconv.Atoi(identifier["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil || number > 3000 || number < 1 {
		log.Println("Numero Invalido")
		fmt.Fprintf(w, "Numero Invalido")
		return
	}

	NumerosRomanos := []string{"M", "CMXC", "CM", "D", "CDXC", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	reference := []int{1000, 990, 900, 500, 490, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	i := 0
	result := ""
	for number >= 0 {
		if number <= 0 || (len(reference) == (i - 1)) {
			break
		}
		j := number
		for j > 0 {
			j = number - reference[i]
			if j >= 0 {
				number -= reference[i]
				result = result + NumerosRomanos[i]
			} else {
				break
			}
		}
		i++
	}
	log.Println(result)
	json.NewEncoder(w).Encode(result)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bienvenido al conversor de Numeros Enteros a Romanos\n\n\nPara poder convertir un numero Entero a Romano se debe acceder mediante una ruta de una API\nLa siguiente ruta es un metodo GET que solo acepta numeros enteros mayores a 0 y menores o iguales a 3000 :\n\n/api/conversor/{id}")
}
