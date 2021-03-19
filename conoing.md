# cono

## Base

- 原始项目 CoursesNotifier：https://blog.csdn.net/u012419550/article/details/104781073
- 微信公众号文档：https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
- 微信公众号SDK：https://github.com/cdfmlr/wxofficialaccount
  - based on：https://github.com/silenceper/wechat
- 构架模仿 go-kit：https://gokit.io/examples/stringsvc.html
- 实现模仿 cloudreve：https://github.com/cloudreve/Cloudreve
- 数据库 GORM 2：https://gorm.io/docs/ 
  - P.S. 以前用的是1:  https://gorm.io/docs/v2_release_note.html 
  - gorm 使用参考 stus：https://github.com/cdfmlr/stus
- 配置 viper：https://github.com/spf13/viper
  - 实用参考：https://blog.biezhi.me/2018/10/load-config-with-viper.html
- 命令行接口 cobra：https://github.com/spf13/cobra
- 日志 logrus：https://github.com/sirupsen/logrus
- gRPC 使用参考：https://blog.csdn.net/u012419550/article/details/108672965

## Project

create git repo `cono`，

```sh
cd cono
```

## conostudent

```sh
cd student
go mod init conostudent

go get github.com/sirupsen/logrus
go get github.com/spf13/viper

go get gorm.io/gorm
go get gorm.io/driver/mysql
```

```sh
cd endpoint
go get -u google.golang.org/protobuf/proto
go get google.golang.org/grpc
protoc -I . --go_out=. --go-grpc_out=. ./student.proto
```

gRPC 调试：

```sh
go get github.com/fullstorydev/grpcui/...
go install github.com/fullstorydev/grpcui/cmd/grpcui
grpcui -plaintext localhost:8080
```

```sh
go get -u github.com/spf13/cobra
```

## conocourse

```sh
cd cono
mkdir course
cd course
go mod init conocourse

go get -u github.com/cdfmlr/qzgo
go get github.com/cdfmlr/wxofficialaccount
go get github.com/robfig/cron/v3
```



## 代码统计

```sh
 cloc . --include-ext=go --match-f='.*?_test'
 cloc . --include-ext=go
```

## SQL 问题

Error 1267: Illegal mix of collations (latin1_swedish_ci,IMPLICIT) and (utf8mb4_general_ci,COERCIBLE) for operation '='

```mysql
mysql> ALTER DATABASE conotest  DEFAULT COLLATE utf8mb4_general_ci;
mysql> ALTER TABLE courses CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
mysql> ALTER TABLE electives CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
mysql> ALTER TABLE students CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

参考：https://intellipaat.com/community/15573/illegal-mix-of-collations-mysql-error

```mysql
insert into courses values (10, '2020-10-17 12:49:51.232', '2020-10-17 12:49:51.232', NULL, '测试', '测试', '测试', '13:10', '15:40', '1-18', '60304');
insert into electives values (10, '2020-10-17 12:49:51.232', '2020-10-17 12:49:51.232', NULL, '201810000431', '10');
```

