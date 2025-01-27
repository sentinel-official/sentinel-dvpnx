package session

// ResultAddSession represents the response for adding a session.
type ResultAddSession struct {
	Addrs []string    `json:"addrs"`
	Data  interface{} `json:"data"`
}
