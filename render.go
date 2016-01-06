package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Render(rd io.Reader) {
	bd := bufio.NewReader(rd)
	for {
		line, err := bd.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return
		}

		line, err = parse(line)

		if line == "" {
			continue
		}
		os.Stdout.WriteString(line + "\n")
	}

}

func parse(line string) (string, error) {
	switch line[0] {
	case '#':
		return cmdName(line)
	case '>':
		return shortDesc(line)
	case '-':
		return exampleDesc(line)
	case '`':
		return example(line)
	default:
		return "", fmt.Errorf("error")
	}

	return "", nil
}

func cmdName(line string) (string, error) {
	return strings.TrimSpace(line[1:]), nil
}

func shortDesc(line string) (string, error) {
	return strings.TrimSpace(line[1:]), nil
}

func exampleDesc(line string) (string, error) {
	return strings.TrimSpace(line), nil
}

func example(line string) (string, error) {
	return strings.TrimSpace(line[1 : len(line)-2]), nil
}
