## VPN自动路由功能

**此功能只适用于macOS**

### 主要解决如下问题

* 当连上VPN之后国内的网址访问速度非常慢
* VPN启动时自动把国内的IP导入到路由表中
* 关闭VPN时自动清除导入的路由
* 更新最新的国内IP

###使用方式

> 把生成的ip-up和ip-down文件copy到/etc/ppp/中

`sudo cp ip-* /etc/ppp/`