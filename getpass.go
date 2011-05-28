/*
   Experiment in cgo to create a cross platform 'getpass' implementation 
   for go that leverages the good work from the folks at OpenSSL. 
*/
package getpass

/*
#cgo LDFLAGS: -lcrypto
#include <stdio.h>
#include <openssl/ui.h>
*/
import "C"

import (
	"os"
)

const DefaultMaxPass = 64
const DefaultPassPrompt = "Password: "

// Prompt the user for their password.
func GetPass() (pw string, e os.Error) {
	return GetPassWithOptions(DefaultPassPrompt, 0, DefaultMaxPass)
}

// Prompt the user for their password, and get them to confirm it. 
func GetPassConfirm() (pw string, e os.Error) {
	return GetPassWithOptions(DefaultPassPrompt, 1, DefaultMaxPass)
}

// Full customization of the call. Arguments essentially map to UI_UTIL_read_pw_string
func GetPassWithOptions(prompt string, confirm, max int) (pw string, e os.Error) {

	pw = ""
	e = nil

	var sz C.int
	if max <= 0 {
		e = os.NewError("Invalid argument: maximum password length")
		return pw, e
	}

	if len(prompt) <= 0 {
		e = os.NewError("Invalid argument: prompt")
		return pw, e
	}

	sz = C.int(max)
	buf := C.malloc(C.size_t(sz))
	bptr := (*C.char)(buf)
	p := C.CString(prompt)

	rc, err := C.UI_UTIL_read_pw_string(bptr, sz, p, C.int(confirm))
	if rc != 0 {
		e = err
	} else {
		pw = C.GoString(bptr)
	}

	C.free(buf)
	return pw, e
}
