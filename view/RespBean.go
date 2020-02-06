package view

type RespBean struct {
	Code int `form:"code" json:"code"`
	Msg  string `form:"msg" json:"msg"`
	Data interface{} `form:"data" json:"data"`
}
func NewRespBean() RespBean{
	return RespBean{}
}

func OkWithMsg(msg string) RespBean  {
	return RespBean{
		Code: 200,
		Msg:  msg,
		Data: nil,
	}
}

func OkWithData(msg string,data interface{}) RespBean{
	return RespBean{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}



func ErrorWithMsg(msg string) RespBean  {
	return RespBean{
		Code: 500,
		Msg:  msg,
		Data: nil,
	}
}

func ErrorWithData(msg string,data interface{}) RespBean{
	return RespBean{
		Code: 500,
		Msg:  msg,
		Data: data,
	}
}


