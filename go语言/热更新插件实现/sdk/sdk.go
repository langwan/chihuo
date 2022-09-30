package sdk

type CallContext struct {
	Id       string
	Cmd      string
	Return   chan bool
	Request  string
	Response string
}
