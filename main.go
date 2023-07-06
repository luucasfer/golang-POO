package main

import (
	"fmt"
)


type Account struct {
	person 			string 		
	agencyNumber 	int 	
	accountNumber 	int 	
	totalValue 		float64 		
}

func (c *Account) cashout(subtractValue float64) string {  //quem chamar essa função eu vou apontar para a Account de quem estiver chamando
	canTakeCash := subtractValue > 0 && subtractValue <= c.totalValue // valor do saque é > 0 e <= ao valor total?
	if canTakeCash {
		c.totalValue -= subtractValue 
		return "Cashout realized with success"
	} else {
		return "Can not take this amount of money"
	}
}


func main(){
	newAccount := Account{}
	newAccount.person = "Lucas"
	newAccount.agencyNumber = 0001
	newAccount.accountNumber = 12345
	newAccount.totalValue = 1000.00
	fmt.Println(newAccount)

	newAccount.cashout(10000.0)
	fmt.Println(newAccount.totalValue)
}