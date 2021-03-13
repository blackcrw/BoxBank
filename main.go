package main

import (
	"fmt"

	bank "ByteBank/system"
)

func main() {
	maria := bank.Bank{
		Account: bank.Person{"Maria", "726.018.050-13"},
		Balance: 1000.12,
	}

	joao := bank.Bank{
		Account: bank.Person{"João", "869.674.550-79"},
		Balance: 12212.12,
	}

	maria.Transfer(joao, 320.53)
	// joao.Transfer(maria, 1000.00)

	// maria.Draw(100.1)
	// joao.Draw(213.14)

	// maria.Deposit(12.00)
	// joao.Deposit(21.00)

	fmt.Print("\n")
	maria.Extract()
	joao.Extract()

}
