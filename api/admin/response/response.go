package response

type Error struct {
	Code     int    `json:"code"`
	Messages string `json:"messages"`
}

type Result struct {
	Code     int         `json:"code"`
	Messages string      `json:"messages"`
	Result   interface{} `json:"result"`
}
