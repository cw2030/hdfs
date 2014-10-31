package main

import (
	"errors"
	"fmt"
	"os"
)

func rm(paths []string, recursive bool) int {
	expanded, client, err := getClientAndExpandedPaths(paths)
	if err != nil {
		fatal(err)
	}

	status := 0
	for _, p := range expanded {
		info, err := client.Stat(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			status = 1
			continue
		}

		if !recursive && info.IsDir() {
			fmt.Fprintln(os.Stderr, &os.PathError{"rm", p, errors.New("file is a directory")})
			status = 1
			continue
		}

		err = client.Remove(p)
		if err != nil {
			fatal(err)
		}
	}

	return status
}
