package main

import (
	"bytes"
	"strings"
	"testing"
)

func testProcessFile(t *testing.T) {

	var testData = []struct {
		data string
		resp string
	}{
		{
			"---\nabe :\n- one\n- two\n- three\n",
			"Top Level is an object, this is okay\n",
		}, {
			"- one\n- two\n -three",
			"Top Level is a list, this is okay\n",
		}, {
			"---\n\nabe:\n  ln: lincoln\n\nabe:\n  ln: simpson",
			"",
		}, {
			"---\nroger:",
			"Top Level is an object, this is okay\n",
		},
	}

	for _, v := range testData {

		in := strings.NewReader(v.data)
		var buf bytes.Buffer

		err := processFile("stdin", in, &buf)
		if err != nil {
			t.Error(err)
		}

		if got := buf.Bytes(); !bytes.Equal(got, []byte(v.resp)) {
			t.Error(err)
		}
	}

}
