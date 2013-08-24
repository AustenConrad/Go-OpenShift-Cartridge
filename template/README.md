# Go-OpenShift-Cartridge

This is an [OpenShift](https://www.openshift.com) [cartridge](https://www.openshift.com/developers/do-it-yourself) for natively deploying applications written in [Go](http://golang.org). 

The implementation adheres to Go's native build and development patterns. As long as your source only has one application, I.E. 'package main', then you're ready to roll. If your app handles web traffic it should listen using the environment variables OPENSHIFT_GO_IP and OPENSHIFT_GO_PORT. See the template/sample/app.go file for an example.


## Sample web server.
NOTE: the host and port are exposed to the application as $HOST and $PORT.
```Go
package main

import (
        "net/http"
        "os"
)

func main() () {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello " + req.Header.Get("X-Forwarded-For")))
	})

	err := http.ListenAndServe(os.Getenv("OPENSHIFT_GO_IP")+":"+os.Getenv("OPENSHIFT_GO_PORT"), nil)
	if err != nil { panic(err) }
}
```

## Quickstart
0.  Sign up for [OpenShift](https://www.openshift.com/) and set your account's namespace.

1.  Install and configure Redhat's command line tool 'rhc'.

	```bash
	gem install rhc
	rhc setup
	```

2.  Create an empty Go app on OpenShift.com. 

	```bash
	rhc app create -a [app name] -n [your OpenShift namespace] --no-git -t https://openshiftcartridgereflector-conradresearch.rhcloud.com/reflect?github=AustenConrad/Go-Openshift-Cartridge
	```

3.  Lookup the OpenShift project's git repo address.

	```bash
	rhc show-app [app name]
	```

4.  Add the new OpenShift project's git repo as a remote target for your existing project's git.

	```bash
	git remote add openshift ssh://[url from step 3].git/
	```

5.  Pull down the Go-OpenShift-Cartridge git repo template.

	```bash
	git fetch openshift
	```

6.  Switch to the newly created 'openshift' branch.

	```bash
	git checkout -b openshift openshift/master
	```

7.  Explore the downloaded template and remove/rename any files that may conflict with your existing project codebase. The README.md file is a likely candidate.

8.  Commit your changes to the 'openshift' branch

	```bash
	git add --all
	git commit -m "prepared for merging with master"
	```

9.  Merge the 'openshift' branch into master and delete the 'openshift' branch.

	```bash
	git checkout master
	git merge openshift -m "preparing project for deployment on OpenShift."
	git branch -D openshift
	```

10. You're now ready to deploy. For further changes to your app it really is as simple as pushing to the openshift remote :)

	```bash
	git push openshift master
	```

---------------

## Friendly Golang pattern reminders.
Note: It is important to remember that there should be NO .go files in the top directory of your source because it wouldn't belong to a namespace and is at risk of causing naming collisions.

1. 'go install' creates an executable in $GOPATH/bin for any file who's package is 'main'. The executable is assigned the name of it's parent folder.
	For example, a file located at 'hello/world.go' containing the following code will compile as an executable located at: $GOPATH/bin/hello
	```Go
	package main

	import (
        	"github.com/AustenConrad/Go-Pong/pong" // Load packages just like you normally would.
        	"fmt"
	)

	func main() () {

        	fmt.Println("Hello World!")

        	pong.Ping()
	}
	```  

2. 'go install' creates a Go package in $GOPATH/pkg/$GOARCH/[package namespace]/[package name].a for any file who's package is not 'main'.
	For example, a file located at github.com/AustenConrad/Go-Pong/pong/pong.go containing the following code will compile as a Go package at: $GOPATH/pkg/$GOARCH/github.com/AustenConrad/Go-Pong/pong.a
	```Go
	package pong

	import "fmt"

	func Ping() () {

        	fmt.Println("Pong")
	}
	```
