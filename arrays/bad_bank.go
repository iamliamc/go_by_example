package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func applyAchar(accumulator string, current string) string {
	// fmt.Println(accumulator, current)
	newString := accumulator + "s" + current
	return newString
}

func ApplyStringToCollection(initialValue string, collection []string, stringToAdd string) string {
	return Reduce(
		collection,
		applyAchar,
		initialValue,
	)
}
