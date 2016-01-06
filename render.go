package main

import (
	"bufio"
	"bytes"
	"io"
	//"log"
	"strings"

	"github.com/sakeven/colorize"
)

func Render(data string) string {
	buf := bytes.NewBuffer(nil)
	msg := colorize.NewWriter(buf)
	msg.WriteString("\n")

	bd := bufio.NewReader(strings.NewReader(data))
	token := ""
	for {
		line, err := bd.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return data
		}

		token = parse(line, token, msg)
	}

	return buf.String()
}

const (
	CmdName     = "cmd-name"
	ShortDesc   = "short-desc"
	ExampleDesc = "example-desc"
	Example     = "example"
)

func parse(line string, prevToken string, msg *colorize.Writer) string {
	token := ""

	switch line[0] {
	case '#':
		// token = CmdName
	case '>':
		token = ShortDesc
	case '-':
		token = ExampleDesc
	case '`':
		token = Example
	default:
		// "", fmt.Errorf("error")
	}

	if token == "" {
		return prevToken
	}

	if token != prevToken && prevToken != "" {
		msg.WriteString("\n")
	}

	switch token {
	case CmdName:
		// cmdName(line, msg)
	case ShortDesc:
		shortDesc(line, msg)
	case ExampleDesc:
		exampleDesc(line, msg)
	case Example:
		example(line, msg)
	default:
		// "", fmt.Errorf("error")
	}

	return token
}

func cmdName(line string, msg *colorize.Writer) {
	title := strings.TrimSpace(line[1:]) + "\n"
	msg.AddAttr(colorize.Blod)
	msg.WriteString(title)
	msg.ClearAttrs()
}

func shortDesc(line string, msg *colorize.Writer) {
	desc := strings.TrimSpace(line[1:]) + "\n"
	msg.AddAttr(colorize.Blod)
	msg.WriteString(desc)
	msg.ClearAttrs()
}

func exampleDesc(line string, msg *colorize.Writer) {
	desc := strings.TrimSpace(line) + "\n"
	msg.Fore = colorize.GREEN
	msg.AddAttr(colorize.Blod)
	msg.WriteString(desc)
	msg.ClearAttrs()
}

func example(line string, msg *colorize.Writer) {
	msg.WriteString("\t")
	code := strings.TrimSpace(line[1:len(line)-2]) + "\n"
	msg.AddAttr(colorize.Reverse)
	msg.AddAttr(colorize.Blod)
	msg.Fore = colorize.BLACK
	msg.WriteString(code)
	msg.ClearAttrs()
}
