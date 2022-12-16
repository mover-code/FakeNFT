/***************************
@File        : pil_test.go
@Time        : 2022/03/04 16:49:14
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : pil_test
****************************/

package util

import (
    "fmt"
    "testing"
)

func TestDraw(t *testing.T) {
    a := "../layers/gem_1/1#bg/bg#4.jpg"
    b := "../layers/gem_1/2#gem/gem#80.png"
    fmt.Println(Draw(a, b, "./new.jpg"))
}
