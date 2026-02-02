package tasks

func StartTask(taskId string) {

}

func TaskRunning(taskId string) bool {
	if taskId == "" {
		return false
	}
	return true
}
