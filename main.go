package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

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

		somethings := []Someting{
			{Problems: "Кутер", Fix: "true", Idk: "asd"},
		}

		id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
		//fmt.Println("id =>", id) //prints the ID from the URL

		//deviceName := "Example Device" // Replace with the desired device name
		mainStructs := getRes()
		foundDevice := getDeviceByName(mainStructs, id)

		/* if foundDevice.Name != "" { // just easy DEBUG...
			// Found the device, use foundDevice for further processing
			fmt.Println("Found device:", foundDevice.Name)
			fmt.Println("Model:", foundDevice)
			// Add more fields as needed
		} else {
			fmt.Println("Device not found")
		} */
		data := PageData{
			DeviceAssetsNames: mainStructs,
			DeviceAssets:      []Device_asset{foundDevice},
			Smt:               somethings,
		}

		tmpl.ExecuteTemplate(w, "base", data)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("html/base.html", "html/create_new_device.html", "html/left_side.html"))

		// TO DO retrive form DB and sent the right data from 'id'

		somethings := []Someting{
			{Problems: "Кутер", Fix: "true", Idk: "asd"},
		}

		id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
		//fmt.Println("id =>", id) //prints the ID from the URL

		//deviceName := "Example Device" // Replace with the desired device name
		mainStructs := getRes()
		foundDevice := getDeviceByName(mainStructs, id)

		/* if foundDevice.Name != "" { // just easy DEBUG...
			// Found the device, use foundDevice for further processing
			fmt.Println("Found device:", foundDevice.Name)
			fmt.Println("Model:", foundDevice)
			// Add more fields as needed
		} else {
			fmt.Println("Device not found")
		} */

		data := PageData{
			DeviceAssetsNames: mainStructs,
			DeviceAssets:      []Device_asset{foundDevice},
			Smt:               somethings,
		}

		tmpl.ExecuteTemplate(w, "base", data)
	}

	idk := func(w http.ResponseWriter, r *http.Request) {
		//tmpl := template.Must(template.ParseFiles("html/base.html", "html/create_new_device.html", "html/left_side.html"))

		//id := r.URL.Query().Get("id") // !!! getting the ID from the website URL

		device := Device_asset{
			Name:        r.FormValue("devname"),
			Model:       r.FormValue("devmodel"),
			Description: r.FormValue("description"),
			Working:     r.FormValue("working") == "on",
			RepairList: ListOfRepairs{
				Problem:       r.FormValue("problem"),
				Fix:           r.FormValue("fix"),
				Description:   r.FormValue("repairDescription"),
				StartedRepair: parseDateTime(r.FormValue("startedRepair")),
				EndedRepair:   parseDateTime(r.FormValue("endedRepair")),
			},
			LatestRepair: parseDateTime(r.FormValue("latestRepair")),
			ScheduledRepair: scheduledRepair{
				DateOfRepair:     parseDateTime(r.FormValue("dateOfRepair")),
				AddedDescription: r.FormValue("addedDescription"),
			},
			CreatedTime: time.Now(),
			LastUpdated: time.Now(),
		}

		sendData(device) //sending data to the db

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/createNewDev/", h2)
	http.HandleFunc("/submit/", idk)

	//http.HandleFunc("/add-film/", h2)

	//define handlers for web-resurces
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) //from where to be accest in the browser, accest(repeat), whats the dir for the css file
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	log.Fatal(http.ListenAndServe(":8000", nil))

}
func parseDateTime(dateTimeStr string) time.Time {
	layout := "2006-01-02T15:04"
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
