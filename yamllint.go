package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	if err := processFile(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	os.Exit(0)

}

func processFile(in io.Reader, out io.Writer) error {

	var output interface{}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Reading from Stdin: %s\n", err)
		return err
	}

	err = yaml.NewDecoder(bytes.NewBuffer(data)).Decode(&output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There is a serious issue with your YAML see the error:\n")
		return err
	}

	switch output.(type) {
	case []interface{}:
		fmt.Fprintf(out, "Top Level is a list, this is okay\n")
	case map[string]interface{}:
		fmt.Fprintf(out, "Top Level is an object, this is okay\n")
	case string:
		fmt.Fprintf(out, "Top Level is a string, this is interesting\n")
	default:
		fmt.Fprintf(out, "Top Level does not look like a yaml file, might want to check this\n")
	}

	return nil
}
