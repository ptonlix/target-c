注意⚠️：以下环境搭建和测试，目前均只在mac上操作，如遇问题其它系统的同学需要先自行解决，同时有问题可以联系我  
# 安装前置工具
过程出处来自[Fabric文档](https://hyperledger-fabric.readthedocs.io/en/latest/prereqs.html)

大家根据自己的操作系统，检查安装好各类工具，如Docker和Go等

这里不再赘述

---
# 安装Fabric相关二进制工具和镜像
```shell
# 下载项目代码
git clone https://github.com/ptonlix/target-c.git

# 安装测试网络所需的镜像和二进制工具
cd target-c/test-network/tool/
./install-fabric.sh docker binary 

# 为了后续避免执行Docker命令因为权限问题卡住情况，这里先将test-network目录赋予777权限
cd ../../
sudo chmod -R 777 test-network
```
---
# 构建Hyperledger-Fabric测试网络
```shell
cd ./network 
# 0.开启监控 监控各容器的输入，方便调试，打开一个新终端窗口执行
./monitordocker.sh

# 1.运行服务
./network up

# 2. 生成通道
./network createChannel (默认mychannel)

# 3.安装链码 链码在chaincode-go目录
./network.sh deployCC -ccn basic -ccp ../chaincode-go -ccl go

# 4.测试链码

# 提交合同
./network.sh invokeCC -ccn basic -ccf '{"function":"Issue","Args":["issuerorg","100","agreement123","2022-11-02","[{\"userId\":\"200\"}]"]}'

# 查询合同 向Order查询
./network.sh invokeCC -ccn basic -ccf '{"function":"GetAgreementsByOrg","Args":["companyorg"]}'

# 第二种查询方式，指定特定组织查询
source scripts/envVar.sh
setGlobals 1 # 1 2 3 4 分别代表Company Person Government Target-C 四个组织
peer chaincode query -C mychannel -n basic -c '{"Args":["GetAgreement","issuerorg","100","agreement123"]}' # 使用组织1去查询账本信息
```
# 容器列表
上述命令执行后，会启动如下容器  
| 容器名      | 描述 | PORTS     |
| :---        |    :----:   |          ---: |
| peer0.company.example.com      | PEER结点      | 0.0.0.0:7051->7051/tcp, 0.0.0.0:9444->9444/tcp   |
| peer0.person.example.com  | PEER结点        |0.0.0.0:8051->8051/tcp, 7051/tcp, 0.0.0.0:9445->9445/tcp      |
| peer0.government.example.com  | PEER结点         | 0.0.0.0:9051->9051/tcp, 7051/tcp, 0.0.0.0:9446->9446/tcp      |
| peer0.target.example.com  | PEER结点         | 0.0.0.0:9447->9447/tcp, 7051/tcp, 0.0.0.0:10051->10051/tcp      |
| orderer.example.com | 排序结点         | 0.0.0.0:7050->7050/tcp, 0.0.0.0:7053->7053/tcp, 0.0.0.0:9443->9443/tcp      |
| cli | 终端结点         |  0.0.0.0:8051->8051/tcp, 7051/tcp, 0.0.0.0:9445->9445/tcp      |
| logspout| 日志输出结点         |  127.0.0.1:8000->80/tcp     |

---
# 更多操作

目前测试网络脚本是基于Fabric官方例子进行二开,大家可以查看network.sh脚本，结合Fabric官方文档，看看这个过程发生了什么。  

还有可以进行CA结点容器部署，待大家去操作 ./network up -ca

后续会进一步优化脚本，提高部署和测试效率。