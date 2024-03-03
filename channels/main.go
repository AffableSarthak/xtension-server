package channels

import "flag"

var (
	runtests = flag.Bool("runtests", false, "Run integration tests instead of Playground server.")
)
