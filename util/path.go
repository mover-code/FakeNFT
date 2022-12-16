/***************************
@File        : path.go
@Time        : 2022/02/18 17:35:43
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 工作目录操作
****************************/

package util

import (
    "os"
    "path/filepath"
)

func GetPath() string {
    path, _ := os.Getwd()
    return path
}

func JoinPath(base, path string) string {
    return filepath.Join(base, path)
}

// Exists 判断文件或目录是否存在
func Exists(path string) bool {
    if _, err := os.Stat(path); err == nil {
        if os.IsExist(err) {
            return true
        }
    } else {
        return false
    }
    return true
}
