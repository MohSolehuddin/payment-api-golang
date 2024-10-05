package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/MohSolehuddin/payment-api-golang/services"
)

type Saldo struct {
	Timestamp time.Time `json:"timestamp"`
	Username string `json:"username"`
	TotalSaldo int `json:"totalSaldo"`
	Out int `json:"out"`
	In int `json:"in"`
}

type Transfer struct {
	Timestamp time.Time `json:"timestamp"`
	Sender string `json:"sender"`
	To string `json:"to"`
	Nominal   string    `json:"nominal"`
}
type DataSaldo struct {
	Data []Saldo `json:"data"`
}
type CustomerDetailType struct {
	Username string `json:"username"`
	Name string `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Gender string `json:"gender"`
}
type CustomerType struct {
	Data []CustomerDetailType `json:"data"`
}
func PaymentHandler(filename string, t *Transfer)  {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			data = []byte(`{"data": []}`)
		}else {
			log.Println("Failed to load file " + filename)
			return
		}
	}
	var dataSaldo DataSaldo
	CustomerData, customerDataErr := ioutil.ReadFile("/storage/customers.json")
	if customerDataErr != nil {
		if os.IsNotExist(customerDataErr) {
			data = []byte(`{"data": []}`)
		}else {
			log.Println("Failed to load file " + filename)
			return
		}
	}
	var customerData CustomerType

	if err := json.Unmarshal(data, &dataSaldo); err != nil {
		log.Println("Failed to unmarshal saldo data: " + err.Error())
		return
	}

	if customerDataErr := json.Unmarshal(CustomerData, &customerData); err != nil {
		log.Println("Failed to unmarshal customer data: " + customerDataErr.Error())
		return
	}
	var ActivityTransfer services.ActivityTransfer
	ActivityTransfer.From = t.Sender
	ActivityTransfer.To = t.To
	ActivityTransfer.Nominal = t.Nominal
	ActivityTransfer.Timestamp = time.Now()
	services.CreateHistory("/storage/history.json", "Transfer activity recorded", &ActivityTransfer )
}
func LoginHandler()  {
	
}
func LogoutHandler()  {
	
}