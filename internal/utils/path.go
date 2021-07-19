package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("GetCurrentPath error:%+v\n", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
