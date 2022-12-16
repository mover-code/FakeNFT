/***************************
@File        : data.go
@Time        : 2022/03/03 13:53:27
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : nft 数据
****************************/

package types

import (
	"github.com/jinzhu/gorm"
)

type (
	// Arrtibute 属性
	Arrtibute struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	}

	// Nft 图片
	NftJson struct {
		Name        string       `json:"name"`        // 名称
		Description string       `json:"description"` // 描述
		Image       string       `json:"image"`       // 图片地址
		Dna         string       `json:"dna"`         // DNA
		Date        int64        `json:"date"`        // 时间
		Rarity      int64        `json:"rarity"`      // 稀有度
		Level       string       `json:"level"`       // 级别
		Arrtibutes  []*Arrtibute `json:"arrtibutes"`  // 部位
	}

	// Config 配置
	Config struct {
		OutSum      int    // 输出图片数量
		Layer       string // 素材路径
		ImageWight  int    // 图片宽度
		ImageHight  int    // 图片高度
		IPFS        bool   // 是否上传至IPFS
		OutPut      string // 输出路径
		CollectName string // 系列名称
		ErrSum      int    // 错误数量
	}

	// Layer元素
	Layer struct {
		Id       int    // 标识
		Weight   int    // weight
		FileName string // 文件名
		Level    string // 级别
		Limit    int    // 限制
	}

	// allLayer
	AllLayer struct {
		Name   string
		Layers []*Layer
	}

	// 数据库

	NFTConf struct {
		gorm.DB
		Layer []NFTLayer `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"` // nft 部位id
		Name  string     // nft 集合名称
		Desc  string     // 描述
	}

	NFTImage struct {
		gorm.DB
		Src string // 图片地址
	}

	NFTLayer struct {
		gorm.DB
		Name string  // 部位名称
		Sort int64   // 标识和生成图片时的排序
		Info []int64 // 部位内容,图片与权重
	}

	NFTLayerInfo struct {
		gorm.DB
		Pic    int64 // 图片
		Weight int64 // 权重来获取稀有度
	}

	NFTFakeData struct {
		gorm.DB
		Nft         int64   // NFTConf
		Name        string  `json:"name"`        // 名称
		Description string  `json:"description"` // 描述
		Pic         int64   `json:"image"`       // 最终合成的图片
		Dna         string  `json:"dna"`         // DNA
		Date        int64   `json:"date"`        // 生成时间
		Rarity      int64   `json:"rarity"`      // 稀有度
		Arrtibutes  []int64 `json:"arrtibutes"`  // 部位 NFTLayerInfo
	}
)

// func init() {
//     db, err := gorm.Open("sqlite3", "gorm.db")
//     defer db.Close()
//     if err != nil {
//         panic("fiald to connect database")
//     }
//     db.AutoMigrate(&NFTConf{}, &NFTLayer{}, &NFTLayerInfo{}, &NFTImage{}, &NFTFakeData{})
// }
