package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go app...")

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
