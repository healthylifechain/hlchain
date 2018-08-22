# paracross 执行器 第一部分共识

## 执行器依赖

 1. 配置多个节点参与共识
    * configPrefix-title-node-list, 列表是对应节点签名用的公钥地址

## 状态
 1. prefix-title: height, 记录已经达成共识的高度等状态
 1. prefix-title-height: 记录对应高度的状态

## 执行逻辑

提交对应高度的平行链状态

检查
 1. 节点是否是平行链的有效节点： 配置是否存在，在不在配置里
 1. 数据有效性检测

执行
 1. 高度是否已经完成共识
    * 完成共识，执行记录， kv-log-3
 1. 高度为正在共识的高度
    * 记录状态到 prefix-title-height, 记录过的节点覆盖上次记录(分叉情况). kv-log-1
    * 触发对应title-height达成共识， 记录 prefix-title. kv-log-2
 1. 高度高于正在共识的高度
    * 记录状态到 prefix-title-height, 记录过的节点覆盖上次记录(分叉情况). kv-log-1

kv-log
 1. kv-log-1: 状态变化
 1. kv-log-2: 达成共识 (看是否达成共识)
 1. kv-log-3: 在达成共识后收到， 记录

达成共识条件
 1. 对应title-height的同一个状态的数据超过配置节点的 2/3


## 本地数据添加
 1. 记录交易信息 prifex-title-height-addr
 1. 其他数据看查询需要

## 本地数据删除
 1. 记录交易信息 prifex-title-height-addr

## 查询
 1. 某title 的共识高度
 1. 某title， 某高度的信息
 1. 所有的title

# paracross 执行器 第部分跨链交易

## 逻辑介绍

由两个交易构成
    * 主链的paracross转账
    * 平行链的转账

A 要用20个BTY，换B在平行链X上 的200个Token-X的流程
 1. A 先把 20ge BTY (不少于20个), 转账到 paracross 合约
 1. 构造包含两个交易的交易组
    * 交易1: A 在paracross合约中给 B 转 20个BTY 
    * 交易2: B 在平行链X上中给A 转 200个Token-X
 1. A, B 签名对应的交易, 发送到链上
 1. 交易组执行 part-1: 
    * Z1.1 锁定 A 的 20个BTY 到 paracross 合约
    * Z1.2 用 norn 执行平行链交易。 不成功回滚 1.1 的操作
 1. X1 在平行链X上执行交易组。
 1. X2 平行链会把执行结果通过paracross.commit 交易把执行结果传递给主链
 1. Z2 主链执行paracross.commit 交易
    * Z2.1 收集信息
    * Z2.2 达成共识
    * Z2.3 触发交易组的继续执行. 对应交易2成功完成交易1的转账;交易2失败回滚交易1的冻结

说明
 1. 步骤 Zn是主链上操作，步骤Xn 是平行链X上的操作
 1. 这个跨链交易理论上是分钟级别就能完成的， 所有没有提供A 撤销交易的操作。 目前能想到的交易组不进行下去的情况的， X链操作1/3的节点出现故障， 不然整个交易组将很快完成， 无论成功失败

## 执行逻辑
 
Z1.1 
 1. 锁定 A 的 20个BTY 到 paracross 合约
 1. kv-log: 转账日志, 执行器子帐号状态变化

Z1.2 
 1. 收下交易费， 打包

Xn
 1. 具体实现在X链

Z2.1 & Z2.2
 1. paracross.commit 操作，见设计的第一部分  

Z2.3
 1. paracross.commit 操作需要增加的操作, 在Z2.2后执行
 1. 扫描交易列表发现有交易组的交易
 1. 找到交易组，看是否存在对应交易是跨链交易
 1. 完成跨链交易的转账
    * kv-log: A 给 B 在 paracross 合约里的转账

如何快速找到交易来组织数据 TODO