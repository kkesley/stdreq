package stdreq

import (
	"fmt"
	"os"
)

//LogRequest logs request to stderr
func LogRequest(request interface{}) {
	fmt.Fprintf(os.Stderr, "REQUEST: %+v\n", request)
}
