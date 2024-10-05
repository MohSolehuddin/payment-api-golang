package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Activity struct {
	Username string `json:"username"`
	Timestamp time.Time `json:"timestamp"`
	Activity  string    `json:"activity"`
}

type ActivityTransfer struct {
	Username string `json:"username"`
	Timestamp time.Time `json:"timestamp"`
	Activity  string    `json:"activity"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Nominal   string    `json:"nominal"`
}

type HistoryData struct {
	Activities        []Activity        `json:"activities"`
	TransferActivities []ActivityTransfer `json:"transfer_activities"`
}

func CreateHistory(filename string, activity string, transfer *ActivityTransfer) {
	activityLog := Activity{
		Timestamp: time.Now(),
		Activity:  activity,
	}

	var history HistoryData

	historyData, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			historyData = []byte(`{"activities": [], "transfer_activities": []}`)
		} else {
			log.Println("Failed to load file " + filename)
			return
		}
	}

	if err := json.Unmarshal(historyData, &history); err != nil {
		log.Println("Failed to unmarshal history data: " + err.Error())
		return
	}

	if transfer != nil {
		transfer.Timestamp = time.Now()
		history.TransferActivities = append(history.TransferActivities, *transfer)
	}

	history.Activities = append(history.Activities, activityLog)

	newHistoryData, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		log.Println("Failed to marshal history data: " + err.Error())
		return
	}

	if err := ioutil.WriteFile(filename, newHistoryData, 0644); err != nil {
		log.Println("Failed to write to file " + filename)
	}
}
