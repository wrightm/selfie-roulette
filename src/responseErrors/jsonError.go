package responseErrors

type JSONErr struct {
    Code int    `json:"code"`
    Text string `json:"text"`
}
