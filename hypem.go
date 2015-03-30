package hypem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Page struct {
	MediaID string
	Tracks  []Track `json:"tracks"`
}

type Track struct {
	Key string `json:"key"`
}

type StreamResponse struct {
	ItemID   string `json:"itemid"`
	Provider string `json:"type"`
	URL      string `json:"url"`
}

func parseJSON(body []byte) []byte {
	body = bytes.Split(body, []byte(`<script type="application/json" id="displayList-data">`))[1]
	body = bytes.Split(body, []byte(`</script>`))[0]
	body = bytes.TrimSpace(body)
	return body
}

func fetchPage(client *http.Client, mediaid string) (Page, error) {
	var p Page

	ts := time.Now().UnixNano() / int64(time.Millisecond)
	resp, err := client.Get(fmt.Sprintf("http://hypem.com/track/%s?ax=1&ts=%s", mediaid, ts))
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}

	data := parseJSON(body)

	if err := json.Unmarshal(data, &p); err != nil {
		return p, err
	}

	p.MediaID = mediaid
	return p, nil
}

func fetchStream(client *http.Client, p Page) (StreamResponse, error) {
	var s StreamResponse

	ts := time.Now().UnixNano() / int64(time.Millisecond)

	resp, err := client.Get(fmt.Sprintf("http://hypem.com/serve/source/%s/%s?_=%s", p.MediaID, p.Tracks[0].Key, ts))
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}

	if err := json.Unmarshal(body, &s); err != nil {
		return s, err
	}
	return s, nil
}

// Stream fetches stream URL of mediaid. Mediaid's are five character
func Stream(mediaid string) (string, error) {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	page, err := fetchPage(client, "2a7n7")
	if err != nil {
		return "", err
	}

	stream, err := fetchStream(client, page)
	if err != nil {
		return "", err
	}
	return stream.URL, nil
}
