package bingx

import (
	"testing"
)

func TestXxx(t *testing.T) {

	Cntrcts()
	XorderSts(1735916532754554880)

	XcloseAllPositions()

	XallOrder("1000BONK-USDT", 50)

	TakeMrktOrder("1000BONK-USDT", 800)

	SetLeverage("BTC-USDT", "LONG", 61)
	SetLeverage("BTC-USDT", "SHORT", 61)

	GtBalance()

}

func TestKi(t *testing.T) {

	Xmain()
}

func TestSetLeverage(t *testing.T) {

	cd, err := Cntrcts()

	if err == nil {

		for _, cntc := range (*cd) {

			SetLeverage(cntc.Symbol,"LONG",50)
			SetLeverage(cntc.Symbol,"SHORT",50)
			
		}
		
		
	}

	

}
