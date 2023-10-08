package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("html/base.html", "html/main_content.html", "html/left_side.html"))

		tmpl.ExecuteTemplate(w, "base", 10)
	}

	// handler function #2 - returns the template block with the newly added film, as an HTMX response
	// h2 := func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(1 * time.Second)

	// 	tmpl := template.Must(template.ParseFiles("index.html"))
	// 	tmpl.ExecuteTemplate(w, "base", )
	// }

	// define handlers
	http.HandleFunc("/", h1)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	//http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
