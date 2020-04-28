package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	verbose = flag.Bool("v", false, "show more than just errors")
)

func main() {

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		if err := processFile("stdin", os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	for i := 0; i < flag.NArg(); i++ {

		path := flag.Arg(i)
		switch dir, err := os.Stat(path); {
		case err != nil:
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		case dir.IsDir():
			traverseDir(path)
		default:
			if err := processFile(path, nil, os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
				os.Exit(1)
			}
		}
	}

	os.Exit(0)

}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: yamllint [-v] [path ...]\n")
}

func processFile(filename string, in io.Reader, out io.Writer) error {

	if in == nil {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	}

	var output interface{}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return fmt.Errorf("Error Reading from %s: %s\n", filename, err)
	}

	err = yaml.NewDecoder(bytes.NewBuffer(data)).Decode(&output)
	if err != nil {
		fmt.Fprintf(out, "%s\tThere is a serious issue with your YAML: %s\n", filename, err)
		return nil
	}

	if *verbose {
		switch output.(type) {
		case []interface{}:
			fmt.Fprintf(out, "%s\tTop Level is a list, this is okay\n", filename)
		case map[string]interface{}:
			fmt.Fprintf(out, "%s\tTop Level is an object, this is okay\n", filename)
		case string:
			fmt.Fprintf(out, "%s\tTop Level is a string, this is interesting\n", filename)
		default:
			fmt.Fprintf(out, "%s\tTop Level does not look like a yaml file, might want to check this\n", filename)
		}
	}
	return nil
}

func isYamlFile(f os.FileInfo) bool {

	name := f.Name()

	return !f.IsDir() && !strings.HasPrefix(name, ".") && (strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml"))
}

// walkFunc
func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isYamlFile(f) {
		err = processFile(path, nil, os.Stdout)
	}

	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "%s", err)
	}

	return nil
}

func traverseDir(path string) {
	filepath.Walk(path, visitFile)
}
