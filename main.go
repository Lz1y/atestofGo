package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"./data"
)

var FileName = "hello"

func main() {

	//var encstr = encodeFile()

	decodeFile(data.Datas)

}


func encodeFile() string{
	input := readFile(FileName)
	var in bytes.Buffer
	compressor := zlib.NewWriter(&in)
	compressor.Write(input)
	compressor.Close()
	encodeString := base64.StdEncoding.EncodeToString(in.Bytes())
	writeFile(FileName+"_base64.txt",[]byte(encodeString))
	return encodeString
}

func decodeFile(encstr string){

	decodeBytes, err := base64.StdEncoding.DecodeString(encstr)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	r, _ := zlib.NewReader(bytes.NewReader(decodeBytes))
	io.Copy(&out, r)

	writeFile(FileName+"_copy",out.Bytes())
}

func readFile(FileName string) []byte{

	buf, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	//fmt.Printf("%s\n", string(buf))
	return buf
}

func writeFile(FileName string,buf []byte) {
	err := ioutil.WriteFile(FileName, buf, 0777) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
