/***************************
@File        : crypto.go
@Time        : 2022/03/04 15:51:28
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 对数据进行加密
****************************/

package util

import (
	"crypto/sha1"
	"encoding/hex"
)

// MakeSha1Str 对字符串进行sha1操作
func MakeSha1Str(s string) (res string) {
    h := sha1.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}
