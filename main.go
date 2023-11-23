package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var cacheDbData []Device_asset
var DEBUG bool = false // debug

func main() {
	fmt.Println("Go app...")
	//cache
	var wg sync.WaitGroup
	cacheDbData = mongoGetAllDevices()
	wg.Add(1)
	go updateCache(&cacheDbData, &wg)

	// define handlers
	http.HandleFunc("/dev/", h1)
	http.HandleFunc("/createNewDev/", h2)
	http.HandleFunc("/submit/", idk)
	http.HandleFunc("/edit/", editDevice)
	http.HandleFunc("/api/alert", alert)
	http.HandleFunc("/api/empty", empty_str)
	http.HandleFunc("/shops/", shops)

	//define handlers for web-resurces
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	log.Fatal(http.ListenAndServe(":8000", nil))

}
