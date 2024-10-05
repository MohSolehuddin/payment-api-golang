package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Activity struct {
	Timestamp time.Time `json:"timestamp"`
	Activity  string    `json:"activity"`
}

type ActivityTransfer struct {
	Timestamp time.Time `json:"timestamp"`
	Activity  string    `json:"activity"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Nominal   string    `json:"nominal"`
}

// CreateHistory mencatat aktivitas ke dalam file JSON.
func CreateHistory(filename string, activity string, transfer *ActivityTransfer) {
	// Mencatat aktivitas umum
	activityLog := Activity{
		Timestamp: time.Now(),
		Activity:  activity,
	}

	var activities []Activity
	var transferActivities []ActivityTransfer

	// Membaca data dari file
	historyData, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			historyData = []byte(`{"data": []}`) // Jika file tidak ada, buat array kosong
		} else {
			log.Println("Failed to load file " + filename)
			return
		}
	}

	// Mengonversi data JSON yang dibaca ke dalam slice
	if err := json.Unmarshal(historyData, &activities); err != nil {
		log.Println("Failed to unmarshal activities: " + err.Error())
		return
	}

	// Jika transfer tidak nil, tambahkan ke transferActivities
	if transfer != nil {
		transfer.Timestamp = time.Now()
		transferActivities = append(transferActivities, *transfer)
	}

	// Menambahkan aktivitas baru
	activities = append(activities, activityLog)

	// Mengonversi kembali ke format JSON
	newHistoryData, err := json.Marshal(activities)
	if err != nil {
		log.Println("Failed to marshal activities: " + err.Error())
		return
	}

	// Menyimpan kembali ke file
	if err := ioutil.WriteFile(filename, newHistoryData, 0644); err != nil {
		log.Println("Failed to write to file " + filename)
	}
}