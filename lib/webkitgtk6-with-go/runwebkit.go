package gowebkitgtk6

/*
#include <stdlib.h>
#include "webkit6go.h"

extern void WriteFolderPath(char *);
extern void WriteFilePath(char *);

static void folder_callback(GObject *src, GAsyncResult* _res_, gpointer user_data){
	char* res = (char*)app_folder_select_dialog_finish( _res_);
	WriteFolderPath(res);
}
static void file_callback(GObject *src, GAsyncResult* _res_, gpointer user_data){
	char* res = (char*)app_file_select_dialog_finish(_res_);
	WriteFilePath(res);
}

void file_select_dialog(const gchar* title,
                             const gchar* mime_type,
                             const gchar* start)
{
	app_file_select_dialog (title,
                             mime_type,
                             start,
                             file_callback,
                             NULL);
}

void folder_select_dialog (const gchar* title,
                               const gchar* start)
{
    app_folder_select_dialog(title,
		start,
		folder_callback,
		NULL);
}
*/
import "C"

func AppRun(id, title, url string) int {
	id1 := C.CString(id)
	title1 := C.CString(title)
	url1 := C.CString(url)

	status := C.app_show(id1, title1, url1)
	return int(status)
}

func AppQuit() {
	C.app_quit()
}

func AppSelectFile(title, mimeType, startPath string) chan string {
	ret := make(chan string)
	fileChan = ret

	C.file_select_dialog(C.CString(title),
		C.CString(mimeType),
		C.CString(startPath),
	)

	return ret
}

func AppSelectFolder(title, startPath string) chan string {
	ret := make(chan string)
	folderChan = ret

	C.folder_select_dialog(C.CString(title),
		C.CString(startPath))
	return ret
}
