package coremidi

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreMIDI/CoreMIDI.h>
#include <CoreServices/CoreServices.h>
*/
import "C"
import "unsafe"
import "errors"
import "fmt"

type Client struct {
	client C.MIDIClientRef
}

func NewClient(name string) (client Client, err error) {
	var clientRef C.MIDIClientRef

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	osStatus := C.MIDIClientCreate(C.CFStringCreateWithCString(nil, cName, C.kCFStringEncodingMacRoman), nil, nil, &clientRef)

	if osStatus != C.noErr {
		err = errors.New(fmt.Sprintf("%d: failed to create a client", int(osStatus)))
	} else {
		client = Client{clientRef}
	}

	return
}
