package funcDBserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PhoneNumbers []PhoneNumber

var db *gorm.DB

func NewConnection() {

	var err error
	dsn := "host=localhost user=postgres password=trippi2005 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection filed")
	}

	fmt.Println(db)
	fmt.Println("CONNECTED")

	err = db.Find(&PhoneNumbers).Error

	if err != nil {
		fmt.Println("Errore", err)
	}

	fmt.Println(PhoneNumbers)

}

func HandleRequests() {

	NewConnection()

	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/phoneNumbers", Wrapper(GetPhoneNumbers, EmptyDecoder)).Methods("GET")
	myRoute.HandleFunc("/phoneNumber/{Id}", Wrapper(GetPhoneNumber, GetKey)).Methods("GET")
	myRoute.HandleFunc("/phoneNumbers/delete/{Id}", Wrapper(DeletePhoneNumber, GetKey)).Methods("DELETE")
	myRoute.HandleFunc("/phoneNumbers/create", Wrapper(addPhoneNumber, getBody)).Methods("POST")

	http.Handle("/", myRoute)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func Wrapper(fn func(interface{}) (interface{}, error), dec func(*http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		payload, err := dec(r)
		if err != nil {
			//todo esci
		}

		resp, _ := fn(payload)
		//todo se err ... tornare qualcosa altro

		jsonData, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Scrive il JSON come risposta
		w.Write(jsonData)
	}
}

func GetKey(r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	key := vars["Id"]

	return key, nil
}

func EmptyDecoder(r *http.Request) (interface{}, error) {
	return nil, nil
}
