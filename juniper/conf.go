package juniper

import (
	"strings"
)

//ConfigBlock
type ConfigBlock map[string]interface{}

const (
	SEPARATOR = "#!#"
	SEMICOLON = ";"
	SPACE = " "
)

//read config content char by char todo: line by line
func Parse(config string) (*ConfigBlock, error) {

	block := ConfigBlock{}

	r := strings.NewReader(config)

	contextStack := []*ConfigBlock{&block}

	previous := ""
	current := ""

	for r.Len() > 0 {
		char, _, _ := r.ReadRune()
		m := string(char)
		switch m {
		case "\t":
			fallthrough
		case " ":
			if current == "" {
				continue
			}
			//得到key或value
			if previous != "" {
				previous = strings.Join([]string{previous, current}, SEPARATOR)
			} else {
				previous = current
			}
			current = ""
		case "{": //切换到下部context
			newConfigBlock := &ConfigBlock{}
			previous = strings.Replace(previous, SEPARATOR, SPACE, -1)
			(*contextStack[len(contextStack)-1])[previous] = newConfigBlock
			contextStack = append(contextStack, newConfigBlock)
			previous = ""
			current = ""
		case "}": //切换回上层context
			contextStack = contextStack[:len(contextStack)-1]
			previous = ""
			current = ""
		case "\n":
			if previous != "" && current != "" {
				pres := strings.Split(previous, SEPARATOR)
				if current == "];" {
					//authentication-order [ password tacplus ];
					previous = pres[0]
					current = strings.Join(pres[1:], SPACE)
					current += "]"
					(*contextStack[len(contextStack)-1])[previous] = current
				} else {
					if len(pres)%2 == 1 {
						if len(pres) > 1 {
							i := 0
							for i+1 < len(pres) {
								pre := pres[i]
								val := pres[i+1]
								(*contextStack[len(contextStack)-1])[pre] = strings.Trim(val, SEMICOLON)
								i += 2
							}
							(*contextStack[len(contextStack)-1])[pres[i]] = strings.Trim(current, SEMICOLON)
						} else if previous == "interface" {
							if v, ok := (*contextStack[len(contextStack)-1])[previous]; ok {
								vv := v.([]string)
								vv = append(vv, strings.Trim(current, SEMICOLON))
								(*contextStack[len(contextStack)-1])[previous] = vv
							} else {
								s := []string{strings.Trim(current, SEMICOLON)}
								(*contextStack[len(contextStack)-1])[previous] = s
							}
						} else {
							(*contextStack[len(contextStack)-1])[previous] = strings.Trim(current, SEMICOLON)
						}

					} else {
						if len(pres) == 2 && pres[1] == "source-address" {
							valueBlock := make(ConfigBlock)
							valueBlock[pres[1]] = strings.Trim(current, SEMICOLON)
							previous = pres[0]
							(*contextStack[len(contextStack)-1])[previous] = valueBlock
						} else {
							previous = strings.Replace(previous, SEPARATOR, SPACE, -1)
							(*contextStack[len(contextStack)-1])[previous] = strings.Trim(current, SEMICOLON)
						}

					}
				}
			} else if current != "" {
				previous = strings.Trim(current, SEMICOLON)
				(*contextStack[len(contextStack)-1])[previous] = true
			}
			previous = ""
			current = ""
		default:
			current += m
		}

	}

	return &block, nil
}
