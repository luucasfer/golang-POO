package accounts

import "customers"

type Account struct {
	Person 			customers.PersonalInfo 	
	AgencyNumber 	int 	
	AccountNumber 	int 	
	totalValue 		float64 		
}

func (c *Account) cashout(subtractValue float64) (string, float64) {  //quem chamar essa função eu vou apontar para a Account de quem estiver chamando
	canTakeCash := subtractValue > 0 && subtractValue <= c.totalValue // valor do saque é > 0 e <= ao valor total?
	if canTakeCash {
		c.totalValue -= subtractValue 
		return "Cashout realized with success - Available value: ", c.totalValue
	} else {
		return "Can not take this amount of money - Available value:", c.totalValue
	}
}

func (c *Account) deposit(depositValue float64) (string, float64) {
	if depositValue > 0{
		c.totalValue += depositValue
		return "Value deposited with success - Available value: ", c.totalValue
	} else {
		return "Can not deposit negative values - Available value: ", c.totalValue
	}
}

func (c *Account) tranfer(transferValue float64, destiny *Account) bool {
	if transferValue < c.totalValue && transferValue > 0{
		c.totalValue -= transferValue
		destiny.deposit(transferValue)
		return true
	} else {
		return false
	}
}


func (c *Account) getTotalValue() float64 {
	return c.totalValue
}