package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Device_asset struct {
	Name            string
	Idk             time.Time
	Working         bool
	Model           string
	RepairList      ListOfRepairs
	CreatedTime     time.Time
	LatestRepair    time.Time
	ScheduledRepair time.Time
}

type ListOfRepairs struct {
	Problem       string
	Fix           string
	Description   string
	StartedRepair time.Time
	EndedRepair   time.Time
}
type Someting struct {
	Problems string
	Fix      string
	Idk      string
}
type PageData struct {
	DeviceAssets []Device_asset
	Smt          []Someting
}

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("html/base.html", "html/main_content.html", "html/left_side.html"))
		id := r.URL.Query().Get("id")
		fmt.Println("id =>", id)

		// TO DO retrive form DB and sent the right data from 'id'
		deviceAssets := []Device_asset{
			{Name: "Example Device", Idk: time.Now(), Working: true, Model: "ABC123", CreatedTime: time.Now(), LatestRepair: time.Now(), ScheduledRepair: time.Now(),
				RepairList: ListOfRepairs{
					Problem: "Broken Screen", Fix: "Replace Screen", Description: "The device's screen is cracked.", StartedRepair: time.Now(), EndedRepair: time.Now()},
			},
			{Name: "Голяма бъркалка", Idk: time.Now(), Working: true},
		}

		somethings := []Someting{
			{Problems: "Кутер", Fix: "true", Idk: "asd"},
		}

		data := PageData{
			DeviceAssets: deviceAssets,
			Smt:          somethings,
		}

		tmpl.ExecuteTemplate(w, "base", data)
	}

	// define handlers
	http.HandleFunc("/", h1)
	//http.HandleFunc("/add-film/", h2)

	//define handlers for web-resurces
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	log.Fatal(http.ListenAndServe(":8000", nil))

}
