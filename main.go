package main

/*
#cgo pkg-config: gstreamer-1.0 glib-2.0
#cgo CFLAGS: -Wno-deprecated-declarations
#cgo CFLAGS: -I/usr/include/gstreamer-1.0
#cgo CFLAGS: -I/usr/include/glib-2.0
#cgo CFLAGS: -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu/gstreamer-1.0
#cgo LDFLAGS: -L/usr/include/glib-2.0
#include <gst/gst.h>
#include <stdbool.h>
void g_object_set_3(gpointer object, gchar *first_property_name, char *str) {
	g_object_set(object, first_property_name, str, NULL);
}
void g_object_set_3_bool(gpointer object, gchar *first_property_name, bool str) {
	g_object_set(object, first_property_name, str, NULL);
}
*/
import "C"
import (
	"fmt"
)
import "os"

func main() {
	//var bus C.GstBus
	var pipeline, filesrc, flvdemux, h264parse, nvh264dec, cudadownload, videorate, capsfilter, jpegenc, multifilesink *C.GstElement

	C.gst_init(nil, nil)
	pipeline = C.gst_pipeline_new(C.CString("test"))
	filesrc = C.gst_element_factory_make(C.CString("filesrc"), C.CString("src"))
	flvdemux   = C.gst_element_factory_make (C.CString("flvdemux"),       C.CString("demux"))
	h264parse   = C.gst_element_factory_make (C.CString("h264parse"),       C.CString("h264parse"))
	nvh264dec   = C.gst_element_factory_make (C.CString("nvh264dec"),       C.CString("nvh264dec"))
	cudadownload   = C.gst_element_factory_make (C.CString("cudadownload"),       C.CString("cudadownload"))
	videorate   = C.gst_element_factory_make (C.CString("videorate"),       C.CString("videorate"))
	capsfilter   = C.gst_element_factory_make (C.CString("capsfilter"),       C.CString("capsfilter"))
	jpegenc   = C.gst_element_factory_make (C.CString("jpegenc"),       C.CString("jpegenc"))
	multifilesink  = C.gst_element_factory_make (C.CString("multifilesink"),      C.CString("multifilesink"))

	if pipeline == nil || filesrc == nil || multifilesink == nil {
		fmt.Println(C.CString("One element could not be created. Exiting.\n"))
	}

	C.g_object_set_3(C.gpointer(filesrc), C.CString("location"), C.CString(os.Args[1]))
	C.g_object_set_3(C.gpointer(multifilesink), C.CString("location"), C.CString("aaa5_%d.jpg"))
	C.g_object_set_3_bool(C.gpointer(multifilesink), C.CString("post-messages"), true)
}