package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db     = map[string]interface{}{}
	dbLock sync.Mutex
)

//Entry is a map entry,fits responses and requests
type Entry struct {
	Key   string      ` json:"key"`
	Value interface{} `json:"value"`
}

//sendResponse sends Response
func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v-%s", entry, err)
	}
}

//dbPostHandler Post Key Value
func dbPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	db[entry.Key] = entry.Value
	sendResponse(entry, w)
}

//dbGetHandler GetValue by Key
func dbGetHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Path[4:] //trim /db
	dbLock.Lock()
	defer dbLock.Unlock()
	value, ok := db[key]
	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found", key), http.StatusNotFound)
		return
	}
	entry := &Entry{
		Key:   key,
		Value: value,
	}
	sendResponse(entry, w)
}

func main() {
	defer recover()

	http.HandleFunc("/db", dbPostHandler)
	http.HandleFunc("/db/", dbGetHandler)
	if err := http.ListenAndServe(":3310", nil); err != nil {
		log.Fatal(err)
	}
}
