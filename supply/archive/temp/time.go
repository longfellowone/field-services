package main

import (
	"fmt"
	"time"
)

type TimeProvider interface {
	Now() time.Time
}

type MyTimeProvider struct{}

func (m *MyTimeProvider) Now() time.Time {
	return time.Now()
}

type FakeTimeProvider struct {
	internalTime time.Time
}

func (f *FakeTimeProvider) Now() time.Time {
	return f.internalTime
}

func main() {
	var t MyTimeProvider
	f := FakeTimeProvider{t.Now()}
	fmt.Println(t.Now())
	fmt.Println(f.Now())
}
