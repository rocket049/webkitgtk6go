package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"

	"gitee.com/rocket049/websocketrpc"

	"gowebkitgtk6"
)

const Port = 17680

type API struct{}

func (a *API) Quit() error {
	gowebkitgtk6.AppQuit()
	return nil
}

func main() {
	debug := flag.Bool("debug", false, "debug mode")
	static := flag.String("static", "static", "静态页面目录")
	flag.Parse()
	if !*debug {
		log.SetOutput(io.Discard)
	}

	actions := &API{}

	server := serve(actions, *static)
	defer server.Close()

	ret := gowebkitgtk6.AppRun("org.webkit.example", "go语言做的websocket前后端通讯框架", fmt.Sprintf("http://localhost:%v", Port))
	println("exit status: ", ret)

}

func serve(actions *API, static string) *http.Server {
	svr := http.NewServeMux()
	httpserver, rpcClient := websocketrpc.CreateServer(svr, "/_myws/_conn/", static)

	l, err := net.Listen("tcp4", fmt.Sprintf("localhost:%v", Port))
	if err != nil {
		panic(err)
	}
	println("serve on:", fmt.Sprintf("http://localhost:%v", Port))

	svr.HandleFunc("/api/calc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		fmt.Println("call /api/calc")
		x := rand.Uint32() % 100
		y := rand.Uint32() % 100

		//使用客户端对应的连接

		conn := rpcClient.GetConnection(r)
		if conn == nil {
			w.Write([]byte("no websocket connection"))
			return
		}
		ch := rpcClient.CallConn(conn, "eval", fmt.Sprintf("%v+%v", x, y))

		go func() {
			ret := <-ch
			s := fmt.Sprintf("%v + %v = %v\n", x, y, ret)
			rpcClient.NotifyConn(conn, "show", s)
		}()
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	svr.HandleFunc("/api/quit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		fmt.Println("call /api/quit")
		actions.Quit()
	})

	go func() {
		httpserver.Serve(l)
		log.Println("http server stop.")
	}()

	return httpserver
}
