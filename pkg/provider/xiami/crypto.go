package xiami

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"
)

const (
	APPKey = "23649156"
)

var (
	reqHeader = map[string]interface{}{
		"appId":      200,
		"platformId": "h5",
	}
)

func signPayload(token string, model interface{}) map[string]string {
	payload := map[string]interface{}{
		"header": reqHeader,
		"model":  model,
	}
	requestBytes, _ := json.Marshal(payload)
	data := map[string]string{
		"requestStr": string(requestBytes),
	}
	dataBytes, _ := json.Marshal(data)
	dataStr := string(dataBytes)

	t := fmt.Sprintf("%d", time.Now().UnixNano()/(1e6))
	signStr := fmt.Sprintf("%s&%s&%s&%s", token, t, APPKey, dataStr)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))

	return map[string]string{
		"t":    t,
		"sign": sign,
		"data": dataStr,
	}
}
