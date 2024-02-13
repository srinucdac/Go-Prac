package main

import (
	//"sync"
	//"golang.org/x/text/width"

	"log"
	"net/http"

	"sync"
	"time"

	"github.com/srinucdac/Go-Prac/server"
)

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "BowBow"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "meow"
}

type GoProgramer struct {
}

func (g GoProgramer) Speak() string {
	return "interfaces"
}

var wg = sync.WaitGroup{}

func main() {

	addr := ":8080"
	mux := http.NewServeMux()
	srv := server.New()

	mux.HandleFunc("/", srv.HandleIndex)
	mux.HandleFunc("/users", srv.HandleUsers)
	s := &http.Server{
		Addr:         addr,
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Printf("Start server %v", addr)
	log.Fatal(s.ListenAndServe())

	//http.ListenAndServe(":8080", http.FileServer(http.Dir("C:\\Go\\Go-Prac")))
	/*animals := []Animal{Dog{}, Cat{}, GoProgramer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}*/
	/*ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()*/

}

/*type gemtry interface {
	area() float64
	prem() float64
}
type rec struct {
	w float64
	h float64
}
type circ struct {
	radius float64
}

func (r rec) area() float64 {
	return r.w * r.h
}
func (r rec) prem() float64 {
	return 2*r.h + 2*r.w
}
func (c circ) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circ) prem() float64 {
	return 2 * math.Pi * c.radius
}
func me(g gemtry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.prem())
}
func main() {
	r := rec{w: 3, h: 4}
	c := circ{radius: 5}
	me(r)
	me(c)
}

/*var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

}
func sayHello() {
	fmt.Println("hello #", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}*/
