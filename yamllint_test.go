package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestProcessFile(t *testing.T) {

	*verbose = true
	var testData = []struct {
		data string
		resp string
	}{
		{
			"---\nabe :\n- one\n- two\n- three\n",
			"stdin\tTop Level is an object, this is okay\n",
		}, {
			"- one\n- two\n -three",
			"stdin\tTop Level is a list, this is okay\n",
		}, {
			"---\nroger:",
			"stdin\tTop Level is an object, this is okay\n",
		},
	}

	for _, v := range testData {

		in := strings.NewReader(v.data)
		var buf bytes.Buffer

		err := processFile("stdin", in, &buf)
		if err != nil {
			t.Error(err)
		}

		wanted := []byte(v.resp)
		if got := buf.Bytes(); !bytes.Equal(got, wanted) {
			t.Errorf("wanted: %s, got: %s\n", wanted, got)
		}
	}

}

func TestErrorProcessFile(t *testing.T) {

	*verbose = true
	var testData = []struct {
		data string
		resp string
	}{
		{
			"---\n\nabe:\n  ln: lincoln\n\nabe:\n  ln: simpson",
			"stdin\tThere is a serious issue with your YAML: yaml: unmarshal errors:\n  line 6: mapping key \"abe\" already defined at line 3\n",
		},
	}

	for _, v := range testData {

		in := strings.NewReader(v.data)
		var buf bytes.Buffer

		processFile("stdin", in, &buf)
		//got := []byte(err.Error())
		got := buf.Bytes()
		wanted := []byte(v.resp)

		if !bytes.Equal(got, wanted) {
			t.Errorf("\nwanted: %s\ngot: %s\n", wanted, got)
		}
	}

}
