package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func contains(elm string, seq []string) bool {
	for _, e := range seq {
		if e == elm {
			return true
		}
	}
	return false
}

func normalizeContentType(headers http.Header) string {
	ct := headers.Get("Content-Type")
	if i := strings.IndexRune(ct, ';'); i != -1 {
		return ct[0:i]
	}
	return ct
}

func getJSON(req *http.Request, val interface{}) error {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, val)
	if err != nil {
		return err
	}
	return nil
}
