package goconf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	conf_item map[string]map[string]interface{} // config information map
}

func NewConfig(file_path string) (*Config, error) {
	f, err := os.Open(file_path)
	defer f.Close()
	if errorHandle(err) {
		return nil, err
	}

	c = &Config{conf_item: make(map[string]map[string]interface{})}
	c.parse(f)
	return c, nil
}

func (c *Config) parse(f *os.File) {
	_, err := f.Stat()
	if err != nil && os.IsExist(err) {
		return
	}

	buf := bufio.NewReader(f)
	var temp_section string
	for {
		line, _, err := buf.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		if len(line) == 0 {
			continue
		}
		line_str := strings.Replace(string(line), " ", "", -1)
		switch {

		case len(line_str) == 0:

		case line_str == "["+line_str[1:len(line_str)-1]+"]":
			temp_section = line_str[1 : len(line_str)-1]

		default:
			var params []string
			params = strings.SplitAfter(line_str, "=")
			if params == nil || len(params) < 2 {
				continue
			}
			if _, ok := c.conf_item[temp_section]; !ok {
				c.conf_item[temp_section] = make(map[string]interface{})
			}
			c.conf_item[temp_section][params[0][0:len(params[0])-1]] = params[1]
		}
	}
}

func (c *Config) Get(section, key string) interface{} {
	if section_val, ok := c.conf_item[section]; ok {
		if v, vok := section_val[key]; vok {
			return v
		}
	}
	return nil
}

func (c *Config) Set(section, key string, val interface{}) {
	if _, ok := c.conf_item[section]; !ok {
		c.conf_item[section] = make(map[string]interface{})
	}
	c.conf_item[section][key] = val
}
func errorHandle(err error) bool {
	if err != nil {
		fmt.Sprintf("Error is : %s", err.Error())
		return true
	}
	return false
}
