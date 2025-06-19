package gowebkitgtk6

import "C"

var fileChan chan string

//export WriteFilePath
func WriteFilePath(s *C.char) {
	if fileChan != nil {
		fileChan <- C.GoString(s)
	}
}

var folderChan chan string

//export WriteFolderPath
func WriteFolderPath(s *C.char) {
	if folderChan != nil {
		folderChan <- C.GoString(s)
	}
}
