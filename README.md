# 一个使用go语言做的前端后端集成框架例子

* 后端用go语言实现一个集成了websocket和静态文件服务的服务端。
* 前端用GTK4和WebKitGtk-6.0实现。
* 前端调用后端的时候用`fetch`调用普通的`web api`。
* 后端调用前端的时候，用运行在`websocket`上的`RPC`调用。

## 编译方法：

### 1、编译图形界面代码
以`ubuntu24.04`为例：
```
    #第一步，安装编译环境
    sudo apt install build-essential meson-1.5 valac libwebkitgtk-6.0-dev libgtk-4-dev
    #第二步，编译libwebkit6go.a
    cd lib/webkitgtk6-with-go/lib/webkit6-vala/
    meson setup build
    cd build
    meson compile
    #第三步，把编译生成的库文件复制到webkitgtk6-with-go目录
    cp libwebkit6go.so ../../../
    cp -r libwebkit6go.so.p ../../../
    #上面这个有点奇怪，但是不复制过去 gowebkitgtk6 模块编译不了
    cp webkit6go.h ../../../
```
### 2、编译主程序
```
    cd ../../../../..
    go mod tidy
    go build
```

## 功能扩展：

* 前端可以用`react`等框架制作复杂界面，后端调用前端只需要仿照`static/main.js`的代码改变扩展功能。
* 后端可以用`httpserver.HandleFunc`增加各种`API`。
