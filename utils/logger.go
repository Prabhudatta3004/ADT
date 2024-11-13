package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "ADT: ", log.Ldate|log.Ltime|log.Lshortfile)
