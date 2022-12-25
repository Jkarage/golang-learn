package main

const bankBallance = 4987

var loanEligible bool

func main() {
	//If statements example
	if hasCredit() {
		loanEligible = true
	} else {
		loanEligible = false
	}

	// Switch case statements
	switch hasCredit() {
	case true:
		loanEligible = true
	case false:
		loanEligible = false
	}

}

var hasCredit = func() bool {
	return bankBallance > 5000
}
