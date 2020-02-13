package QiniuClient

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestToken(t *testing.T) {
	accessKey := "WdlLw-oFN1mbj-0vngrbQ8Si39u0dgackq0v9L4T"
	secretKey := "HoKZLcwE1BtYWe9ze__jmwWsfqPeAcbHURN48t9A"
	bucket := "note-gin"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	log.Println(upToken)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	var localFile = "D:\\Go_Code\\note-gin\\utils\\QiniuClient\\1.jpg"
	bucket = "if-pbl"
	key := "1.png"
	data := []byte{}
	file, _ := os.Open(localFile)
	data, _ = ioutil.ReadAll(file)
	dataLen := int64(len(data))
	err := resumeUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}

}
