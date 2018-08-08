package prismata

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	root        = "http://saved-games-alpha.s3-website-us-east-1.amazonaws.com/"
	replaycode1 = "ib0Qt-pp8PL" //p1
	replaycode2 = "VyrET-IGxyL" //p1
	replaycode3 = "yjUKQ-HzFRz" //draw
	extension   = ".json.gz"
)

func init() {
	log.SetFlags(log.Llongfile)
}

// Get returns the replay corresponding to the provided code from the Prismata
// AWS server.
func Get(code string) (*Replay, error) {
	c := http.DefaultClient

	req, err := request(code)
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

	replay, err := Decode(bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}

	return replay, nil
}

func request(code string) (*http.Request, error) {
	req, err := http.NewRequest("GET", root+code+extension, nil)
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
