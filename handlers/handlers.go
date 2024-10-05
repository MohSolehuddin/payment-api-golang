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
	Timestamp  time.Time `json:"timestamp"`
	Username   string    `json:"username"`
	TotalSaldo int       `json:"totalSaldo"`
	Out        int       `json:"out"`
	In         int       `json:"in"`
}

type Transfer struct {
	Timestamp time.Time `json:"timestamp"`
	Sender    string    `json:"sender"`
	To        string    `json:"to"`
	Nominal   int       `json:"nominal"`
}

type DataSaldo struct {
	Data []Saldo `json:"data"`
}

type CustomerDetailType struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	Gender    string    `json:"gender"`
}

type CustomerType struct {
	Data []CustomerDetailType `json:"data"`
}

func PaymentHandler(filename string, t *Transfer) {
	// Membaca file saldo
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			data = []byte(`{"data": []}`)
		} else {
			log.Println("Failed to load file " + filename)
			return
		}
	}

	CustomerData, customerDataErr := ioutil.ReadFile("/storage/customers.json")
	if customerDataErr != nil {
		if os.IsNotExist(customerDataErr) {
			CustomerData = []byte(`{"data": []}`)
		} else {
			log.Println("Failed to load customer data " + customerDataErr.Error())
			return
		}
	}

	var dataSaldo DataSaldo
	var customerData CustomerType

	if err := json.Unmarshal(data, &dataSaldo); err != nil {
		log.Println("Failed to unmarshal saldo data: " + err.Error())
		return
	}

	if err := json.Unmarshal(CustomerData, &customerData); err != nil {
		log.Println("Failed to unmarshal customer data: " + err.Error())
		return
	}

	senderValid := false
	receiverValid := false

	for _, customer := range customerData.Data {
		if customer.Username == t.Sender {
			senderValid = true
		}
		if customer.Username == t.To {
			receiverValid = true
		}
	}

	if !senderValid || !receiverValid {
		log.Println("Sender or Receiver not found in customer data")
		return
	}

	for i, saldo := range dataSaldo.Data {
		if saldo.Username == t.Sender {
			if saldo.TotalSaldo < t.Nominal {
				log.Println("Insufficient balance")
				return
			}
			dataSaldo.Data[i].TotalSaldo -= t.Nominal
			dataSaldo.Data[i].Out += t.Nominal
		}
		if saldo.Username == t.To {
			dataSaldo.Data[i].TotalSaldo += t.Nominal
			dataSaldo.Data[i].In += t.Nominal
		}
	}

	newSaldoData, err := json.Marshal(dataSaldo)
	if err != nil {
		log.Println("Failed to marshal updated saldo data: " + err.Error())
		return
	}

	if err := ioutil.WriteFile(filename, newSaldoData, 0644); err != nil {
		log.Println("Failed to write updated saldo to file " + filename)
		return
	}

	var ActivityTransfer services.ActivityTransfer
	ActivityTransfer.From = t.Sender
	ActivityTransfer.To = t.To
	ActivityTransfer.Nominal = string(t.Nominal)
	ActivityTransfer.Timestamp = time.Now()

	services.CreateHistory("/storage/history.json", "Transfer activity recorded", &ActivityTransfer)
}

func LoginHandler() {
}

func LogoutHandler() {
}
