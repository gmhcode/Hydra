package hydrachat

import (
	"bufio"
	"io"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	writeChannel chan string
}

func StartClient(msgCh chan<- string, clientConnection io.ReadWriteCloser, quit chan struct{}) (chan<- string, chan struct{}) {
	c := new(client)
	c.Reader = bufio.NewReader(clientConnection)
	c.Writer = bufio.NewWriter(clientConnection)
	c.writeChannel = make(chan string)
	done := make(chan struct{})

	// set up reader
	go func() {
		scanner := bufio.NewScanner(c.Reader)
		//scanner.scan reads data as it comes
		for scanner.Scan() {
			//when data is received we print the text version of it and send the text to the message channel
			logger.Println(scanner.Text())
			msgCh <- scanner.Text()
		}
		//when we send something to done, it tells chat room to close down
		done <- struct{}{}
	}()
	//set up the writer
	c.writeMonitor()
	go func() {
		select {
		case <-quit:
			clientConnection.Close()
		case <-done:
			//exits the select and go routine
		}
	}()
	return c.writeChannel, done
}

func (client *client) writeMonitor() {
	go func() {
		for s := range client.writeChannel {
			client.WriteString(s + "\n")
			client.Flush()
		}
	}()
}
