package RedisClient

var client = RedisInit()

func AccessRecord() {
	client.Incr("AccessRecord")
}

func GetAccessRecord() (result string) {
	result = client.Get("AccessRecord").Val()
	return
}
func DelAccessRecord() {
	client.Del("AccessRecord") //清除redis记录
}
