package ctrl

import (
	"giftCode/handle"
	"github.com/gin-gonic/gin"
)

func AdminCreatGiftcode(c *gin.Context){
	des := c.Query("des")
	GiftNum := c.Query("GN")
	ValidPeriod :=c.Query("VP")
	GiftContent :=c.Query("GC")
	CreatePer := c.Query("CP")
	if len(des)== 0 && len(GiftNum)== 0 && len(ValidPeriod)== 0 && len(GiftContent)== 0 && len(CreatePer)== 0 {
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}
	retMap := handle.HandleAdminCreatGiftcode(des,GiftNum,ValidPeriod,GiftContent,CreatePer)
	c.JSON(200,gin.H{
		"condition":retMap["condition"],
		"GiftCode" : retMap["GiftCode"],
	})
}
func AdminInquireGiftCode(c *gin.Context){
	GiftCode := c.Query("giftCode")
	if len(GiftCode) == 0  {
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}
	retMap,infoMap := handle.HadnleAdminInquireGiftCode(GiftCode)
	c.JSON(200, gin.H{
		"condition": retMap["condition"],
		"GiftCode":  GiftCode,
		"data" : infoMap,
	})

}
func Client(c *gin.Context)  {
	GiftCode := c.Query("giftCode")
	userName := c.Query("usr")

	if len(GiftCode) == 0 && len(userName) == 0{
		c.JSON(200,gin.H{
			"condition":"error",
		})
		return
	}
	ret := handle.HandleClient(GiftCode,userName)
	c.JSON(200,gin.H{
		"condition" :ret["condition"] ,
		"GiftContent" : ret["GiftContent"],
		"GiftCode" : ret["GiftCode"],
	})
}