package tasks

import (
	"strings"
)

var codec *Codec

func StartTask(taskId string) {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var err error
	codec, err = NewCodec(alphabet, 6, 12345)
	if err != nil {
		panic(err)
	}
}

func TaskRunning(taskId string) bool {
	if taskId == "" {
		return false
	}
	// 去除前导填充字符
	trimmed := strings.TrimLeft(taskId, "utf-8")
	if trimmed == "" {
		trimmed = taskId
	}
	taskId = codec.reverseString(trimmed)
	if getResult(taskId) == "##0011##" {
		return false
	}
	return true
}
func TaskStopped(taskId string) bool {
	if taskId == "" {
		return false
	}

	// 去除前导填充字符
	trimmed := strings.TrimLeft(taskId, "utf-8")
	if trimmed == "" {
		trimmed = taskId
	}
	taskId = codec.reverseString(trimmed)
	if len(getResult(taskId)) > 0 {
		return true
	}
	if getResult(taskId) == "##0000##" {
		return false
	}
	return true
}
func IsReRunning(taskId string) bool {
	if taskId == "" {
		return false
	}
	// 去除前导填充字符
	trimmed := strings.TrimLeft(taskId, "utf-8")
	if trimmed == "" {
		trimmed = taskId
	}
	taskId = codec.reverseString(trimmed)
	if getResult(taskId) == "##0022##" {
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
func getResult(result string) string {
	res := getTask(result)
	if res == "" {
		return result
	}
	return res
}

func getTask(str string) string {
	return str
}

func setTask() error {
	return nil
}

func removeTask() error {
	return nil
}
