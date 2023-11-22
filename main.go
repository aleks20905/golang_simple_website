package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var cacheDbData []Device_asset
var DEBUG bool = false

func main() {
	fmt.Println("Go app...")
	//cache
	var wg sync.WaitGroup
	cacheDbData = mongoGetAllData()
	wg.Add(1)
	go updateCache(&cacheDbData, &wg)

	// define handlers
	http.HandleFunc("/dev/", h1)
	http.HandleFunc("/createNewDev/", h2)
	http.HandleFunc("/submit/", idk)
	http.HandleFunc("/edit/", editDevice)
	http.HandleFunc("/api/alert", alert)
	http.HandleFunc("/api/empty", empty_str)

	//define handlers for web-resurces
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func updateCache(cacheDbData *[]Device_asset, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Fetch data from the database
		newData := mongoGetAllData()

		// Update the cache with fresh data
		*cacheDbData = newData

		// Sleep for some time before the next update
		time.Sleep(5 * time.Second)
	}
}
