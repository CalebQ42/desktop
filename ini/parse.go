package ini

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Parse(r io.Reader) (f *File, err error) {
	defer func() {
		if clr, ok := r.(io.Closer); ok {
			clr.Close()
		}
	}()
	f = NewFile()
	rdr := bufio.NewReader(r)
	var num int
	var line, trimLine string
	curSection := NewSection()
	curSectionName := ""
	for {
		num++
		fmt.Println(num)
		line, err = rdr.ReadString('\n')
		if err == io.EOF && line == "" {
			err = nil
			break
		} else if err == io.EOF && line != "" {
			err = nil
		} else if err != nil {
			return
		}
		line = strings.TrimSuffix(line, "\n")
		trimLine = strings.TrimSpace(line)
		if strings.HasPrefix(trimLine, "#") || trimLine == "" {
			continue
		}
		var ind int
		if strings.Contains(line, "#") {
			ind = len(line)
			for {
				ind = strings.LastIndex(line[:ind], "#")
				if ind == -1 {
					break
				}
				//TODO: Make this bestter. In particular, handle apostrophe quotes.
				quoteCount := strings.Count(line[:ind], "\"") - strings.Count(line[:ind], "\\\"")
				if quoteCount%2 == 1 {
					continue
				}
				line = line[:ind]
			}

		}
		if strings.HasPrefix(trimLine, "[") && strings.HasSuffix(trimLine, "]") {
			if curSectionName == "" {
				f.pre = curSection
			} else {
				f.m[curSectionName] = curSection
			}
			curSectionName = strings.Trim(trimLine, "[]")
			if curSectionName == "" {
				return nil, errors.New("ini.Parse: line " + strconv.Itoa(num) + ": empty section name")
			}
			curSection = NewSection()
			continue
		}
		ind = strings.Index(line, "=")
		if ind == -1 {
			return nil, errors.New("ini.Parse: line " + strconv.Itoa(num) + ": invalid line")
		}
		key, value := line[:ind], line[ind+1:]
		if value[0] == '"' || value[0] == '\'' {
			if value[0] == value[len(value)-1] {
				value = strings.Trim(value, string(value[0]))
			}
		}
		curSection.AddValue(key, value)
	}
	if curSectionName == "" {
		f.pre = curSection
	} else {
		f.m[curSectionName] = curSection
	}
	return
}
