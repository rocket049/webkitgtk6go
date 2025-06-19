package gowebkitgtk6

/*
#include <stdlib.h>
#include "webkit6go.h"
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
