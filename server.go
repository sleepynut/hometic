package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Pair struct {
	DeviceID int
	UserID   int
}

func main() {
	fmt.Println("hello hometic : I'm Gpher!!")

	r := mux.NewRouter()
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("addr:", addr)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("...starting...")
	log.Fatal(server.ListenAndServe())
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var p Pair
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	defer r.Body.Close()
	fmt.Printf("pair: %#v\n", p)

	// w.Write([]byte(`{"status":"active"}`))
}
