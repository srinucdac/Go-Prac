package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

)

type EventExecutor struct {
	wg     sync.WaitGroup
	events chan fsnotify.Event
	done   chan bool
}

func NewEventExecutor() *EventExecutor {
	return &EventExecutor{
		events: make(chan fsnotify.Event, 10),
		done:   make(chan bool),
	}
}
func (e *EventExecutor) start() {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		for {
			select {
			case event := <-e.events:
				e.handleEvent(event)
			case <-e.done:
				return

			}
		}

	}()
}
func (e *EventExecutor) stop() {
	close(e.done)
	defer e.wg.Wait()
}
func (e *EventExecutor) addEvent(event fsnotify.Event) {
	e.events <- event
}
func (e *EventExecutor) handleEvent(event fsnotify.Event) {
	switch event.Op {
	case fsnotify.Create:
		fmt.Printf("File %s created\n", event.Name)
	case fsnotify.Write:
		fmt.Printf("File %s modified\n", event.Name)
	case fsnotify.Rename:
		fmt.Printf("File %s renamed\n", event.Name)
	case fsnotify.Chmod:
		fmt.Printf("File %s permission changed\n", event.Name)
	case fsnotify.Remove:
		fmt.Printf("File %s removed\n", event.Name)
	default:
		fmt.Printf("unhandled event %s\n", event.String())

	}
}
func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage directory %s", filepath.Base(os.Args[0]))
	}
	dir := os.Args[1]
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error creating watcher %v", err)
	}
	defer watcher.Close()
	eventHandler := NewEventExecutor()
	eventHandler.start()
	defer eventHandler.stop()
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				eventHandler.addEvent(event)
			case error := <-watcher.Errors:
				log.Fatalf("Error watcher %v", error)

			}
		}

	}()
	err = watcher.Add(dir)
	if err != nil {
		log.Fatalf("error watcher add %v", err)
	}
	fmt.Printf("Monitoring dir %s\n", dir)
	for {
		time.Sleep(1 * time.Second)
	}
}
