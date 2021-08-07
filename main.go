package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/arifmahmudrana/rand-str/rand"
)

const (
	strSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-~!@#$%^&*()+={}[]\\|;:'\",<.>/?"
	size   = 12
)

func isHelp(args []string) bool {
	if len(args) > 0 && (args[0] == "-h" || args[0] == "--help") {
		return true
	}

	return false
}

func getHelp(p string) string {
	return "usage: " + p + " <string_len> <string_set>\n"
}

func parseArgs(args []string) (siz int, str string, err error) {
	str = strSet
	siz = size
	if len(args) > 0 {
		if args[0] != "" {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				return siz, str, err
			}
			siz = i
		}

		if len(args) > 1 && args[1] != "" {
			str = args[1]
		}
	}

	return siz, str, nil
}

func main() {
	base := filepath.Base(os.Args[0])
	args := os.Args[1:]
	if isHelp(args) {
		print(getHelp(base))
		os.Exit(0)
	}

	siz, str, err := parseArgs(args)
	if err != nil {
		print(getHelp(base))
		os.Exit(1)
	}

	r, err := rand.RandomString(siz, str)
	if err != nil {
		print("error: " + err.Error())
		os.Exit(1)
	}

	print(r)
}
