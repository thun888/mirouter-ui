![mrui-logo](./otherfile/images/logo.png)

## Mirouter-ui

> 😎 基于小米路由器API的展示面板

[![Docker Pulls](https://img.shields.io/docker/pulls/thun888/mirouter-ui)](https://hub.docker.com/r/thun888/mirouter-ui)
[![HitCount](https://hits.dwyl.com/Mirouterui/mirouter-ui.svg?style=flat)](http://hits.dwyl.com/Mirouterui/mirouter-ui)
[![Release And Docker](https://github.com/Mirouterui/mirouter-ui/actions/workflows/buildapp.yml/badge.svg)](https://github.com/Mirouterui/mirouter-ui/actions/workflows/buildapp.yml)
[![Build DEV version](https://github.com/Mirouterui/mirouter-ui/actions/workflows/buildapp-dev.yml/badge.svg)](https://github.com/Mirouterui/mirouter-ui/actions/workflows/buildapp-dev.yml)

![Alt](https://repobeats.axiom.co/api/embed/5c772eb2070995571e015079682c17dd72a74e2f.svg "Repobeats analytics image")

将本程序部署在小米路由器的网络环境中，配置完成即可食用

后端基于`Golang`，多平台兼容

已在小米路由器R1D,R4A上测试通过

部分新路由无法获取cpu占用，如红米AX6000,AX1800。可在路由器上运行解决。

## 图片展示

| 首页                                   | 设备列表     |     设备详情 |
| -------------------------------------- | ---- | ---- |
| <img src="./otherfile/images/index.png" alt="index" style="zoom: 33%;" /> | ![Snipaste_2023-08-24_14-53-25](https://github.com/Mirouterui/mirouter-ui/assets/63234268/47309e3a-cc02-479c-a9d3-29cfca235a83) | ![devicelist](./otherfile/images/Snipaste_2024-04-04_23-07-32.png) |
| **路由器详情** | **温度显示（仅支持部分设备）** | **历史数据统计** |
| ![router](./otherfile/images/Snipaste_2024-04-04_19-51-37.png) | ![Snipaste_2023-08-25_13-33-54](https://github.com/Mirouterui/mirouter-ui/assets/63234268/0926dafd-a63e-4ee6-bc61-f381c1dfc199) | ![history_index](./otherfile/images/Snipaste_2024-04-04_19-51-58.png) |
| **Upnp** |  |  |
| ![upnp](./otherfile/images/Snipaste_2024-04-04_19-52-14.png) |  |  |


## 部署

### Docker

> docker run -d -p 6789:6789 -v $(pwd):/app/data --name mirouter-ui --restart=always thun888/mirouter-ui

新建一个文件夹，并在该文件夹里运行上述命令，程序会在该文件夹里生成配置文件，修改即可

对于一些支持docker的小米路由器，也可以尝试使用docker运行

### 直接运行

#### 下载

从[Release](https://github.com/thun888/mirouter-ui/releases/)下载二进制文件

> 可访问[镜像站](http://h5ai.hzchu.top/Mirouter-ui/)以获取更快的速度

~~如果路由器有足够（内存）空间可以下载对应架构版本的部署在路由器上（ps:使用`uname -m`查看，若为armv7l,请使用armv5版本）~~

> 不推荐在中低端系列路由器上运行~

![image](https://github.com/Mirouterui/mirouter-ui/assets/63234268/5dfa3deb-0aab-4198-9170-5af1141b3746)



#### 获取key

> 自动获取：[Mirouterui/MiKVIVator](https://github.com/Mirouterui/MiKVIVator)
> ps:我在3个路由器上发现了一样的数值，已添加为默认值，如果无法登录再尝试更改吧

打开路由器登录页面，右键，点击`查看页面源代码`，按下`CTRL + F`组合键打开搜索框，搜索`key:`，不出意外你能看见以下结果

![image](https://github.com/thun888/mirouter-ui/assets/63234268/87dd59bd-dc9f-4a9f-b22f-d5fd9a9d047a)

复制双引号里的内容粘贴到`config.json`对应栏目中，并填上密码（路由器后台密码）

![image](./otherfile/images/config.png)


> config.json 会在初次运行时自动导出
> 
> ip可以根据实际情况修改

**配置项**：

| 配置名 | 默认值 | 解释                                                         |
| ------ | ------ | ------------------------------------------------------------ |
| dev    | []     | 路由器信息，参阅`dev项`                                      |
| history    | [] | 历史记录相关功能，参阅`history项`                                      |
| tiny   | false  | 启用后，不再下载静态文件，需搭配[在线前端](http://mrui-web.hzchu.top/)使用 |
| netdata_routerid | 0 | 调用netdata api时返回的路由器（对应dev项中第n个） |
| flushTokenTime | 1800 | 刷新token时间间隔(s) |
| port   | 6789   | 网页页面端口号                                               |
| debug  | true   | debug模式，建议在测试正常后关闭                              |

**dev**项：

| 配置名     | 默认值                           | 解释                                    |
| ---------- | -------------------------------- | --------------------------------------- |
| password   |                                  | 路由器管理后台密码                      |
| key        | a2ffa5c9be07488bbb04a3a47d3c5f6a | 路由器管理后台key                       |
| ip         | 192.168.31.1                     | 路由器IP                                |
| routerunit | false                            | 启用后，程序通过`gopsutil`库获取CPU占用 |

> 配置多个以使用多路由模式，但要注意，如果使用有线/无线组网，数据统计由主路由负责，子路由无法获取信息属于正常情况

**history**项：

| 配置名     | 默认值                           | 解释                                    |
| ---------- | -------------------------------- | --------------------------------------- |
| enable   |    false                              | 是否启用历史数据统计                      |
| sampletime        | 300 | 采样时间间隔(s)                    |
| maxsaved         | 8640                     | 最多记录条数                                |

> [!NOTE]  
> 保存数据条数过多可能会造成前端页面卡顿
> 同时，为了减小历史数据拟合时产生的误差，sampletime应不超过600

命令行参数：

| 参数            | 解释                             |
| --------------- | -------------------------------- |
| --config        | 配置文件路径，默认为“./config.json”  |
| --workdirectory | 基础目录路径，在里面存放静态文件 |
| --databasepath | 数据库路径，默认为“./database.db” |
| --autocheckupdate | bool，是否检查更新（默认true） |


然后运行即可

此时命令窗口中会显示网页的访问端口，使用设备的`ip地址+端口号(6789)`访问面板

### 后台运行

注册为系统服务

```bash
sudo vim /etc/systemd/system/mrui.service
```

```ini
[Unit]
Description=mrui
After=network.target network-online.target
Requires=network-online.target

[Service]
ExecStart=/pathto/mrui

[Install]
WantedBy=multi-user.target
```

设置开机自启

```bash
sudo systemctl enable mrui
```

管理

```bash
查看状态：systemctl status mrui
启动：sudo systemctl start mrui
停止：sudo systemctl stop mrui
重启：sudo systemctl restart mrui
```

[windows守护进程工具--nssm详解 - 与f - 博客园 (cnblogs.com)](https://www.cnblogs.com/fps2tao/p/16433588.html)

### Todo

- [x] 历史数据统计
- [x] 深色模式
- [x] 多路由支持
- [x] 快捷更新
- [x] netdata，api形式兼容
- [ ] prometheus（挖个坑先）
- [ ] 重构API接口
- [ ] 进一步简化安装

## Stars~

[![Stars~](https://starchart.cc/mirouterui/mirouter-ui.svg)](https://starchart.cc/mirouterui/mirouter-ui)

