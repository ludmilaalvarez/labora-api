package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

// Estructura BankAccount que guarda el saldo de la cuenta
type BankAccount struct {
	Balance int
	mutex sync.Mutex
}

// Método que incrementa el saldo de la cuenta con la cantidad dada
func (account *BankAccount) Deposit(amount int) {
	account.Balance += amount
}

// Método que decrementa el saldo de la cuenta con la cantidad dada
func (account *BankAccount) Withdraw(amount int) {
	account.Balance -= amount
}

// Función que realiza una transacción (depósito o retiro) en la cuenta
func processTransaction(account *BankAccount, transactionType string, amount int) {
	(*account).mutex.Lock()
	switch transactionType {
	case "deposit":
		account.Deposit(amount)
		fmt.Printf("Depósito: %d. Balance actual: %d\n", amount, account.Balance) // imprime la cantidad depositada y el saldo actual
	case "withdraw":
		account.Withdraw(amount)
		fmt.Printf("Retiro: %d. Balance actual: %d\n", amount, account.Balance) // imprime la cantidad retirada y el saldo actual
	}
	(*account).mutex.Unlock()
}

func main() {
	//var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano()) // semilla para generar números aleatorios
	account := &BankAccount{Balance: 1000} // inicializa una cuenta con un saldo de 1000
	fmt.Printf("Balance inicial: %d\n", account.Balance) // imprime el saldo inicial

	for i := 0; i < 5; i++ { // loop para generar 5 transacciones
		transactionType := "" // inicializa el tipo de transacción
		if i%2 == 0 { // si el índice es par, la transacción es un depósito
			transactionType = "deposit"
		} else { // si el índice es impar, la transacción es un retiro
			transactionType = "withdraw"
		}
		amount := rand.Intn(500) + 1 // genera una cantidad aleatoria entre 1 y 500 para la transacción
		go processTransaction(account, transactionType, amount) // inicia una goroutine para procesar la transacción
	}

	time.Sleep(3 * time.Second) // espera a que todas las goroutines finalicen
	fmt.Printf("Balance final: %d\n", account.Balance) // imprime el saldo final
}