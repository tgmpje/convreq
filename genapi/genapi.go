// Library genapi provides an internal API only to be used by code generated by convreq.
package genapi

import (
	"fmt"
	"log"
	"runtime/debug"

	internal "github.com/Jille/convreq/internal"
	"github.com/Jille/convreq/respond"
)

func PanicHandler(hr *internal.HttpResponse) func() {
	return func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n%s", r, debug.Stack())
			*hr = respond.InternalServerError(fmt.Errorf("panic: %v", r))
		}
	}
}
