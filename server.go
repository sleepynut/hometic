package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sleepynut/hometic/logger"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello hometic : I'm Gopher!!")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(logger.Middleware)

	r.Handle("/pair-device", PairDeviceHandler(NewCreatePairDevice(db))).Methods(http.MethodPost)

	addr := fmt.Sprintf("%s:%s", host(), os.Getenv("PORT"))
	fmt.Println("addr:", addr)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("starting...")
	log.Fatal(server.ListenAndServe())
}

func host() string {
	h := os.Getenv("HOST")
	if h == "" {
		return "0.0.0.0"
	}

	return h
}

type Pair struct {
	DeviceID int64
	UserID   int64
}

func PairDeviceHandler(device Device) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.L(r.Context()).Info("pair-device")

		var p Pair
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		defer r.Body.Close()
		fmt.Printf("pair: %#v\n", p)

		err = device.Pair(p)
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{"status":"active"}`))
	}
}

type Device interface {
	Pair(p Pair) error
}

type CreatePairDeviceFunc func(p Pair) error

func (fn CreatePairDeviceFunc) Pair(p Pair) error {
	return fn(p)
}

func NewCreatePairDevice(db *sql.DB) CreatePairDeviceFunc {
	return func(p Pair) error {
		_, err := db.Exec("INSERT INTO pairs VALUES ($1,$2);", p.DeviceID, p.UserID)
		return err
	}
}
