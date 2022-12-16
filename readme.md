# 根据图层批量生成nft工具

## 输入

按照layers demo 中示例添加图层由内到外命名方式按照示例

```tree
.
├── 1Background
│   └── Black#1.png
├── 2Eye\ ball
│   ├── Red#50.png
│   └── White#50.png
├── 3Eye\ color
│   ├── Cyan#1.png
│   ├── Green#1.png
│   ├── Pink#1.png
│   ├── Purple#1.png
│   ├── Red#1.png
│   └── Yellow#10.png
├── 4Iris
│   ├── Large#20.png
│   ├── Medium#20.png
│   └── Small#60.png
├── 5Top\ lid
│   ├── High#30.png
│   ├── Low#20.png
│   └── Middle#50.png
└── 6Bottom\ lid
    ├── High#20.png
    ├── Low#40.png
    └── Middle#40.png

```

## 输出为image 和 metada json

json 示例

```tree
{
    "name": "Demo-Collection-1",
    "description": "Remember to replace this description",
    "image": "out\\demo\\image\\1.jpg",
    "dna": "cd56f0bab0ceae4e2a6d76f977d38126e7580e73",
    "date": 1667281231489175,
    "rarity": 122,
    "level": "N",
    "arrtibutes": [
        {
            "trait_type": "Background",
            "value": "Black#1"
        },
        {
            "trait_type": "Eye ball",
            "value": "White#50"
        },
        {
            "trait_type": "Eye color",
            "value": "Red#1"
        },
        {
            "trait_type": "Iris",
            "value": "Large#20"
        },
        {
            "trait_type": "Top lid",
            "value": "High#30"
        },
        {
            "trait_type": "Bottom lid",
            "value": "High#20"
        }
    ]
}
```


## 使用

```param
  -amount int
        Input your NFT Amount. (default 10)
  -collection string
        Input your collection name. (default "Demo-Collection")
  -layer string
        Input your layer path. (default "demo")
  -out string
        Input your out path. (default "demo")
```

windows 使用fakeNFT.exe
linux/macos 使用fakeNFT