package main

import (
	"log"
	"net/http"
	"time"
)

func parseDateTime(dateTimeStr string) time.Time {
	layout := "2006-01-02T15:04"
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func empty_str(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("DELETE SOMETING ")
	w.Write([]byte("")) // send empty string to the front end
}

func alert(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("it worked somehow SHOW SOMETING ")
	w.Write([]byte(`
	<div id="modal">
		<div class="modal-content">
			<h1>Modal Dialog</h1>
			This is the modal content.
			You can put anything here, like text, or a form, or an image.
			<br>
			<br>
				<button class="content-button" hx-post="/api/empty" hx-trigger="click" hx-swap="outerHTML" hx-target="#modal">
				close
				</button>
		</div>
	</div>
	`))
}

func getDeviceByName(devices []Device_asset, dev_name string) Device_asset {
	for _, device := range devices {
		if device.Name == dev_name {
			return device
		}
	}
	// Return an empty Device_asset if no match is found
	return Device_asset{}
}
