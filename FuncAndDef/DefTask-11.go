package main

import "fmt"

// создается структура Account
type Account struct {
	Balance int
}

func main() {
	initialBalance := 1000
	// тут мы зачисляем на баланс 1000
	account := &Account{Balance: initialBalance}

	defer printBalance("Изначальный баланс", account.Balance) // 1000
	defer printBalance("Текущий баланс", account.Balance)     // 1000
	defer printAccountBalance("Указатель на баланс", account) // 1300
	//

	account.Balance += 500           // 1500
	updateBalance(account, 200)      // 1300
	account = &Account{Balance: 300} // account указывает на новый обьект
}

// функция обновления баланса с указателем на стуктуру *Account
func updateBalance(acc *Account, amount int) {
	acc.Balance -= amount
}

// фунцкия вывода Баланса на консоль
func printBalance(label string, amount int) {
	fmt.Printf("%s: %d\n", label, amount)
}

// это указатель на стуктуру Account
func printAccountBalance(label string, acc *Account) {
	fmt.Printf("%s: %d\n", label, acc.Balance)
}
