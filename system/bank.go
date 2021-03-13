package bank

import "fmt"

// Bank :: Struct
type Bank struct {
	Account Person
	Balance float32
}

// Person :: Struct
type Person struct {
	Name string
	Cpf  string
}

// Transfer :: Function
func (bank *Bank) Transfer(account Bank, money float32) {
	if bank.Balance >= money {
		account.Balance += money
		bank.Balance -= money

		fmt.Printf("Transferência no valor de R$%.2f, feita com sucesso para %s\n", money, account.Account.Name)
	} else {
		fmt.Println("Abistinência do valor requisitado.")
	}
}

// Draw  :: Function
func (bank *Bank) Draw(money float32) {
	if bank.Balance >= money {
		bank.Balance -= money

		fmt.Printf("O valor R$%.2f foi sacado com sucesso!\n", money)
	} else {
		fmt.Println("Não foi possível sacar o valor requisitado.")
	}

}

// Deposit :: Function
func (bank *Bank) Deposit(money float32) {
	bank.Balance += money

	fmt.Printf("O valor R$%.2f foi depositado com sucesso!\n", money)
}

func (bank Bank) Extract() {
	fmt.Printf("O valor que você %s tem guardado na sua conta é %.2f\n", bank.Account.Name, bank.Balance)
}
