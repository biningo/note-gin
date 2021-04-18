package view

/**
*@Author lyer
*@Date 2/22/21 17:01
*@Describe
**/

var HttpMsg = map[int]string{

	SUCCESS:                  "successful",
	ERROR:                    "ERROR",
	ERROR_FILE_TYPE:          "文件类型不允许",
	ERROR_FILE_NOT_EXIST:     "文件不存在",
	ERROR_TEMP_SAVE:          "文章保存失败",
	ERROR_RECOVER:            "文章恢复失败",
	ERROR_FILE_IS_EXIST:      "存在同名文件",
	FILE_IS_EXIST_AND_UPDATE: "存在同名文件，文件已经更新",
}

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_FILE_TYPE          = 10000
	ERROR_FILE_NOT_EXIST     = 10001
	ERROR_TEMP_SAVE          = 10002
	ERROR_RECOVER            = 10003
	ERROR_FILE_IS_EXIST      = 10004
	FILE_IS_EXIST_AND_UPDATE = 10005
)
