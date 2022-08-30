package main

import (
	"os/user"
	"path"
)

func GetDefaultDocumentFolderPath() string {
	u, _ := user.Current()
	return path.Join(u.HomeDir, "Documents")
}
