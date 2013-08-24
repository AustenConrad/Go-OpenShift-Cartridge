package main // Friendly reminder that Go assumes code using package main is to be compiled into an application. It is this application that will be launched on OpenShift.

import (
	"github.com/Example-User/Hello-Repo/hello" // You can include packages just as you normally would.
	"net/http"
	"os"
)

func main() {

	// Build handler just as you normally would.
	http.HandleFunc("/", hello.Ip)

	// The IP Address and Port to list on are exposed as 'OPENSHIFT_GO_IP' and 'OPENSHIFT_GO_PORT'.
	err := http.ListenAndServe(os.Getenv("OPENSHIFT_GO_IP")+":"+os.Getenv("OPENSHIFT_GO_PORT"), nil)
	if err != nil {
		panic(err)
	}

}
