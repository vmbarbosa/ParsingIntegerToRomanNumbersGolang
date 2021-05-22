package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gorilla/mux"
)

func Router2() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/api/conversor/{id}", conversor).Methods("GET")
	return router
}

func Test_conversor(t *testing.T) {
	//Test 1

	request, _ := http.NewRequest("GET", "/api/conversor/35", nil)
	response := httptest.NewRecorder()
	Router2().ServeHTTP(response, request)
	respuesta, _ := ioutil.ReadAll(response.Body)

	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(string(respuesta), " ")

	if replace != `"XXXV" ` {
		t.Errorf("Test 1 - getTasks() = %v, want %v", replace, `"XXXV" `)
	}

	//Test 2

	request, _ = http.NewRequest("GET", "/api/conversor/p", nil)
	response = httptest.NewRecorder()
	Router2().ServeHTTP(response, request)
	respuesta, _ = ioutil.ReadAll(response.Body)

	if string(respuesta) != "Numero Invalido" {
		t.Errorf("Test 2 - getTasks() = %v, want %v", response.Body, "Numero Invalido")
	}
}

func Test_indexRoute(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	request, _ := http.NewRequest("GET", "", nil)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: request,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexRoute(tt.args.w, tt.args.r)
		})
	}
}
