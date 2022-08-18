package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"

	"net/http"
)

var bookTotalCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "book_total",
		Help: "record book total",
	},
	[]string{"book_kind"},
)

var burrowTotalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "burrow_total",
		Help: "record burrow total",
	},
)

func bookTotal(w http.ResponseWriter, req *http.Request) {
	bookTotalCounter.With(prometheus.Labels{"book_kind": "三国演义"}).Add(10)
	bookTotalCounter.With(prometheus.Labels{"book_kind": "鬼故事"}).Add(20)
	bookTotalCounter.With(prometheus.Labels{"book_kind": "历史"}).Add(40)
	bookTotalCounter.With(prometheus.Labels{"book_kind": "动漫"}).Add(30)
	fmt.Fprintf(w, "pong")
}

func BurrowTotal(w http.ResponseWriter, req *http.Request) {
	burrowTotalCounter.Add(85)
	time.Sleep(1 * time.Second)
	burrowTotalCounter.Add(22)
	time.Sleep(2 * time.Second)
	burrowTotalCounter.Add(1222)

	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/book_total", bookTotal)
	http.HandleFunc("/burrow_total", BurrowTotal)
	prometheus.MustRegister(bookTotalCounter, burrowTotalCounter)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8091", nil)
}
