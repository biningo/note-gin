package CacheCount

import (
	"github.com/go-redis/redis"
	"strconv"
)

type CacheCountImpl struct {
}

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "118.178.180.115:6379",
		DB:       1,
		Password: "biningo",
	})

	return client
}

func (CacheCountImpl) SetSubFileTotal(FolderID int64, val int) {

	client := RedisInit()
	client.Set("SubFileTotal:"+strconv.Itoa(int(FolderID)), val, -1)
}

func (CacheCountImpl) GetSubFileTotal(FolderID int64) (int, bool) {
	client := RedisInit()

	flag := client.Exists("SubFileTotal:" + strconv.Itoa(int(FolderID))).Val()

	if flag != 1 {
		return 0, false
	} else {
		valStr := client.Get("SubFileTotal:" + strconv.Itoa(int(FolderID))).Val()
		value, _ := strconv.Atoi(valStr)
		return value, true
	}

}

func (CacheCountImpl) AddSubFileTotal(FolderID int64, val int) {
	client := RedisInit()
	client.IncrBy("SubFileTotal:"+strconv.Itoa(int(FolderID)), int64(val))
}

func (CacheCountImpl) SubtractSubFileTotal(FolderID int64, val int) {
	client := RedisInit()
	client.DecrBy("SubFileTotal:"+strconv.Itoa(int(FolderID)), int64(val))
}

func (CacheCountImpl) DelSubFileTotal(FolderID int64) {
	client := RedisInit()
	client.Del("SubFileTotal:" + strconv.Itoa(int(FolderID)))
}

//FolderCount

func (CacheCountImpl) SetSubFolderTotal(FolderID int64, val int) {

	client := RedisInit()
	client.Set("SubFolderTotal:"+strconv.Itoa(int(FolderID)), val, -1)
}

func (CacheCountImpl) GetSubFolderTotal(FolderID int64) (int, bool) {
	client := RedisInit()

	flag := client.Exists("SubFolderTotal:" + strconv.Itoa(int(FolderID))).Val()

	if flag != 1 {
		return 0, false
	} else {
		valStr := client.Get("SubFolderTotal:" + strconv.Itoa(int(FolderID))).Val()
		value, _ := strconv.Atoi(valStr)
		return value, true
	}

}

func (CacheCountImpl) AddSubFolderTotal(FolderID int64, val int) {
	client := RedisInit()
	client.IncrBy("SubFolderTotal:"+strconv.Itoa(int(FolderID)), int64(val))
}

func (CacheCountImpl) SubtractSubFolderTotal(FolderID int64, val int) {
	client := RedisInit()
	client.DecrBy("SubFolderTotal:"+strconv.Itoa(int(FolderID)), int64(val))
}

func (CacheCountImpl) DelSubFolderTotal(FolderID int64) {
	client := RedisInit()
	client.Del("SubFolderTotal:" + strconv.Itoa(int(FolderID)))
}
