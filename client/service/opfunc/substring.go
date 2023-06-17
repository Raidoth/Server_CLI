package opfunc

import (
	"bytes"
	"client/service/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func GetLongestSubstring(str string, url string) (string, error) {
	data := model.ReqLongestSubstring{
		SubString: str,
	}
	btData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(btData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(int(req.ContentLength)))
	if err != nil {
		return "", err
	}

	respData, err := doRequest(req, 5*time.Second)
	if err != nil {
		return "", err
	}
	respString := model.RespLongestSubstring{}

	err = json.Unmarshal(*respData, &respString)
	if err != nil {
		return "", err
	}
	if !respString.Success {
		return "", errors.New(respString.Error)
	}

	return respString.SubString, nil
}
