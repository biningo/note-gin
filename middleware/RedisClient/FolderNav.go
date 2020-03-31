package RedisClient

import (
	"note-gin/model"
)

func ChangeFolderNav(folder model.Folder) (nav []string) {
	client := RedisInit()
	defer client.Close()

	if folder.Title == "Home" {
		client.Del("folder_nav")
		return
	}

	length := client.LLen("folder_nav").Val()

	nav = client.LRange("folder_nav", 0, length-1).Val()
	//注意 这里的nav顺序是反的  0是最后一个目录
	if len(nav) > 0 && nav[0] == folder.Title { //如果page=1 同时还是本目录 则不执行下面的操作
		return
	}

	result := []string{}

	for i, v := range nav {
		if folder.Title == v {
			client.LTrim("folder_nav", int64(i), length-1) //【】前后都包括 所以这里-1 当然超过了也没事 不会报错
			result = client.LRange("folder_nav", 0, client.LLen("folder_nav").Val()-1).Val()
			nav = result
		}
	}

	//如果是新项
	if len(result) == 0 {
		arr := []string{folder.Title}
		nav = append(arr, nav...)
		client.LPush("folder_nav", folder.Title)
	}
	return
}

func GetCurrentNav() (nav []string) {
	client := RedisInit()
	defer client.Close()

	length := client.LLen("folder_nav").Val()
	if length > 0 {
		nav = client.LRange("folder_nav", 0, length-1).Val()
	}
	return
}
