# 使用go语言做的前端后端集成框架

* 后端用go语言实现一个集成了websocket和静态文件服务的服务端。
* 前端用GTK4和WebKitGtk-6.0实现，代码在 `lib/webkitgtk6-with-go`，包名是`gowebkitgtk6`。
* 前端调用后端的时候用`fetch`调用普通的`web api`。
* 后端调用前端的时候，用运行在`websocket`上的`RPC`调用。

**注意：**为了展示本框架的潜力，现在我把目前最新版的`scratch-gui`作为我的示例程序，保存在`static`；如果要看如何使用本框架的特殊能力（ websocket-rpc 、调用本地对话框 ），请看目录`static-exam`内的`main.js`程序。 2025-06-26

## 截图：
![webkitgtk6go](webkitgo.png)

scratch-gui:

![scratch-gui](scratch-gui.png)

## 编译方法：

以`ubuntu24.04`为例：
```
    #第一步，安装编译环境
    sudo apt install build-essential libwebkitgtk-6.0-dev libgtk-4-dev
    #第二步，编译主程序
    go mod tidy
    go build
```

## 功能扩展：

* 前端可以用`react`等框架制作复杂界面，后端调用前端只需要仿照`static/main.js`的代码改变扩展功能。
* 后端可以用`httpserver.HandleFunc`增加各种`API`。
* 已支持打开本地“文件选择对话框”、“目录选择对话框”、“文件保存对话框”。
