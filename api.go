package main

import (
	"log"
	"net/http"
	"sync"
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
func updateCacheDevices(cacheDbData *[]Device_asset, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Fetch data from the database
		newData := mongoGetAllDevices()

		// Update the cache with fresh data
		*cacheDbData = newData

		// Sleep for some time before the next update
		time.Sleep(5 * time.Second)
	}
}

func updateCacheShops(cacheDbData *[]Shops, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Fetch data from the database
		newData := mongoGetAllShops()

		// Update the cache with fresh data
		*cacheDbData = newData

		// Sleep for some time before the next update
		// select with normal 5 sec or when sending new dato to be reloaded ???????????? #TODO
		time.Sleep(5 * time.Second)
	}
}
