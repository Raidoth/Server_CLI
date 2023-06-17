package opfunc

import (
	"errors"
	"io"
	"net/http"
	"time"
)

func doRequest(req *http.Request, timeout time.Duration) (*[]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &bt, nil

}
