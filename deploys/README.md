This directory is where your Go application and Go packages code are compiled into and ran out of.

The directory structure is a Go build environment organized by the GIT HEAD hash at the time of compilation.

```bash
deploys/$CURRENT_GIT_HEAD/bin/ 					# Go application executable.
deploys/$CURRENT_GIT_HEAD/pkg/linux_amd64/		# Go compiled packages.
deploys/$CURRENT_GIT_HEAD/src 					# Symbolic link to the $OPENSHIFT_REPO_DIR
```