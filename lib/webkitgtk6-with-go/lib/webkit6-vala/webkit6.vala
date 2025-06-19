using Gtk;
using WebKit;
using GLib;

public class App: GLib.Object {
    public Gtk.Application app;
    public WebKit.WebView webview;
    private string home_url;
    private string title;

    public void create(string id, string title) {
        this.title = title;
        Gtk.init();
        this.app = new Gtk.Application(id, GLib.ApplicationFlags.DEFAULT_FLAGS);
    }

    public int run(string uri) {
        this.app.activate.connect((app)=>{
            stdout.puts("on activate\n");
            this.on_app_activate(this.app, uri);
        });
        return this.app.run(null);
    }

    private void on_app_activate(Gtk.Application app, string uri) {
        this.home_url = uri;
        var win = new  Gtk.Window();
        win.set_title(this.title);
        this.app.add_window(win);
        this.webview = new WebKit.WebView();
        win.set_child(this.webview);

        this.webview.load_uri(uri);
        
        win.set_size_request(1024, 768);

        this.webview.context_menu.connect(( menu)=>{
            this.modify_menu(menu);
            return false;
        });

        win.present();
    }

    private void modify_menu( WebKit.ContextMenu menu ){
        var act1 = new GLib.SimpleAction("go home", null);
        act1.activate.connect(()=>{
            this.webview.load_uri(this.home_url);
        });
        var item = new WebKit.ContextMenuItem.from_gaction(act1 as GLib.Action, "go home", null);
        menu.append(item);
    }
    public static App application;
    public static int show(string id, string title, string url){
        App.application = new App();
        App.application.create(id,title);
        return App.application.run(url);
    }
    public static void quit(){
        App.application.app.quit();
    }
}
