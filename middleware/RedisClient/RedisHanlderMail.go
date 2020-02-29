package RedisClient

var client = RedisInit()

func AccessRecord(body string) {

	client.LPush("AccessRecord", body)
}

func GetAccessRecord() (result []string) {
	result = client.LRange("AccessRecord", 0, client.LLen("AccessRecord").Val()).Val()
	return
}
func DelAccessRecord() {
	client.Del("AccessRecord") //清除redis记录
}
