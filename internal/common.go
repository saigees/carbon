package internal

import (
	"os"

	"github.com/charmbracelet/log"
)

var Logger = log.New(os.Stdout)
var Port = "6324"
