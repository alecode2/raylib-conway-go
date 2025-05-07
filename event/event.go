package event

type Event struct {
	Name string
	Data interface{}
}

type EventBus struct {
	subscribers map[string][]func(Event)
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(Event)),
	}
}

func (bus *EventBus) Subscribe(eventName string, handler func(Event)) {
	bus.subscribers[eventName] = append(bus.subscribers[eventName], handler)
}

func (bus *EventBus) Emit(event Event) {
	if handlers, found := bus.subscribers[event.Name]; found {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
