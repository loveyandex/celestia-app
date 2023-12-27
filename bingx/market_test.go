package bingx

import (
	"fmt"
	"strconv"
	"testing"
)

func TestXPricexx(t *testing.T) {

	Price, err2 := XPrice("1000BONK-USDT")

	if err2 != nil {

		fmt.Printf("err2: %v\n", err2)
	}

	fmt.Printf("f: %v\n", Price)

	b, err := GtBalance()

	if err != nil {

		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("b.AvailableMargin: %v\n", b.AvailableMargin)
	
	bavl, err2 := strconv.ParseFloat(b.AvailableMargin, 64)

	if err2 != nil {
	
	}

	how:=bavl/Price/100

	fmt.Printf("how: %v\n", how)


	TakeMrktOrder("1000BONK-USDT",int64(how))

}
