# Repair Scheduler and Archiver Tool
## NOT UPDATED: <<10:11:2023>>

The Repair Scheduler and Archiver Tool is a web application designed to manage the scheduling and archival of repairs for various machines. The tool maintains detailed records of each machine, including its name, model, description, current working state, repair history, and scheduled maintenance. It is built using the Go programming language to ensure high performance and reliability.

### Features
* Machine Management: Maintain a list of machines with detailed information including name, model, and descriptions.
* Current Status: Track whether each machine is in a working state.
* Repair History: Archive past repairs with details such as the nature of the problem, the fix applied, time consumed, and the date of the last repair.
* Cost Tracking: Maintain records of parts and labor costs for each repair.
* Scheduled Repairs: Manage and schedule upcoming repairs to ensure timely maintenance.
  
> Additional Features
* Repair Shops: Maintain information about repair shops, including location and contact details.
* Parts Inventory: Track parts used in repairs, including supplier information and prices.
* Location : Records for the location of repair shops and parts suppliers.


![My Remote Image](https://i.gyazo.com/5e21e01029bd75948cbe0f598371d000.png)

### Technologies Used
* Backend: Go (Golang) for building a robust and efficient server-side application.
* Web Framework: Echo for routing and handling HTTP requests.
* Database: PostgreSQL for storing machine data and repair records.
* Frontend: HTML, CSS, and HTMX for a responsive and interactive user interface.

### 👉 Set Up for `Windows` 

> 👉 **Step 1** - Download the code from the GH repository (using `GIT`) and download mongodb drivers

```bash
$ git clone https://github.com/aleks20905/golang_simple_website.git
$ cd golang_simple_website
```

<br />

> 👉 **Step 2** - Start the APP `localy`

```bash
$ go run .
```
<br />
<br />

### 👉 Set Up for `Linux` 

> 👉 **Step 1** - Download the code from the GH repository (using `GIT`) 

```bash
$ git clone https://github.com/aleks20905/golang_simple_website.git
$ cd golang_simple_website
$ go get go.mongodb.org/mongo-driver/mongo
```
>  **or** use 'gh' insted, but u will need to have gh install
```bash
$ gh repo clone aleks20905/golang_simple_website
$ cd golang_simple_website
$ go get go.mongodb.org/mongo-driver/mongo
```
<br />

> 👉 **Step 2** - Start the APP `localy`

```bash
$ go run *.go
```

Visit `http://localhost:8000` in your browser. The app should be up & running.

