package main

import (
	"fmt"
	a "accounts"
	c "customers"
)


func payment(whichAccount verifyAccount, paymentValue float64) {
	whichAccount.cashout(paymentValue)
}

type verifyAccount interface {
	cashout(value float64) string
}

func main(){
	account1 := a.Account{
		Person: c.PersonalInfo{
			Name: "Lucas", 
			Document: 00011122233, 
			Profession: "Developer",
			}, 
		AgencyNumber: 0001, 
		AccountNumber: 12345, 
		TotalValue: 1100.00,
	}

	account2Info := a.Account{
		Person: c.PersonalInfo{
			Name: "Maria", 
			Document: 33344455566, 
			Profession: "Developer",
			}, 
	}
	account2 := a.Account{
		account2Info, 
		AgencyNumber: 0001, 
		AccountNumber: 54321, 
		TotalValue: 1000.00,
	}
	
	fmt.Println(account1)
	fmt.Println(account1.cashout(100.0))
	fmt.Println(account1.deposit(100.0))


	fmt.Println(account2.tranfer(200.0, &account1))
	fmt.Println(account1.getTotalValue())
	fmt.Println(account2.getTotalValue())

	payment(&account1, 60)

}