package main

import (
	"fmt"
	"giftCode/handle"
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


func Test_HandleClient(t *testing.T){
	retM := handle.HandleClient("r1czr5u2","nna")
	if retM["GiftCode"] == "r1czr5u2" && retM["GiftContent"] == "Insufficient quantity"{
		fmt.Println("Test_HandleClient pass")
	}else {
		fmt.Println("Test_HandleClient error")
	}




}