package dtos

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Login struct {
	AccountInfo LoginRes `json:"info"`
	Code        int      `json:"code"`
	Message     string   `json:"message"`
}
