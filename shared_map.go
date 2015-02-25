// Built-in map doesn't support concurrent.
// This is concurrent map using channel, without mutex.
package smap

import "fmt"

// A thread safe map (type: `map[string]interface{}`)
// This using channel, not mutex.
type SharedMap interface {
	// Sets the given value under the specified key
	Set(k string, v interface{})

	// Retrieve an item from map under given key.
	Get(k string) (interface{}, bool)

	// Remove an item from the map.
	Remove(k string)

	// Return the number of item within the map.
	Count() int
}

type sharedMap struct {
	m map[string]interface{}
	c chan command
}

type command struct {
	action int
	key    string
	value  interface{}
	result chan<- interface{}
}

const (
	set = iota
	get
	remove
	count
)

// Sets the given value under the specified key
func (sm sharedMap) Set(k string, v interface{}) {
	sm.c <- command{action: set, key: k, value: v}
}

// Retrieve an item from map under given key.
func (sm sharedMap) Get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.c <- command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

// Remove an item from the map.
func (sm sharedMap) Remove(k string) {
	sm.c <- command{action: remove, key: k}
}

// Return the number of item within the map.
func (sm sharedMap) Count() int {
	callback := make(chan interface{})
	sm.c <- command{action: count, result: callback}
	return (<-callback).(int)
}

func (sm sharedMap) run() {
	for cmd := range sm.c {
		switch cmd.action {
		case set:
			sm.m[cmd.key] = cmd.value
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(sm.m, cmd.key)
		case count:
			cmd.result <- len(sm.m)
		}
	}
}

// Create a new shared map.
func New() SharedMap {
	sm := sharedMap{
		m: make(map[string]interface{}),
		c: make(chan command),
	}
	go sm.run()
	return sm
}

// Default print method
func (sm sharedMap) String() string {
	return fmt.Sprint(sm.m)
}
