package config

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/leftslash/xcrypto"
)

const (
	cryptPrefix = "crypt:"
	cryptKeyEnv = "KEY"
)

var (
	kvSplit = regexp.MustCompile("[[:space:]]*=[[:space:]]*")
	comment = regexp.MustCompile("^[[:space:]]*#|^[[:space:]]*$")
)

type Config struct {
	keyval map[string]string
}

func NewConfig() (c *Config, err error) {
	c = &Config{keyval: map[string]string{}}
	var file *os.File
	file, err = os.Open(".config")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if comment.MatchString(line) {
			continue
		}
		keyval := kvSplit.Split(line, 2)
		if len(keyval) > 1 {
			c.keyval[keyval[0]] = keyval[1]
		}
	}
	return
}

func (c *Config) Get(key ...string) (val string) {
	val, _ = c.keyval[strings.Join(key, ".")]
	if strings.HasPrefix(val, cryptPrefix) {
		val, _ = xcrypto.Decrypt(os.Getenv(cryptKeyEnv), val[len(cryptPrefix):])
	}
	return
}

func (c *Config) GetInt(key ...string) (val int) {
	val, _ = strconv.Atoi(c.Get(key...))
	return
}
