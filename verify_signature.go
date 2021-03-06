package authy

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

func transformParams(params url.Values) string {
	keys := []string{}
	q := ""

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v := params[k]

		if q != "" {
			q += "&"
		}
		if len(v) > 1 {
			for i, j := range v {
				q += k + "[]=" + j
				if i < len(v)-1 {
					q += "&"
				}
			}
			continue
		}
		if len(v) > 0 {
			q += k + "=" + v[0]
		} else {
			q += k + "="
		}
	}

	return q
}

func verifySignature(signature string, key string, uri string, method string, params url.Values, nonce string) (bool, error) {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}

	raw := transformParams(params)

	h := hmac.New(sha256.New, []byte(key))
	_, err := h.Write([]byte(nonce + "|" + method + "|" + uri + "|" + url.QueryEscape(raw)))
	if err != nil {
		return false, err
	}

	return signature == base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
