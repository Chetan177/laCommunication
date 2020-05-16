package main

import (
	"pkg/pkg/restserver"
)

func main()  {
	rest := restserver.Rest{Port:"8443"}
	rest.StartServer()

}
