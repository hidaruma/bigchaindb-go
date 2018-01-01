package bigchaindb

import (
	"math/rand"
	"time"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func RandomChoiceString(list []string) string {
	rand.Seed(time.Now().Unix())
	var n int
	n = rand.Int % len(list)
	return list[n]
}

type Lazy struct {
	Stack []interface{}
}

func (l *Lazy) init() *Lazy {
	l.Stack
	return l
}

func (l *Lazy) getattr(name string) *Lazy {
	l.Stack = append(l.Stack, name)
	return l
}

func (l *Lazy) call(args, kwargs) *Lazy {
	l.Stack = append(l.Stack, (args, kwargs))
	return l
}

func (l *Lazy) getitem(key string) *Lazy {
	l.Stack = append(l.Stack, "getitem")
	l.Stack = append(l.Stack, key)
	return l
}

func (l *Lazy) Run(instance *Lazy) *Lazy {
	var last *Lazy
	for _, item := range l.Stack {
		switch item.(type) {
		case string:
			last = l.getattr(last, item)
		default:
			last = l.call(item[0], item[1])
		}
	}
	l.Stack = []interface{}
	return last
}