package sensemicroservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	jtltojson "github.com/nicholasvuono/jtl-to-json"
)

func getProtocolLevelResultsList(w http.ResponseWriter, r *http.Request) {
	data := db.readAll("plu")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	checkErr(err)
}

func getBrowserLevelResultsList(w http.ResponseWriter, r *http.Request) {
	data := db.readAll("blu")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	checkErr(err)
}

func addProtocolLevelResult(w http.ResponseWriter, r *http.Request) {
	var res *jtltojson.Result
	b, err := r.GetBody()
	checkErr(err)
	err = json.NewDecoder(b).Decode(res)
	checkErr(err)
	db.write("plu", res)
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Result was added to the database successfully!")
	checkErr(err)
}

func addBrowserLevelResult(w http.ResponseWriter, r *http.Request) {
	var res *jtltojson.Result
	b, err := r.GetBody()
	checkErr(err)
	err = json.NewDecoder(b).Decode(res)
	checkErr(err)
	db.write("blu", res)
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Result was added to the database successfully!")
	checkErr(err)
}

func downloadDatabaseBackup(w http.ResponseWriter, r *http.Request) {
	err := db.backup(w)
	checkErr(err)
}

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/results/plu/list", getProtocolLevelResultsList).Methods("GET")
	r.HandleFunc("/results/blu/list", getBrowserLevelResultsList).Methods("GET")
	r.HandleFunc("/results/plu/add", addProtocolLevelResult).Methods("POST")
	r.HandleFunc("/results/blu/add", addBrowserLevelResult).Methods("POST")
	r.HandleFunc("/backup", downloadDatabaseBackup).Methods("GET")
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":10000", routes()))
}
