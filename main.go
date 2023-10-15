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
	DeviceAssetsNames []Device_asset
	DeviceAssets      []Device_asset
	Smt               []Someting
}

func getDeviceByName(devices []Device_asset, name string) Device_asset {
	for _, device := range devices {
		if device.Name == name {
			return device
		}
	}
	// Return an empty Device_asset if no match is found
	return Device_asset{}
}

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("html/base.html", "html/main_content.html", "html/left_side.html"))

		// TO DO retrive form DB and sent the right data from 'id'
		deviceAssets := []Device_asset{
			{Name: "Example Device", Idk: time.Now(), Working: true, Model: "ABC123", CreatedTime: time.Now(), LatestRepair: time.Now(), ScheduledRepair: time.Now(),
				RepairList: ListOfRepairs{Problem: "Broken Screen", Fix: "Replace Screen", Description: "The device's screen is cracked.", StartedRepair: time.Now(), EndedRepair: time.Now()},
			},
			{Name: "Голяма бъркалка", Idk: time.Now(), Working: true},
		}

		somethings := []Someting{
			{Problems: "Кутер", Fix: "true", Idk: "asd"},
		}

		id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
		//fmt.Println("id =>", id) //prints the ID from the URL

		//deviceName := "Example Device" // Replace with the desired device name
		foundDevice := getDeviceByName(deviceAssets, id)

		if foundDevice.Name != "" { // just easy DEBUG...
			// Found the device, use foundDevice for further processing
			fmt.Println("Found device:", foundDevice.Name)
			fmt.Println("Model:", foundDevice)
			// Add more fields as needed
		} else {
			fmt.Println("Device not found")
		}

		data := PageData{
			DeviceAssetsNames: deviceAssets,
			DeviceAssets:      []Device_asset{foundDevice},
			Smt:               somethings,
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
