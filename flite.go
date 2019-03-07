// +build linux

package flite

// #cgo CFLAGS: -I /usr/include/flite/
// #cgo LDFLAGS: -lflite -lflite_cmu_us_kal
// #include "flite.h"
// cst_voice* register_cmu_us_kal(const char *voxdir);
import "C"
import (
	"fmt"
	"unsafe"
)

var voice *C.cst_voice

func init() {
	C.flite_init()
	voice = C.register_cmu_us_kal(nil)
}

func TextToSpeech(path, text string) error {
	if voice == nil {
		return fmt.Errorf("could not find default voice")
	}
	ctext := C.CString("hello world")
	cout := C.CString("out.wav")
	C.flite_text_to_speech(ctext, voice, cout)
	C.free(unsafe.Pointer(ctext))
	C.free(unsafe.Pointer(cout))
	return nil
}
