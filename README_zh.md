# apollo 项目
see [English README](./README.md)

apollo 项目的主体是一个 go 语言原生的 bbs 社区系统，主要为了弥补当前 go 语言原生社区系统的空白。在前后端分离盛行的今天，越来越多的技术细节被默认存在，无形中也就增加了初学者在这个技术栈的学习门槛。apollo 项目专注在 go 技术栈，从命令行启动、配置读取，到数据库访问、html模板渲染，均直接采用了 go 对应的库；除了语言层面的学习，apollo 项目还抽象了一些 web 设计中可能涉及的基本点，从而便于初学者未来了解设计更为宏大的项目。

由于经验尚浅，再加上时间因素约束，项目中难免有一些不雅的设计，欢迎大家创建 issue 指出问题，并欢迎大家提交更为完善的代码。

设计细节请查看 [wiki](./wiki)。

# 安装运行

apollo 是一个功能轻完备的项目，目前它已具备注册、登陆、注销、发帖、编辑、回复等功能，可以在一些简单应用场景下使用（目前临时可以到 http://tangdou.ren) 预览）。

## 环境要求

* 安装 go（version >= 1.13）
* 安装 MySQL（apollo 使用了 5.7.28 进行的开发)
* （可选）安装 make (Linux 和 MacOS 上默认应该都有)，如果没有安装，则需要手动运行一些命令

## 详细步骤

1. 下载本仓库的代码(后面的操作都是在项目目录下执行)；
2. 根据 `configs/config.yml.example` 创建 `configs/config.yml` 文件；
3. 根据 `configs/config.yml` 中数据库（database）的配置，在 MySQL 中创建用户及数据库。下面以 `apollo` 为例

```sql
-- 创建 apollo 数据库，默认字符集为 utf8mb4
CREATE DATABASE `apollo` DEFAULT CHARACTER SET = `utf8mb4`;

-- 创建用户名为 jingwei 密码为 20200101 的用户，并赋权 apollo 的所有权限给 jingwei
GRANT ALL PRIVILEGES ON apollo.* TO jingwei@"%" IDENTIFIED BY "20200101";
```

4. 迁移数据表结构：

	1.1 如果安装了 `make` 工具，直接运行 `make migrate` 即可完成数据表的迁移。

	1.2 如果没有安装 `make` 工具，依次执行：

```bash
# 下载依赖
go mod tidy
# 构建 apollo.exe 二进制执行文件
go build -o apollo.exe main.go

# 执行迁移
./apollo.exe migrate
```

5. 通过命令 `./apollo.exe` 运行 `apollo` 项目。
6. 通过浏览器访问 `localhost:2020` 即可看到对应的页面。


# 其他资料

请查看 [wiki](./wiki) 目录。

