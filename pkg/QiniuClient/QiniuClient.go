package QiniuClient

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func GetToken() string {
	accessKey := "WdlLw-oFN1mbj-0vngrbQ8Si39u0dgackq0v9L4T"
	secretKey := "HoKZLcwE1BtYWe9ze__jmwWsfqPeAcbHURN48t9A"
	bucket := "note-gin"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
