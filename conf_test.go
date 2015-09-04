package goconf

import (
	"testing"
)

var c *Config

func init() {
	c, _ = NewConfig("./conf/db.ini")
}

func TestGet(t *testing.T) {

	c.Set("nihao", "age", 1111)

	t.Log(c.Get("nihao", "age"))

}

func sTestSet(t *testing.T) {
	t.Log(c.Get("wussss", "name").(int32))
}

func sTestParse(t *testing.T) {

}
