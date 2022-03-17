package utils

import "fmt"

func ManageMsg(b []byte) []byte {
	msg := string(b)
	fmt.Printf("[INFO]%+v\n", msg)
	ToJson(b)
	return b //getMsg(b)
}
