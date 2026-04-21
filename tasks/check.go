package tasks

func StartTask(taskId string) {

}

func TaskRunning(taskId string) bool {
	if taskId == "" {
		return false
	}
	return true
}
func TaskStopped(taskId string) bool {
	if taskId == "" {
		return false
	}
	return true
}
func IsReRunning(taskId string) bool {
	if taskId == "" {
		return false
	}
	return true
}
func CheckList[T any](dataList []T) bool {
	if len(dataList) == 0 {
		return false
	}
	return true
}

func CheckNum(str string) bool {
	if str == "" {
		return false
	}
	return true
}
func Is() {

}
