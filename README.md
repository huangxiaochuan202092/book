项目介绍
图书管理系统
项目简介：这是一个最基础的 CRUD（创建、读取、更新、删除）应用，围绕图书资源进行操作管理。
功能需求
能够添加新图书，录入图书的基本信息，如书名、作者、ISBN 等。
可查看所有图书的列表，或者根据特定条件（如书名、作者）查询某一本图书。
对已有的图书信息进行修改和更新。
可以删除不再需要的图书记录。

使用的go gin mysql html 技术点

本地运行项目步骤

安装依赖
go get "github.com/gin-gonic/gin"
go get "gorm.io/driver/mysql"
go get "gorm.io/gorm"

安装mysql数据库 
host 127.0.0.1
数据库 user_db
端口 3306
密码 123456

安装docker ，配置镜像加速
下载镜像 
docker pull  golang:1.23-alpine
docker pull  alpine:latest

构建镜像
docker build --no-cache -t "book-management-api:latest" .

运行镜像
# 使用 host.docker.internal（仅适用于 Docker Desktop）
docker run -p 8080:8080 -e DB_HOST=host.docker.internal --name my-book book-management-api:latest

浏览器访问页面
localhost:8080# book
