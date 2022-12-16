/***************************
@File        : main.go
@Time        : 2022/03/03 13:47:19
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : golang 根据素材生成nft图片
****************************/

package main

import (
	"fake-nft/cmd"
	"flag"
)

var (
	count         = flag.Int("amount", 10, "Input your NFT Amount.")
	layer         = flag.String("layer", "demo", "Input your layer path.")
	out           = flag.String("out", "demo", "Input your out path.")
	collection    = flag.String("collection", "Demo-Collection", "Input your collection name.")
	Input_flagvar int
)

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
}

func main() {
	Init()
	flag.Parse()
	if *count <= 0 {
		panic("argument `amount` value not supported !")
	}
	cmd.StartCreating(cmd.NewConf(*count, *layer, *out, *collection, 2000, 2000, false))
}
