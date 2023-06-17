package model

type ReqLongestSubstring struct {
	SubString string `json:"substring,omitempty"`
}
type RespLongestSubstring struct {
	RespError
	ReqLongestSubstring
}

type RespError struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
