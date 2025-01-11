package session

type ResultAddSession struct {
	Addrs []string    `json:"addrs"`
	Data  interface{} `json:"data"`
}
