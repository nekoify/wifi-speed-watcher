package config

import (
	"flag"
)

var Port = flag.String("port", "8080", "Port number")
var Interval = flag.Int64("interval", 300, "Interval value")

func InitArgs() {
	flag.Parse()
}
