package main

import (
	"fmt"

	"github.com/genesiskoo/test-go/banking"
)

func main() {

	account := banking.Account{Owner: "koo", Balance: 10000}
	fmt.Println(account)

}
