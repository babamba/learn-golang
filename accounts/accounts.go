package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

//  메모리 자체를 리턴한다. create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// receiver 에서는 새로운 복사본이 생성된다. 포인터가 다르단 얘기
//func (a Account) Deposit(amount int) {

// pointer receiver : *를 통해 기존 receiver 처럼 복사하지 않고 메서드를 호출한 account를 사용하게 처리
func (a *Account) Deposit(amount int) {
	//fmt.Println("Gonna deposit : ", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount your account
func (a *Account) Withdraw(amount int) error {
	// Go lang 에는 try - catch 같은 게 없다.
	// 직접 error를 리턴해줘야하고 직접 체크해야한다.
	// a.balance -= amount
	if a.balance < amount {
		//return error.Error() 기본
		return errNoMoney
	}
	a.balance -= amount
	return nil // js의 null 같은 것
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// 메모리가아니라 getter setter 처럼 할수도있음.
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	//return "whatever you want"
	return fmt.Sprint(a.Owner(), "'s account.\nHas : ", a.Balance())
}
