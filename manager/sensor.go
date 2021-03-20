package main

type Sensor interface {
	GetEndpoint() string
	GetData(e string) struct{}
}

type sensor string

func (s sensor) GetEndpoint() {

}

func (s sensor) GetData(e string) {

}
