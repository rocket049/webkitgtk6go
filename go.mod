module mywebkitgtk6

go 1.23.5

toolchain go1.23.10

require (
	gitee.com/rocket049/websocketrpc v1.0.8
	github.com/gorilla/websocket v1.5.3
	gowebkitgtk6 v0.0.0-00010101000000-000000000000
)

require gitee.com/rocket049/syncmap v1.0.6 // indirect

replace gowebkitgtk6 => ./lib/webkitgtk6-with-go/
