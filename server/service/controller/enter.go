package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/stock"
	"runtime"
	"strings"
	"time"
)

type ControllerGroup struct {
	StockController
}

var stockBuyServer stock.StockBuyServer

func getMethodName() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fullName := fmt.Sprintf("%s", fn.Name())
	lastDotIndex := strings.LastIndex(fullName, ".")
	return fullName[lastDotIndex+1:]
}

func redisGet(name string) (str string, err error) {
	client := global.GVA_REDIS
	str, err = client.Get(context.Background(), name).Result()
	return str, err
}

func redisGetList(name string, list interface{}) (err error) {
	client := global.GVA_REDIS
	kk, err := client.Get(context.Background(), name).Result()
	if err == nil {
		json.Unmarshal([]byte(kk), &list)
		return nil
	}
	return err
}

func redisSetList(name string, list interface{}) (err error) {
	client := global.GVA_REDIS
	str, err := json.Marshal(list)
	client.Set(context.Background(), name, str, time.Minute*120).Err()
	return nil
}

func redisSetTotal(name string, total int) (err error) {
	client := global.GVA_REDIS
	client.Set(context.Background(), name, total, time.Minute*120).Err()
	return nil
}

func redisSetStr(name string, str string) (err error) {
	client := global.GVA_REDIS
	client.Set(context.Background(), name, str, time.Minute*120).Err()
	return nil
}
