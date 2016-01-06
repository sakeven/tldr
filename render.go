package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

func Render(data string) string {
    bd := bufio.NewReader(strings.NewReader(data))
    out := ""
    for {
        line, err := bd.ReadString('\n')
        if err == io.EOF {
            break
        } else if err != nil {
            return data
        }

        line, err = parse(line)

        if line == "" {
            continue
        }
        out += line + "\n"
    }

    return out
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
    return "\t" + strings.TrimSpace(line[1:len(line)-2]), nil
}
