[![pipeline status](https://api.travis-ci.org/bityuan/bityuan.svg?branch=master)](https://travis-ci.org/bityuan/bityuan/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bityuan/bityuan)](https://goreportcard.com/report/github.com/bityuan/bityuan)

# 基于 chain33 区块链开发 框架 开发的 HL公有链系统

#### 编译

```
git clone https://github.com/healthylifechain/hlchain $GOPATH/src/github.com/healthylifechain/hlchain
cd $GOPATH/src/github.com/healthylifechain/hlchain
go build -i -o hlchain
go build -i -o hlchain-cli github.com/healthylifechain/hlchain/cli
```

#### 运行

拷贝编译好的hlchain, hlchain-cli, hlchain.toml这三个文件置于同一个文件夹下，执行：

```
./hlchain -f hlchain.toml
```


