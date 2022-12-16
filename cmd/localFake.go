/***************************
@File        : fake.go
@Time        : 2022/03/03 14:17:20
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 生成nft图片入口
****************************/

package cmd

import (
	"fake-nft/types"
	"fake-nft/util"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var DNA = map[string]int{}
var Last = time.Now()
var START = time.Now()
var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

// NewConf 配置
func NewConf(outSum int, layer, output, collectionName string, weight, height int, upload bool) *types.Config {
	// 默认demo数据
	baseLayer := "layers"
	baseOut := "out"
	if outSum <= 0 {
		outSum = 10
	}
	if layer == "" {
		layer = util.JoinPath(baseLayer, "layer")
	} else {
		layer = util.JoinPath(baseLayer, layer)
	}
	util.Mkdir(layer)
	if output == "" {
		output = util.JoinPath(baseOut, "demo")
	} else {
		output = util.JoinPath(baseOut, output)
	}
	util.Mkdir(output)
	if collectionName == "" {
		collectionName = "demo"
	}
	if weight <= 0 || height <= 0 {
		weight, height = 512, 512
	}
	// 初始化输出
	util.Mkdir(util.JoinPath(output, "json"))
	util.Mkdir(util.JoinPath(output, "image"))

	return &types.Config{
		OutSum:      outSum,
		Layer:       layer,
		OutPut:      output,
		ImageWight:  weight,
		ImageHight:  height,
		CollectName: collectionName,
		IPFS:        upload,
		ErrSum:      1,
	}
}

// StartCreating 开始创建nft
func StartCreating(c *types.Config) {
	// panic(c.CollectName)
	layers, err := loadLayers(c)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(c.ErrSum)

	// allInfo := []*types.NftJson{}
	errCount := 0
	success := 0
	i := 0
MAKENFT:
	for {
		dna, level, pics, arrtibute := createDNA(layers)
		// fmt.Println(len(pics))
		// panic("hhh")
		sha1Dna := util.MakeSha1Str(dna)
		if _, ok := DNA[sha1Dna]; !ok {
			i += 1
			DNA[sha1Dna] = i

			for j := 0; j < len(pics); j++ {
				if j == 0 {
					util.Draw(util.JoinPath(c.Layer, pics[0]), util.JoinPath(c.Layer, pics[1]), util.JoinPath(util.JoinPath(c.OutPut, "image"), fmt.Sprintf("%v.jpg", success+1)))
				}
				if j > 1 {
					util.Draw(util.JoinPath(util.JoinPath(c.OutPut, "image"), fmt.Sprintf("%v.jpg", success+1)), util.JoinPath(c.Layer, pics[j]), util.JoinPath(util.JoinPath(c.OutPut, "image"), fmt.Sprintf("%v.jpg", success+1)))
				}
			}
			success += 1
			fmt.Println("process--", success, c.CollectName, level, time.Since(Last))
			Last = time.Now()
			fmt.Println("----", len(DNA))

			info := &types.NftJson{
				Name:        fmt.Sprintf("%v-%v", c.CollectName, success),
				Description: "Remember to replace this description",
				Image:       util.JoinPath(util.JoinPath(c.OutPut, "image"), fmt.Sprintf("%v.jpg", success)),
				Dna:         sha1Dna,
				Date:        time.Now().UnixMicro(),
				Rarity:      TotalWeight(arrtibute),
				Level:       level,
				Arrtibutes:  arrtibute,
			}
			util.Write(&info, util.JoinPath(util.JoinPath(c.OutPut, "json"), fmt.Sprintf("%v-%v.json", success, level)))
			// allInfo = append(allInfo, info)
		} else {
			errCount += 1
		}
		if errCount > c.ErrSum || success >= c.OutSum {
			break MAKENFT
		}
	}
	fmt.Println("stop...", "error:", errCount, "success:", success, "times:", time.Since(START))

	// util.Write(&allInfo, util.JoinPath(util.JoinPath(c.OutPut, "json"), fmt.Sprintf("all.json")))

}

func loadLayers(c *types.Config) ([]*types.AllLayer, error) {
	childrens, err := util.GetChildFile(c.Layer)
	if err != nil {
		return nil, err
	}
	layers := []*types.AllLayer{}
	for _, child := range childrens {
		var elements []*types.Layer
		elems, _ := util.GetChildFile(util.JoinPath(c.Layer, child))
		for id, element := range elems {
			elements = append(elements, &types.Layer{
				Id:       id,
				FileName: element,
				Weight:   util.FindNum(element),
				Level:    util.FindLevel(element),
				Limit:    util.FindLimit(element),
			})

			// fmt.Println(child, element, "-", util.FindNum(element), util.FindLevel(element), util.FindLimit(element))
		}
		layers = append(layers, &types.AllLayer{
			Name:   child,
			Layers: elements,
		})
		c.ErrSum = c.ErrSum * len(elements)
	}
	c.ErrSum = c.ErrSum * 100000
	return layers, nil
}

func createDNA(layers []*types.AllLayer) (dna, level string, pics []string, arrtibutes []*types.Arrtibute) {
	level = "N"
	index := 0
	// fmt.Println("开始生成DNA", layers)
	d := -1
	b := -1

	for j, v := range layers {
		totalWeight := 0
		index += 1
		for _, element := range v.Layers {
			// fmt.Println(element.Limit)
			if element.Limit != 1 {
				totalWeight += element.Weight
			}
			// fmt.Println(element.FileName, element.Id, element.Weight)
		}
		// fmt.Println(totalWeight)
		// 随机方式
		// 	randNum := r1.Intn(totalWeight)

		// ELEMENT:
		// 	for i, element := range v.Layers {
		// 		randNum -= element.Weight
		// 		if randNum < 0 && element.Limit != 1 {
		// 			if index == len(layers) {
		// 				dna += fmt.Sprintf("%v:%v#%v", element.Id, v.Name, element.FileName)
		// 			} else {
		// 				dna += fmt.Sprintf("%v:%v#%v-", element.Id, v.Name, element.FileName)
		// 			}
		// 			pics = append(pics, util.JoinPath(v.Name, element.FileName))
		// 			arrtibutes = append(arrtibutes, &types.Arrtibute{
		// 				TraitType: v.Name[1:],
		// 				Value:     strings.Split(util.SplitFileName(element.FileName)[0], "&")[0],
		// 			})
		// 			if element.Limit > 1 {
		// 				d = i
		// 				b = j
		// 				// fmt.Println(el.FileName, "----")
		// 			}
		// 			break ELEMENT
		// 		}
		// 	}

		i := r1.Intn(len(v.Layers))
		// i := 0
		element := v.Layers[i]
		// fmt.Println("=")
		// fmt.Println(element.FileName, element.Level, element.Limit)
		if element.Limit != 1 {
			if index == len(layers) {
				dna += fmt.Sprintf("%v:%v#%v", element.Id, v.Name, element.FileName)
			} else {
				dna += fmt.Sprintf("%v:%v#%v-", element.Id, v.Name, element.FileName)
			}
			pics = append(pics, util.JoinPath(v.Name, element.FileName))
			arrtibutes = append(arrtibutes, &types.Arrtibute{
				TraitType: v.Name[1:],
				Value:     strings.Split(util.SplitFileName(element.FileName)[0], "&")[0],
			})
			if element.Limit > 1 {
				d = i
				b = j
			}
		}

	}
	if d != -1 && b != -1 {
		// fmt.Println(layers[b].Layers[d], d, b)
		sha1Dna := util.MakeSha1Str(dna)
		if _, ok := DNA[sha1Dna]; !ok {
			if layers[b].Layers[d].Limit > 1 {
				layers[b].Layers[d].Limit -= 1
				fmt.Println("limit layer", layers[b].Layers[d].FileName, layers[b].Layers[d].Limit, layers[b].Layers[d].Level)
				level = layers[b].Layers[d].Level
				if layers[b].Layers[d].Limit == 1 {
					layers[b].Layers = append(layers[b].Layers[:d], layers[b].Layers[d+1:]...)
				}
			}
		}
	}
	return
}

func TotalWeight(layers []*types.Arrtibute) (weight int64) {
	for _, l := range layers {
		weight += int64(util.FindNum(l.Value))
	}
	return
}
