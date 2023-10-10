package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Device_asset struct {
	Name    string
	Idk     int
	Working bool
	Model   string
	//CereatedTime time
	//LatestStatus time
	//LatestRepair time
}

/*
	type ListOfRepairs struct{
		Problem string
		Fix string
		Description string
		StartedRepair time
		EndedRepair time

}
*/
func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("html/base.html", "html/main_content.html", "html/left_side.html"))
		id := r.URL.Query().Get("id")
		fmt.Println("id =>", id)
		deviceassets := map[string][]Device_asset{
			"deviceassets": {
				{Name: "Кутер", Idk: 10, Working: true, Model: "mazda"},
				{Name: "Голяма бъркалка", Idk: 15, Working: true},
			},
		}

		tmpl.ExecuteTemplate(w, "base", deviceassets)
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
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	//http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
