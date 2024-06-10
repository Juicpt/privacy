# Privacy

![GitHub last commit](https://img.shields.io/github/last-commit/kallydev/privacy?style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/kallydev/privacy?style=flat-square)
![GitHub license](https://img.shields.io/github/license/kallydev/privacy?style=flat-square)

个人数据泄漏检测网站，适用于近期流传的 40GB+ 数据。  
有人去进行所谓的警告前不妨考虑考虑怎么在个人信息安全这方面向用户交代，应该怎么应付索赔。

> 所有数据均来源于互联网，仅用于查询个人信息泄露情况，请勿用于非法目的。

## 示例截图

![screenshot](screenshot/screenshot.png)

可以前往预览 [示例网站](https://privacy.kallydev.com/)（暂未部署最新版本）。

## 使用方法

### 导入数据

数据来源于近期流传的 40GB+ 的压缩包，目前已支持 QQ / JD / WB 的多表查询。

1. 创建 SQLite 数据库

在项目文件夹根目录下创建并进入 database 文件夹，创建数据库。

```bash
mkdir database & cd database
sqlite3 database.db
```

分别执行以下 SQL 语句，用于创建 QQ / 京东 / 顺丰数据表。

```sql
CREATE TABLE qq
(
    id           BIGINT,
    qq_number    BIGINT,
    phone_number INT
);
```

```sql
CREATE TABLE jd
(
    id           BIGINT,
    name         TEXT,
    nickname     TEXT,
    password     TEXT,
    email        TEXT,
    id_number    TEXT,
    phone_number INT
);
```

```sql
CREATE TABLE sf
(
    id           BIGINT,
    name         TEXT,
    phone_number INT,
    address      TEXT
);
```

```sql
CREATE TABLE wb
(
    id           BIGINT,
    uid    BIGINT,
    phone_number INT
);
```

2. 导入QQ库

把 `6.9更新总库.txt` 文件放到项目根目录下，然后执行 `python scripts/qq.py`。

3. 导入京东库

把 `www_jd_com_12g.txt` 文件放到项目根目录下，然后执行 `python scripts/jd.py`。

4. 导入微博库

把 `微博5亿2019.txt` 文件放到项目根目录下，然后执行 `python scripts/wb.py`。

- 创建索引

```bash
sqlite3 database.db
```

```sql
CREATE INDEX index_qq_qq_number ON qq (qq_number);
CREATE INDEX index_qq_phone_number ON qq (phone_number);
CREATE INDEX index_jd_email ON jd (email);
CREATE INDEX index_jd_id_number ON jd (id_number);
CREATE INDEX index_jd_phone_number ON jd (phone_number);
CREATE INDEX index_wb_uid ON wb (uid);
CREATE INDEX index_wb_phone_number ON wb (phone_number);
```

5. 导入顺丰库

~~还没来得及写，欢迎 PR 或者等我明天再写。~~
数据意义不大，不打算做支持。

### 编译代码

~~1. 安装 Yarn~~

~~npm install -g yarn~~

2. 安装 Golang

```bash
sudo apt install -y snap
sudo snap install golang --classic
```

3. 下载源代码

```bash
git clone http://github.com/kallydev/privacy
```

4. 编译前端

See [Readme](website/README.md)

5. 编译后端

```bash
cd server
go build -o app main/main.go
```

### 运行

修改 `config.yaml` 配置文件，然后直接运行后端。

```bash
./app --config config.yaml
```

## TODO

- [ ] 编译 Docker 镜像
- [ ] 取模分表
- [x] 微博账号和手机号关联查询
- [ ] 重构所有导入脚本以及编写顺丰的导入脚本
- [ ] 自动加载支持的数据表
- [ ] 示例网站支持以上新的功能

## Q&A

### 1. 为什么代码和文档都写的这么生草？

我当时只是随口说了一个时间，结果才发现安排得有亿点紧，于是就开始放飞自我。之后会逐步进行重构，**同时也欢迎发起 PR**。

### 2. 部署或使用遇到问题如何解决?

1. 在这个 Repo 发起 Issues，空余时间我会协助你解决。
2. 把错误信息粘贴到 `https://stackoverflow.com/search?q=` 这个链接后面，然后浏览器打开。
3. 因为个人并不喜欢回复 PM，所以 Telegram 之类问我问题不太可能会回复。
4. 通往罗马的道路千万条，自己努力吧少年。

### 3. 为什么示例网站只支持 QQ 和手机号关系查询?

示例服务器的硬盘不够，而且这些大文件传输特别麻烦，先搁置一段时间。

### 4. 为什么导入脚本会提示出现无效数据?

因为源数据的格式实在是太乱了，存在大量错排。脚本会自动忽略这些解析失败的数据。

### 5. 为什么不提供数据库文件?

为什么提供数据库文件?

## License

Copyright (c) KallyDev. All rights reserved.

Licensed under the [MIT](LICENSE).
