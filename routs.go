package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/base.html", "html/main_content.html", "html/left_side.html"))

	id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
	//fmt.Println("id =>", id) //prints the ID from the URL

	//deviceName := "Example Device" // Replace with the desired device name
	mainStructs := mongoGetAllData()
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
		DeviceAsset:       foundDevice,
	}

	tmpl.ExecuteTemplate(w, "base", data)
}

func h2(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/base.html", "html/create_new_device.html", "html/left_side.html"))

	//id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
	//fmt.Println("id =>", id) //prints the ID from the URL

	mainStructs := mongoGetAllData()

	data := PageData{
		DeviceAssetsNames: mainStructs,
	}

	tmpl.ExecuteTemplate(w, "base", data)
}
func idk(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("html/base.html", "html/create_new_device.html", "html/left_side.html"))
	fmt.Println("asd")
	//id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
	if len(r.FormValue("devname")) > 1 { // #checks if the deviceName is longer than 2 characters
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
		mongoSendData(device) //sending the data to the db
	}

	http.Redirect(w, r, "/dev", http.StatusSeeOther)
}
func editDevice(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/base.html", "html/edit_device.html", "html/left_side.html"))

	id := r.URL.Query().Get("id") // !!! getting the ID from the website URL
	fmt.Println("id =>", id)      //prints the ID from the URL

	mainStructs := mongoGetAllData()

	foundDevice := getDeviceByName(mainStructs, id)

	data := PageData{
		DeviceAssetsNames: mainStructs,
		DeviceAsset:       foundDevice,
	}

	tmpl.ExecuteTemplate(w, "base", data)
}
