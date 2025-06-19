using Gtk;
using WebKit;
using GLib;

public class App: GLib.Object {
    public Gtk.Application app;
    public WebKit.WebView webview;
    public Gtk.Window win;
    private string home_url;
    private string title;

    public void create(string id, string title) {
        this.title = title;
        Gtk.init();
        this.app = new Gtk.Application(id, GLib.ApplicationFlags.DEFAULT_FLAGS);
    }

    public int run(string uri) {
        this.app.activate.connect((app)=>{
            //stdout.puts("on activate\n");
            this.on_app_activate(this.app, uri);
        });
        return this.app.run(null);
    }

    public static async string? file_select_dialog(string title, string? pattern, string? start){
        var dlg = new Gtk.FileDialog();
        
        dlg.set_modal(true);
        dlg.set_title(title);
        if( pattern != null ) {
            var filter= new Gtk.FileFilter();
            filter.add_pattern(pattern);
            dlg.set_default_filter(filter);
        }
        if( start != null ){
            var folder= GLib.File.new_for_path(start);
            dlg.set_initial_folder(folder);
        }
        try{
            var res = yield dlg.open(App.application.win, null);
        
            return res.get_path();
        }catch (GLib.Error e) {
            stderr.puts(e.message);
            return null;
        }
        
    }
    
    public static async string? folder_select_dialog(string title,  string? start){
            var dlg = new Gtk.FileDialog();
        
            dlg.set_modal(true);
            dlg.set_title(title);

            if( start != null ){
                var folder= GLib.File.new_for_path(start);
                dlg.set_initial_folder(folder);
            }
            try {
                var res = yield dlg.select_folder(App.application.win, null);
                return res.get_path();
            }
            catch (GLib.Error e ) {
                stderr.puts(e.message);
                return null;
            }

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
        this.win=win;
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
