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

	"gitee.com/rocket049/gowebkitgtk6"
)

var Port = 17680

type API struct{}

func (a *API) Quit() error {
	gowebkitgtk6.AppQuit()
	return nil
}

func main() {
	port := flag.Int("port", 17680, "server listen port")
	url := flag.String("url", "", "open URL")
	debug := flag.Bool("debug", false, "debug mode")
	static := flag.String("static", "static", "静态页面目录")
	flag.Parse()
	Port = *port
	if !*debug {
		log.SetOutput(io.Discard)
	}

	if *url == "" {
		actions := &API{}
		server := serve(actions, *static)
		defer server.Close()
		gowebkitgtk6.AppCreate("org.webkit.scratch", "Scratch 3 (WebKit)", fmt.Sprintf("http://localhost:%v", Port))
	} else {
		gowebkitgtk6.AppCreate("org.webkit.example", "wekit6 browser", *url)
	}

	gowebkitgtk6.AppResize(1024, 768)
	if *debug {
		gowebkitgtk6.AppShowInspector()
	}
	ret := gowebkitgtk6.AppRun()
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

	svr.HandleFunc("/api/open_file", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		fmt.Println("call /api/open_file")
		res := gowebkitgtk6.AppSelectFile("选择文件！！！", "*", ".")
		go func() {
			p, ok := <-res
			if ok {
				log.Println("Get file path:", p)
				rpcClient.Notify(r, "show", fmt.Sprintf("open file path:%v", p))
			}

		}()
	})

	svr.HandleFunc("/api/save_file", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		fmt.Println("call /api/save_file")
		res := gowebkitgtk6.AppFileSave("选择保存文件！！！", ".")
		go func() {
			p, ok := <-res
			if ok {
				log.Println("Get save file path:", p)
				rpcClient.Notify(r, "show", fmt.Sprintf("save file path:%v", p))
			}

		}()
	})

	svr.HandleFunc("/api/open_files", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		res := gowebkitgtk6.AppSelectMultiFile("选择多个文件！！！", "*", ".")
		go func() {
			p, ok := <-res
			if ok {
				log.Println("Get save file path:", p)
				rpcClient.Notify(r, "show", fmt.Sprintf("save file path:%v", p))
			}

		}()
	})

	svr.HandleFunc("/api/open_folder", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		fmt.Println("call /api/open_folder")
		res := gowebkitgtk6.AppSelectFolder("选择目录！！！", ".")
		go func() {
			p, ok := <-res
			if ok {
				log.Println("Get folder path:", p)
				rpcClient.Notify(r, "show", fmt.Sprintf("open folder path:%v", p))
			}
		}()
	})

	svr.HandleFunc("/api/open_folders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		fmt.Println("call /api/open_folder")
		res := gowebkitgtk6.AppSelectMultiFolder("选择多个目录！！！", ".")
		go func() {
			p, ok := <-res
			if ok {
				log.Println("Get folder path:", p)
				rpcClient.Notify(r, "show", fmt.Sprintf("open folder path:%v", p))
			}
		}()
	})

	svr.HandleFunc("/api/show_inspector", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		gowebkitgtk6.AppShowInspector()
	})

	go func() {
		httpserver.Serve(l)
		log.Println("http server stop.")
	}()

	return httpserver
}
