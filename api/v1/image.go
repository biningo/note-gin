package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io/ioutil"
	"log"
	"note-gin/pkg/QiniuClient"
	"note-gin/view"
)

/**
*@Author lyer
*@Date 2/22/21 15:41
*@Describe
**/

var accessKey = QiniuClient.QiniuAccessKey
var secretKey = QiniuClient.QiniuSecretKey
var bucket = "note-gin"

func ImgUpload(c *gin.Context) {

	fileUp, _ := c.FormFile("img")
	file, _ := fileUp.Open()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

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
	bucket = "note-gin"
	key := fileUp.Filename
	data := []byte{}

	manager := storage.NewBucketManager(mac, &cfg)
	FileInfo, err := manager.Stat(bucket, key)
	if FileInfo.Fsize != 0 { //图片存在
		url := "http://gin-note.binnb.top/" + key
		c.JSON(200, view.OkWithData("图片已经存在!", url))
		return
	}

	data, _ = ioutil.ReadAll(file)
	dataLen := int64(len(data))
	err = resumeUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "http://gin-note.binnb.top/" + key
	log.Println("上传图片：", url)
	c.JSON(200, view.OkWithData("图片上传成功!", url))
}

func ImgDelete(c *gin.Context) {

	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := "note-gin"
	key := c.Query("img_name")

	err := bucketManager.Delete(bucket, key)
	if err != nil {
		c.JSON(200, view.ErrorWithMsg("云存储图片删除失败!"))
	} else {
		c.JSON(200, view.OkWithMsg("云存储图片删除成功!"))
	}
}
