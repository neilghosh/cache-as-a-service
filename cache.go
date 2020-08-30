package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

type cacheItem struct {
	Key   string
	Value string
	//TTL   int
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case "GET":
		keys, ok := r.URL.Query()["key"]
		var key string

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
		} else {
			key = keys[0]
			//TODO Multi Get
		}

		var response, err = getCache(ctx, key)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}

	case "POST":
		var request cacheItem
		//TODO Strict Checking for Key and Value
		//TODO Multi Set, Array of items
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if putCache(ctx, request) == true {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func getCache(c context.Context, key string) (string, error) {

	item0, err := memcache.Get(c, key)
	if err != nil && err != memcache.ErrCacheMiss {
		//return err
		log.Printf("Listening on port %v", err)

	}
	if err == nil {
		log.Printf("memcache hit: Key=%q Val=[% x]\n", item0.Key, item0.Value)
	} else {
		log.Printf("memcache miss\n")
		return "", errors.New("memcache miss")
	}
	return string(item0.Value), nil
}

func putCache(c context.Context, request cacheItem) bool {
	item1 := &memcache.Item{
		Key:   request.Key,
		Value: []byte(request.Value),
	}

	if err := memcache.Set(c, item1); err != nil {
		log.Printf("Eerror creating key %v", err)
		return false
	}
	log.Printf("Created Item %v", item1)
	return true
}
