package main

import (
	"giftCode/config"
	"giftCode/router"
)

func main()  {


	config.InitClient()
	r:=router.SetUpRount()
	r.Run(":8080")
}
