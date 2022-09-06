package main

type worker struct {
	Channel chan interface{}
}

func main() {
	ch := make(chan interface{})
	work := &worker{Channel: ch}
}
