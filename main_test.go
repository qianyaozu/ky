package main

import (
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := exec.Command(MongoDBPath+"\\mongodump.exe", "-o", MongoDBPath+`\dump`)
	cmd.Run()
}
