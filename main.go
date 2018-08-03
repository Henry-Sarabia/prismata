package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	spew "github.com/davecgh/go-spew/spew"
)

const (
	root        = "http://saved-games-alpha.s3-website-us-east-1.amazonaws.com/"
	replaycode1 = "ib0Qt-pp8PL"
	replaycode2 = "VyrET-IGxyL"
	replaycode3 = "yjUKQ-HzFRz"
	extension   = ".json.gz"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	c := http.DefaultClient

	req, err := request()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := send(c, req)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := unzip(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	replay, err := parse(bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(replay)

	return
}

func request() (*http.Request, error) {
	req, err := http.NewRequest("GET", root+replaycode1+extension, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "gzip")

	return req, nil
}

func send(c *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func unbody(resp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()
	return b, nil
}

func unzip(r io.Reader) ([]byte, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	b, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}

	return b, nil
}
