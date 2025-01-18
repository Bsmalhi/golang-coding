package util

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func profiling() { 
	log.Println("booting on localhost:8080") 
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}