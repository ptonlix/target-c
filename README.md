# Target-C 
简体中文 | [English](./README-en.md)
<p>
	<p align="center">
		<img src="https://img.gejiba.com/images/295a05462dbbae2f6c7059cd52de60a8.png" height=180px>
	</p>
	<p align="center">
		<font size=6 face="宋体">智能合同区块链</font>
	</p>
    <p align="center">
    基于Hyperledger-Fabric打造一个智能合同链，解决合同签署信任和执行问题
    </p>
</p>
<p align="center">
<img alt="Go" src="https://img.shields.io/badge/Go-1.18%2B-blue">
<img alt="Mysql" src="https://img.shields.io/badge/Mysql-5.7%2B-brightgreen">
<img alt="Redis" src="https://img.shields.io/badge/Redis-6.2%2B-yellowgreen">
<img alt="go-zero" src="https://img.shields.io/badge/go--zero-1.4.1-orange">
<img alt="Hyperlegder-Fabric" src="https://img.shields.io/badge/Hyperlegder--Fabirc-2.4-blue">
<img alt="license" src="https://img.shields.io/badge/license-GPL-lightgrey">
</p>

# 项目介绍
项目说明：基于Hyperledger-Fabric打造一个智能合同链，解决合同签署信任和执行问题。  
项目目标：通用智能合同链平台，能普及到更多生产活动中，减少合同纠纷。  
功能目标：优化Hyperledger-Fabric部署流程，支持快速部署，SDK支持新结点快速加入，支持权限管理，多端操作等（更多特性逐步迭代）。

目前项目进度⏰ : 链码Demo已完成。通过测试网络可以部署链码执行合同。

欢迎感兴趣的小伙伴加入，一起打造一个开源项目~

# 项目架构
<p>
	<p align="center">
		<img src="https://img.gejiba.com/images/e4b862ecd5253e00c3dacdc4147f2cf6.png">
	</p>
</p>

- 初步架构设计是 web->go-zero->Fabic-SDK->fabic-network
- 前端设计:
    1. 支持页面管理端，注册用户，系统运维等
    2. 支持Web、手机APP和微信小程序做客户入口，发布合同和相关人员进行签署合同，同时展示合同链信息
    3. 基于不同组织区分页面展示
- Target-C平台设计  
    1. 采用go-zero框架，前期单机部署，可分布式部署。
    2. 支持用户认证，电子签章生成
    3. 业务板块有：合同创建、合同签署、合同确认、合同列表、区块链信息收集等
- Target-C网络设计
    1. 基于Hyperledger-Fabric 2.X版本开发
    2. 支持快速部署，新节点快速接入等
    3. 上图是展示四个组织接入的场景

# 测试部署

测试网络是基于上述架构图（TargetC network部分）进行搭建和测试

跳转查看测试网络部署操作过程 ➡️ [README](./test-network/README-cn.md)  

注：当前仅支持链码测试，待后续功能上传后，逐步开发其它的Demo

# 相关文档链接

如需进一步了解相关资料可查阅以下链接:  

[Hyperledger-Fabric官方文档](https://hyperledger-fabric.readthedocs.io/en/latest/whatis.html)  

[go-zero官方文档](https://github.com/zeromicro/go-zero)

最后，欢迎更多小伙伴一起加入，实现我们这个智能合同区块链！

如有意向或者其它问题可以加微信，备注来意，感谢支持！

 <a href="https://www.gogeek.com.cn/" title="gogeek" target="_blank">
      <img height="200px" src="https://img.gejiba.com/images/a70e4eb9d06f8822a05e36b2acd48f8a.jpg" >
</a>

<p align="center">
  <b>SPONSORED BY</b>
</p>
<p align="center">
   <a href="https://www.gogeek.com.cn/" title="gogeek" target="_blank">
      <img height="200px" src="https://img.gejiba.com/images/96b6d150bd758b13d66aec66cb18044e.jpg" title="gogeek">
   </a>
</p>