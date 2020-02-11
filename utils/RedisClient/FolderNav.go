package RedisClient

import (
	"note-gin/model"
)

func ChangeFolderNav(folder model.Folder) (nav []string) {
	RedisInit()
	client := RedisClient
	if folder.Title == "Home" {
		client.Del("folder_nav")
		return
	}

	length := client.LLen("folder_nav").Val()

	nav = client.LRange("folder_nav", 0, length-1).Val()
	result := []string{}

	for i, v := range nav {
		if folder.Title == v {

			client.LTrim("folder_nav", int64(i), length-1)
			result = client.LRange("folder_nav", 0, client.LLen("folder_nav").Val()-1).Val()
			nav = result

		}
	}

	//如果是新目录
	if len(result) == 0 {
		arr := []string{folder.Title}
		nav = append(arr, nav...)
		client.LPush("folder_nav", folder.Title)

	}

	//log.Println(nav)
	return
}
