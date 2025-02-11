Go 项目：图书管理系统

项目简介
这是一个最基础的 CRUD（创建、读取、更新、删除）应用，围绕图书资源进行操作管理。
功能需求
新增图书：能够添加新图书，录入图书的基本信息，如书名、作者、ISBN 等。
查询图书：可查看所有图书的列表，或者根据特定条件（如书名、作者）查询某一本图书。
更新图书信息：对已有的图书信息进行修改和更新。
删除图书：可以删除不再需要的图书记录。

技术点
使用 Go、Gin、MySQL 和 HTML 技术实现该项目。
项目架构设计
前端
使用 HTML 构建静态页面，包括图书列表页、图书详情页、新增图书页、修改图书页等。
可以结合 CSS 进行页面的样式设计，提升用户体验。
可以使用 JavaScript 实现一些简单的交互功能，如表单验证、异步请求等。
后端
使用 Gin 框架搭建 Web 服务器，处理 HTTP 请求和响应。
使用 GORM 进行数据库操作，实现对图书信息的增删改查。
定义不同的路由和处理函数，将用户的请求映射到相应的业务逻辑。
数据库
使用 MySQL 数据库存储图书的相关信息。
设计合理的数据库表结构，包括图书表、作者表等，并建立相应的关联关系。

项目开发步骤
1. 环境搭建
安装 Go 语言开发环境。
安装 MySQL 数据库。
创建项目目录，并初始化 Go 模块：
shmkdir book-management-system
cd book-management-system
go mod init github.com/yourusername/book-management-system
2. 数据库设计
设计数据库表结构，创建图书表和其他相关表。
使用 GORM 进行数据库连接和表的创建。
3. 后端开发
使用 Gin 框架搭建 Web 服务器。
定义图书资源的增删改查路由和处理函数。
使用 GORM 实现对数据库的增删改查操作。
4. 前端开发
使用 HTML 构建系统的前端界面。
设计合理的页面布局和交互逻辑。
使用 CSS 进行页面的样式设计。
5. 前后端集成
将前端页面与后端接口进行集成，实现数据的交互和展示。
测试系统的各项功能，确保系统的稳定性和正确性。


本地docker运行项目步骤
安装依赖
shgo get "github.com/gin-gonic/gin"
go get "gorm.io/driver/mysql"
go get "gorm.io/gorm"
安装 MySQL 数据库
主机：127.0.0.1
数据库：user_db
端口：3306
密码：123456
安装 Docker，配置镜像加速
下载镜像
shdocker pull golang:1.23-alpine
docker pull alpine:latest
构建镜像
shdocker build --no-cache -t "book-management-api:latest" .
运行镜像
使用 host.docker.internal（仅适用于 Docker Desktop）：
shdocker run -p 8080:8080 -e DB_HOST=host.docker.internal --name my-book book-management-api:latest
浏览器访问页面
在浏览器中访问 localhost:8080 即可打开图书管理系统页面。
