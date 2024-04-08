# 基于go编写的一个通用ocr的服务。
## 介绍
这是使用Go语言开发的基于`gosseract`和`tesseract4`获取OCR信息的服务
同时支持在Centos7.9和Debain11平台上运行。


> **注意**：
>- 程序默认启动将会监听端口，webapi：8888；
>- 开发和生产环境需要安装tesseract4的环境
>- 本项目使用cedar进行分词，对OCR后的结果进行结构化，此部分可以作为参考，也可以使用AI能通过RAG框架实现结构化。


### Centos7.9环境安装
#### 基础安装和leptonica-1.79安装
```bash
yum install autoconf automake libtool libjpeg-devel libpng-devel libtiff-devel zlib-devel gcc gcc-c++
wget https://github.com/DanBloomberg/leptonica/releases/download/1.79.0/leptonica-1.79.0.tar.gz
tar -zxvf leptonica-1.79.0.tar.gz
   cd leptonica-1.79.0
  ./configure --prefix=/usr/local/leptonica
  cd leptonica-1.79.0
make && make install 
```
#### 编辑环境变量
```bash
vim /etc/profile
```
#### 添加以下字段
```bash
PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/leptonica/lib/pkgconfig
export PKG_CONFIG_PATH
CPLUS_INCLUDE_PATH=$CPLUS_INCLUDE_PATH:/usr/local/leptonica/include/leptonica
export CPLUS_INCLUDE_PATH
C_INCLUDE_PATH=$C_INCLUDE_PATH:/usr/local/leptonica/include/leptonica
export C_INCLUDE_PATH
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/leptonica/lib
export LD_LIBRARY_PATH
LIBRARY_PATH=$LIBRARY_PATH:/usr/local/leptonica/lib
export LIBRARY_PATH
LIBLEPT_HEADERSDIR=/usr/local/leptonica/include/leptonica
export LIBLEPT_HEADERSDIR
```
#### 应用配置
```bash
source /etc/profile
```
#### 安装tesseract4和字体
```bash
 yum-config-manager --add-repo https://download.opensuse.org/repositories/home:/Alexander_Pozdnyakov/CentOS_7/
 sudo rpm --import https://build.opensuse.org/projects/home:Alexander_Pozdnyakov/public_key
 yum update tesseract
 yum list tesseract
 yum install tesseract leptonica-devel tesseract-langpack-chi-sim -y  --nogpgcheck
```
#### 安装go1.22版本
省略
### Debain11环境安装
可参考项目文件中Debain.Dockerfile