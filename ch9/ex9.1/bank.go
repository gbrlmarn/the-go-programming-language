// Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program. The result should indicate whether the transaction succeded or failed due to insufficient funds. The message sent to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can sent the boolean result back to Withdraw.
package bank

type transaction struct {
	amount chan int
	ok     chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balances
var withdraw transaction      // send amount to withdraw

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraw.amount <- amount
	return <-withdraw.ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw.amount:
			if amount <= balance {
				balance -= amount
				withdraw.ok <- true
			} else {
				withdraw.ok <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
