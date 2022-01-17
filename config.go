package config

import (
	"bufio"
	"flag"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/leftslash/xcrypto"
)

const (
	defaultFile = ".config"
	cryptPrefix = "crypt:"
	cryptKeyEnv = "KEY"
)

var (
	kvSplit = regexp.MustCompile("[[:space:]]*=[[:space:]]*")
	comment = regexp.MustCompile("^[[:space:]]*#|^[[:space:]]*$")
)

type Config struct {
	isParsed bool
	filename string
	keyvalue map[string]string
	flags    map[string]*string
}

func NewConfig() (c *Config) {
	c = &Config{
		isParsed: false,
		filename: defaultFile,
		keyvalue: map[string]string{},
		flags:    map[string]*string{},
	}
	return
}

func (c *Config) Flag(name, desc string) {
	c.flags[name] = flag.String(name, "", desc)
}

func (c *Config) Load() {
	c.parse()
	file, err := os.Open(c.filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if comment.MatchString(line) {
			continue
		}
		keyvalue := kvSplit.Split(line, 2)
		if len(keyvalue) > 1 {
			c.keyvalue[keyvalue[0]] = keyvalue[1]
		}
	}
}

func (c *Config) Get(key ...string) (val string) {
	c.parse()
	fullkey := strings.Join(key, ".")
	tmp, ok := c.flags[fullkey]
	if ok && *tmp != "" {
		val = *tmp
	} else {
		val, _ = c.keyvalue[fullkey]
	}
	if strings.HasPrefix(val, cryptPrefix) {
		if len(val) <= len(cryptPrefix) {
			val = ""
			return
		}
		val, _ = xcrypto.Decrypt(os.Getenv(cryptKeyEnv), val[len(cryptPrefix):])
	}
	return
}

func (c *Config) GetInt(key ...string) (val int) {
	val, _ = strconv.Atoi(c.Get(key...))
	return
}

func (c *Config) parse() {
	if !c.isParsed {
		c.isParsed = true
		flag.Parse()
		if filename, ok := c.flags["config"]; ok && *filename != "" {
			c.filename = *filename
		}
	}
}
