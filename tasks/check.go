package tasks

import (
	"strings"
	"sync"
)

var codec *Codec
var taskCache = struct {
	sync.RWMutex
	data map[string]string
}{
	data: make(map[string]string),
}

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
	taskCache.RLock()
	defer taskCache.RUnlock()
	return taskCache.data[str]
}

// setTask 设置任务状态到本地缓存
// str 格式可以是 "taskId=status" 或单独的任务ID
func setTask(str string) error {
	if str == "" {
		return nil
	}

	taskCache.Lock()
	defer taskCache.Unlock()

	// 检查是否包含等号，格式为 "key=value"
	if idx := strings.Index(str, "="); idx > 0 && idx < len(str)-1 {
		key := str[:idx]
		value := str[idx+1:]
		taskCache.data[key] = value
	} else {
		// 如果没有等号，将整个字符串作为key，设置一个默认状态
		taskCache.data[str] = "" // 默认状态为空
	}
	return nil
}

// removeTask 从缓存中移除任务
func removeTask(str string) error {
	if str == "" {
		return nil
	}

	taskCache.Lock()
	defer taskCache.Unlock()
	delete(taskCache.data, str)
	return nil
}
