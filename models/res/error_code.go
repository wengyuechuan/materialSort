package res

const (
	SuccessCode   int = 200 // 成功
	ArgumentError int = 400 // 参数错误
	SystemError   int = 500 // 系统错误
)

var (
	ErrorMap = map[int]string{
		ArgumentError: "参数错误",
		SystemError:   "系统错误",
		SuccessCode:   "成功",
	}
)
