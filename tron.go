package api

import (
	"shop/config"
	"shop/tronApi"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Tronrddress(c *gin.Context) {
	logger := zap.NewExample()
	defer logger.Sync()
	tronapi := tronApi.NewTronApiEngine(config.Tron.Url, config.Tron.Key, logger)
	// blok, _ := tronapi.GetNowBlock()
	add, privateKey, _ := tronapi.Trconaddress()
	var datas config.Tronaddessconf
	datas.Add = add
	datas.Key = privateKey
	//
	var set string
	if config.Redisconf.Lasting == "1" {
		set, _ = config.Rdb.Set(config.Ctx, add, privateKey, 0).Result()
	} else {
		set, _ = config.Rdb.Set(config.Ctx, add, privateKey, 10*time.Hour).Result()
	}

	if set == "OK" {
		res := make(map[string]interface{})
		res["address"] = add
		config.Resdata.Date = res
		c.JSON(200, config.Resdata)
	} else {
		config.Resdata.Code = 0
		config.Resdata.Date = ""
		c.JSON(200, config.Resdata)
	}

}
