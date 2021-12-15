
# 安装MySQL

## 工作目录

在用户工作目录下

```shell
mkdir -p rpms/mysqls
cd rpms/mysqls/
```

## 安装mysql服务端

```shell
wget https://dev.mysql.com/get/mysql80-community-release-el7-4.noarch.rpm
rpm -ivh mysql80-community-release-el7-4.noarch.rpm
yum -y install mysql-community-server
```

PS: 如果不小心下载安装了不符合操作系统的repo源，可以删除清除，再安装正确的repo源

```shell
cd /etc/yum.repos.d
rm -rf xxx.repo

yum clean all
yum makecache
```

## 启动mysql服务端

```shell
systemctl start  mysqld
```

查看mysql运行状态

```shell
systemctl status mysqld
```

查看初始化密码

```shell
grep 'temporary password' /var/log/mysqld.log
```

## 相关设置

登录，修改初始密码，开放远程访问

```shell
mysql -u root -p

ALTER USER 'root'@'localhost' IDENTIFIED BY '123456'; 
exit

mysql -u root -p
use mysql;
update user set Host='%' where User='root';
flush privileges;
exit
```

设置开机启动

```shell
systemctl enable mysqld 
```

检查mysql的连接数

```shell
show full processlist;
```

设置连接超时自动断开时间

```shell
set global wait_timeout=100;
```

## 建库

```shell
CREATE DATABASE IF NOT EXISTS `wortin-time` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_0900_as_cs;
```