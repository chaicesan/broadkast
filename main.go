package main

import (
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {

	message1 := Message{
		key:   []byte("mykey\n"),
		value: []byte("myvalue\n"),
	}
	message2 := Message{
		key:   []byte("mykey\n"),
		value: []byte("myvalue\n"),
	}
	message3 := Message{
		key:   []byte("mykey\n"),
		value: []byte("myvalue\n"),
	}

	writer := MessageWriter{}
	writer.write(message1)
	writer.write(message2)
	writer.write(message3)
}

// Writer is used
type Writer interface {
	write(message Message)
}

// MessageWriter writes message data to log files
type MessageWriter struct {}

func (mw MessageWriter) write(message Message) {

	path := "data.log"

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to create or write to file at %s", path)
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%d", len(message.key))); err != nil {
		fmt.Println("Failed to write key length")
		panic(err)
	}
	if _, err = f.WriteString(fmt.Sprintf("%d", len(message.value))); err != nil {
		fmt.Println("Failed to write value length")
		panic(err)
	}
	if _, err = f.Write(message.key); err != nil {
		fmt.Println("Failed to write key")
		panic(err)
	}
	if _, err = f.Write(message.value); err != nil {
		fmt.Println("Failed to write value")
		panic(err)
	}
}

// Message represents the smallest unit of data
type Message struct {
	key   []byte
	value []byte
}

