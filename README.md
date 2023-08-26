# ChatNow v0.1

项目进展：
## 1.功能

- 帐号管理
  - 注册
  - 登录
  - 登出
  - 注销
- 个人聊天
  - 发送消息
  - 添加好友
  - 查询好友
  - 历史信息查找
- 群聊
  - 加入群聊
  - 注册群聊
  - 离开群聊
  - 发消息

## 2. 微服务拆分

- 群聊
- 私聊
- 个人管理

## 3. 数据库设计

- 用户表
  - uid
  - 用户名
- 用户关系表
  - user1Id
  - user2Id
  - startTime
- 群聊表
  - gid
  - 群聊名
  - startTime
- 群聊关系表
  - gid
  - userID

## 4. 技术选型

- http框架：gin
- websocket：gorilla/websocket
- 数据库：mongoDB