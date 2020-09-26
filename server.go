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

	r.Handle("/pair-device", CustomHandlerFunc(PairDeviceHandler(NewCreatePairDevice(db)))).Methods(http.MethodPost)

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

type CustomResponseWriter interface {
	JSON(statusCode int, data interface{})
}

type CustomHandlerFunc func(CustomResponseWriter, *http.Request)

func (handler CustomHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler(&JSONResponseWriter{w}, r)
}

type JSONResponseWriter struct {
	http.ResponseWriter
}

func (w *JSONResponseWriter) JSON(statusCode int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func PairDeviceHandler(device Device) func(w CustomResponseWriter, r *http.Request) {
	return func(w CustomResponseWriter, r *http.Request) {
		logger.L(r.Context()).Info("pair-device")

		var p Pair
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.JSON(http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()
		fmt.Printf("pair: %#v\n", p)

		err = device.Pair(p)
		if err != nil {
			w.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		w.JSON(http.StatusOK, []byte(`{"status":"active"}`))
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
