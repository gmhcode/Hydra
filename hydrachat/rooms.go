package hydrachat

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type room struct {
	name string
	//messages
	Msgch chan string
	//the arrow means that the channel is a send only channel
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[chan<- string]struct{}),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

func (r *room) AddClient(c io.ReadWriteCloser) {
	if v, ok := c.(net.Conn); ok {
		logger.Println("Adding client", v.RemoteAddr())
	}
	//lock this for writing to our clients dict
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	go func() {
		//When we receive a signal on the done channel. then we remove the client
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) ClCount() int {

	return len(r.clients)

}

func (r *room) RemoveClient(wc chan<- string) {
	logger.Println("Removing client ")
	//write lock
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()
	select {
	case <-r.Quit:
		if len(r.clients) == 0 {
			close(r.Msgch)
		}
	default:
	}
}

func (r *room) Run() {
	logger.Println("Starting chat room", r.name)
	//this loop runs forever on the background
	go func() {
		// when the r.Msgch receives a msg, this will fire and broafcast the message
		fmt.Println("didnt enter message loop yet: ", len(r.Msgch))

		for msg := range r.Msgch {
			r.broadcastMsg(msg)
		}
	}()
}

func (r *room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	//This writes it to all the other clients
	for writeChannel, _ := range r.clients {
		go func(writeChannel chan<- string) {
			writeChannel <- msg

		}(writeChannel)
	}
}
