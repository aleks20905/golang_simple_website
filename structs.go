package main

import (
	"time"
)

type Device_asset struct {
	Name        string
	Model       string
	Description string
	Working     bool

	RepairList      ListOfRepairs
	LatestRepair    time.Time
	ScheduledRepair scheduledRepair

	CreatedTime time.Time
	LastUpdated time.Time
}
type scheduledRepair struct {
	DateOfRepair     time.Time
	AddedDescription string
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
	DeviceAsset       Device_asset
	Smt               []Someting
	Shops             []Shops
	Shop              Shops
}
type Shops struct {
	Name         string
	PhoneNumber  string
	Website      string
	AdditionInfo string
	Review       string
	Address      string
	Color        string
}
