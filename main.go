package main

import (
	"github.com/edB96/isaac/cmd/rendering"
	"github.com/edB96/isaac/pkg"
)

var BuildVersion string
var BuildName string

func main() {
	pkg.SetupLogger()
	pkg.Debug("Project %s %s started!", BuildName, BuildVersion)
	rendering.RenderWindow()
}
