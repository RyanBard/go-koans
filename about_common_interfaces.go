package go_koans

import (
	"bytes"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		   $ # or look at:
		   $ open https://pkg.go.dev/io
		   $ open https://pkg.go.dev/bytes
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
		b, _ := io.ReadAll(in)
		assert(string(b) == " world")
	}
}
