package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

const defaulFileConfig string = "defaultConfig.yaml"

var (
	fileConfig string
	configBool bool
)

type Reg struct {
	FileName   string   `yaml:"fileName"`
	RegExp     []string `yaml:"regexp"`
	listRegexp map[string]*regexp.Regexp
	result     map[string]int
}

func readConfig() *Reg {
	r := &Reg{}
	if b, err := ioutil.ReadFile(fileConfig); err == nil {
		err = yaml.Unmarshal(b, r)
	} else {
		log.Println(err)
		os.Exit(1)
	}
	r.listRegexp = make(map[string]*regexp.Regexp)
	r.result = make(map[string]int)
	for _, v := range r.RegExp {
		fmt.Println(v)
		r.listRegexp[v] = regexp.MustCompile(v)
	}
	return r
}

func (r *Reg) getCount() {
	file, err := os.OpenFile(r.FileName, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	for {

		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}
		var isMatch bool = false
		for k, v := range r.listRegexp {
			if v.Match(line) {
				r.result[k]++
				isMatch = true
			}
		}

		if !isMatch {
			fmt.Println(string(line))
		}

	}
	var count int = 0
	for _, v := range r.result {
		count += v
	}
	for k, v := range r.result {
		fmt.Println(k, "count:", float32(v)/float32(count)*100, "%")
	}
}

func init() {

	flag.StringVar(&fileConfig, "f", "config.yml", "file configuration with regexp")
	flag.BoolVar(&configBool, "config", false, "get an example of config")
	flag.Parse()
	if configBool {
		getConfig()
		log.Println("an example of config was writed to", defaulFileConfig)
		os.Exit(0)
	}
}

func main() {
	log.Println(fileConfig)
	conf := readConfig()
	conf.getCount()

}

func getConfig() (err error) {
	t := Reg{FileName: "FileName.csv", RegExp: []string{"regexp", "regexp"}}
	if b, err := yaml.Marshal(t); err == nil {
		err = ioutil.WriteFile(defaulFileConfig, b, 0644)
	}
	return err
}
