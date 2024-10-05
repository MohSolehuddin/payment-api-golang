# Payment RESTful API - Golang

Welcome to the **Payment RESTful API**, a backend system developed in **Golang** that facilitates seamless payment functionality between merchants and banks. This project aims to provide a simplified yet robust API that supports login, payment, and logout processes for customers.

## ✨ Key Features

- **Login**  
  Customers can log into the system. If a customer does not exist, the login request is rejected, ensuring secure access.
- **Payment**  
  After logging in, customers can transfer funds without a limit on the amount. Payments can only be sent to registered customers, keeping transactions within trusted users.
- **Logout**  
  Once transactions are completed, customers can safely log out of the system.

## 📊 Transaction History and Audit

All customer activities, such as login, payment, and logout, are logged into a history file for future audit and tracking purposes. This ensures transparency and accountability for every transaction.

## 💾 Data Management

Customer, merchant, and transaction history data are simulated using **JSON files**, making the system lightweight and easy to set up without requiring a full database.

## 🚀 Getting Started

To get started with this project, clone the repository and follow the instructions below to run the API on your local machine.

### Prerequisites

- Go 1.16 or later installed
- Git installed

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/YourUsername/payment-restful-api-golang.git
   ```
2. Navigate to the project directory:

   ```bash
   cd payment-restful-api-golang
   ```

3. Run the API:
   ```bash
   go run main.go
   ```

## 🔐 Security and Expansion

If you're looking for a more complex implementation with advanced security features such as **JWT (JSON Web Token)** for authentication, check out a similar project developed in **Java** using **Spring Boot** and **Spring Security**. You can find the repository here:

[Loan Application API - Java, Spring Boot](https://github.com/MohSolehuddin/loan-app)

## 🛠️ Tech Stack

- **Golang**: Backend programming language for API development.
- **JSON**: For data simulation (customers, merchants, transaction history).

## 👤 Author

**[Moh Solehuddin]**
