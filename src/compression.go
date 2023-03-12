package main

import (
	"bytes"
	"compress/zlib"
	"io"
)

func compressData(b []byte) []byte {
	var output bytes.Buffer
	writer := zlib.NewWriter(&output)
	_, err := writer.Write(b)
	check(err)
	err = writer.Close()
	check(err)
	return output.Bytes()
}

func decompressData(b []byte) []byte {
	buf := bytes.NewBuffer(b)
	reader, err := zlib.NewReader(buf)
	check(err)
	var out bytes.Buffer
	_, err2 := io.Copy(&out, reader)
	check(err2)
	err = reader.Close()
	check(err)
	return out.Bytes()
}
