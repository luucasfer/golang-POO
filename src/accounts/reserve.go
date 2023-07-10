package accounts

import "customers"

type ReserveAccount struct {
	Name									customers.PersonalInfo
	AgencyNumber, AccountNumber, Operation	int
	totalValue								float64
}


func (c *ReserveAccount) cashout(subtractValue float64) (string, float64) {  //quem chamar essa função eu vou apontar para a Account de quem estiver chamando
	canTakeCash := subtractValue > 0 && subtractValue <= c.totalValue // valor do saque é > 0 e <= ao valor total?
	if canTakeCash {
		c.totalValue -= subtractValue 
		return "Cashout realized with success - Available value: ", c.totalValue
	} else {
		return "Can not take this amount of money - Available value:", c.totalValue
	}
}

func (c *ReserveAccount) deposit(depositValue float64) (string, float64) {
	if depositValue > 0{
		c.totalValue += depositValue
		return "Value deposited with success - Available value: ", c.totalValue
	} else {
		return "Can not deposit negative values - Available value: ", c.totalValue
	}
}

func (c *ReserveAccount) getTotalValue() float64 {
	return c.totalValue
}