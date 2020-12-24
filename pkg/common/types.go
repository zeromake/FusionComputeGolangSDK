package common

type HttpError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}


type Page struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
