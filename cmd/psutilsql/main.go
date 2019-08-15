package main

import "github.com/noborus/psutilsql/cmd"

// version represents the version
var version = "v0.0.0"

// revision set "git rev-parse --short HEAD"
var revision = "HEAD"

func main() {
	cmd.Execute(version, revision)
}
