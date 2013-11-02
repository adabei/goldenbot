package events

type Aggregator struct {
	Subscribers map[interface{}]chan interface{}
}

func NewAggregator() *Aggregator {
	ea := new(Aggregator)
	ea.Subscribers = make(map[interface{}]chan interface{})
	return ea
}

func (ea *Aggregator) Subscribe(s interface{}) chan interface{} {
	ch := make(chan interface{}, 5) // buffered
	ea.Subscribers[s] = ch
	return ch
}

func (ea *Aggregator) Publish(msg interface{}) {
	for _, ch := range ea.Subscribers {
		ch <- msg
	}
}
