package redis

type Message struct {
	Fn    string // get/set
	Key   string // Can never be empty
	Value any    // Empty if Fn is 'get'
}
