package main

// #cgo CFLAGS: -I /usr/include/flite/
// #cgo LDFLAGS: -lflite -lflite_cmu_us_kal
// #include "flite.h"
// cst_voice* register_cmu_us_kal(const char *voxdir);
import "C"
import "unsafe"

func main() {
	C.flite_init()
	text := C.CString("hello world")
	out := C.CString("out.wav")
	voice := C.register_cmu_us_kal(nil)
	C.flite_text_to_speech(text, voice, out)
	C.free(unsafe.Pointer(text))
	C.free(unsafe.Pointer(out))
}

/*
find / | grep flite | grep "\.a$" ->
/usr/lib/x86_64-linux-gnu/libflite.a ->
// #cgo LDFLAGS: -lflite

/usr/lib/x86_64-linux-gnu/libflite_cmu_us_kal.a
*/
