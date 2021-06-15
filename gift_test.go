package main

import (
	"fmt"
	"giftCode/util"
	"testing"
)

func Test_GetRandCode(t *testing.T){

	retString := util.GetRandCode(8)
	if len(retString) != 8{
		fmt.Println("Test_GetRandCode error")
	}else {
		fmt.Println("Test_GetRandCode pass")
	}
}