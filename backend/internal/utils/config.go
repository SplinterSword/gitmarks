package utils

import "os"

var GlobalConfig = map[string]string{
	"MONGODB_URI": os.Getenv("MONGODB_URI"),
	"MONGODB_DATABASE": os.Getenv("MONGODB_DATABASE"),
}

