// injection
// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"
// 	_ "github.com/lib/pq"
// )

// type Pair struct {
// 	DeviceID int
// 	UserID   int
// }

// type PairDeviceHandler struct {
// 	createPairDevice CreatePairDevice
// }

// func (ph *PairDeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	var p Pair
// 	err := json.NewDecoder(r.Body).Decode(&p)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}
// 	defer r.Body.Close()
// 	fmt.Printf("pair: %#v\n", p)

// 	err = ph.createPairDevice(p)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	w.Write([]byte(`{"status":"active"}`))
// }

// type CreatePairDevice func(p Pair) error

// func createPairDevice(p Pair) error {
// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec("INSERT INTO pairs VALUES (?,?);", p.DeviceID, p.UserID)
// 	return err
// }

// func main() {
// 	fmt.Println("hello hometic : I'm Gopher!!")

// 	r := mux.NewRouter()
// 	r.Handle("/pair-device", &PairDeviceHandler{createPairDevice}).Methods(http.MethodPost)

// 	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
// 	fmt.Println("addr:", addr)

// 	server := http.Server{
// 		Addr:    addr,
// 		Handler: r,
// 	}

// 	log.Println("starting...")
// 	log.Fatal(server.ListenAndServe())
// }

//================================================================================
// closure function: higher-order function
// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	fmt.Println("hello hometic : I'm Gopher!!")

// 	r := mux.NewRouter()
// 	r.Handle("/pair-device", PairDeviceHandler(createPairDevice)).Methods(http.MethodPost)

// 	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
// 	fmt.Println("addr:", addr)

// 	server := http.Server{
// 		Addr:    addr,
// 		Handler: r,
// 	}

// 	log.Println("starting...")
// 	log.Fatal(server.ListenAndServe())
// }

// type Pair struct {
// 	DeviceID int64
// 	UserID   int64
// }

// func PairDeviceHandler(createPairDevice CreatePairDeviceFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var p Pair
// 		err := json.NewDecoder(r.Body).Decode(&p)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			json.NewEncoder(w).Encode(err.Error())
// 			return
// 		}
// 		defer r.Body.Close()
// 		fmt.Printf("pair: %#v\n", p)

// 		err = createPairDevice(p)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(err.Error())
// 			return
// 		}

// 		w.Write([]byte(`{"status":"active"}`))
// 	}
// }

// type CreatePairDeviceFunc func(p Pair) error

// func createPairDevice(p Pair) error {
// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec("INSERT INTO pairs VALUES (?,?);", p.DeviceID, p.UserID)
// 	return err
// }

//================================================================================
// interface

// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	fmt.Println("hello hometic : I'm Gopher!!")

// 	r := mux.NewRouter()
// 	r.Handle("/pair-device", PairDeviceHandler(createPairDevice{})).Methods(http.MethodPost)

// 	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
// 	fmt.Println("addr:", addr)

// 	server := http.Server{
// 		Addr:    addr,
// 		Handler: r,
// 	}

// 	log.Println("starting...")
// 	log.Fatal(server.ListenAndServe())
// }

// type Pair struct {
// 	DeviceID int64
// 	UserID   int64
// }

// func PairDeviceHandler(device Device) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var p Pair
// 		err := json.NewDecoder(r.Body).Decode(&p)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			json.NewEncoder(w).Encode(err.Error())
// 			return
// 		}
// 		defer r.Body.Close()
// 		fmt.Printf("pair: %#v\n", p)

// 		err = device.Pair(p)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(err.Error())
// 			return
// 		}

// 		w.Write([]byte(`{"status":"active"}`))
// 	}
// }

// type Device interface {
// 	Pair(p Pair) error
// }

// type createPairDevice struct {
// }

// func (createPairDevice) Pair(p Pair) error {
// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec("INSERT INTO pairs VALUES ($1,$2);", p.DeviceID, p.UserID)
// 	return err
// }

//================================================================================
// handler

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello hometic : I'm Gopher!!")

	r := mux.NewRouter()
	r.Handle("/pair-device", PairDeviceHandler(CreatePairDeviceFunc(createPairDevice))).Methods(http.MethodPost)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("addr:", addr)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("starting...")
	log.Fatal(server.ListenAndServe())
}

type Pair struct {
	DeviceID int64
	UserID   int64
}

func PairDeviceHandler(device Device) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Pair
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		defer r.Body.Close()
		fmt.Printf("pair: %#v\n", p)

		err = device.Pair(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

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

func createPairDevice(p Pair) error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO pairs VALUES ($1,$2);", p.DeviceID, p.UserID)
	return err
}
