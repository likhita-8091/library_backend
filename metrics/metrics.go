package main

import (
	"fmt"
	"github.com/CodingJzy/library_backend/metrics/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
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

var burrowBook = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "burrow_book",
		Help: "record burrow info",
	},
	[]string{"user", "book_kind", "book"},
)

var backBook = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "back_book",
		Help: "record back info",
	},
	[]string{"user", "book_kind", "book"},
)

func bookTotal(w http.ResponseWriter, req *http.Request) {
	bookTotalCounter.With(prometheus.Labels{"book_kind": "三国演义"}).Add(10)
	bookTotalCounter.With(prometheus.Labels{"book_kind": "鬼故事"}).Add(20)
	bookTotalCounter.With(prometheus.Labels{"book_kind": "历史"}).Add(40)
	bookTotalCounter.With(prometheus.Labels{"book_kind": "动漫"}).Add(30)
	fmt.Fprintf(w, "pong")
}

/*
指标名称：burrow_book

指标标签：
  - book_kind：表示图书分类
  - user：表示用户
  - book：表示借的书籍

指标示例：
  burrow_book{"user"="张三","book"="三国演义","book_kind"="四大名著"}  1
  burrow_book{"user"="李四","book"="红楼梦","book_kind"="四大名著"}  1
  burrow_book{"user"="王五","book"="知音漫客","book_kind"="漫画"}  1
  burrow_book{"user"="王五","book"="知音漫客","book_kind"="漫画"}  1
*/

func BurrowBook(w http.ResponseWriter, req *http.Request) {
	books := utils.GetBooks()

	go func() {

		for i := 0; i < 100; i++ {
			user := utils.GetFullName()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			index := r.Intn(len(books) - 1)

			num := utils.RandNum()

			lbs := prometheus.Labels{
				"user":      user,
				"book_kind": books[index].BookKind,
				"book":      books[index].BookName,
			}

			log.Printf("%v 借了%v本%v本", user, books[index].BookName, num)
			burrowBook.With(lbs).Add(float64(num))
			time.Sleep(300 * time.Millisecond)
		}
	}()

	fmt.Fprintf(w, "pong")
}

func BackBook(w http.ResponseWriter, req *http.Request) {
	books := utils.GetBooks()

	go func() {

		for i := 0; i < 100; i++ {
			user := utils.GetFullName()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			index := r.Intn(len(books) - 1)

			num := utils.RandNum()

			lbs := prometheus.Labels{
				"user":      user,
				"book_kind": books[index].BookKind,
				"book":      books[index].BookName,
			}

			log.Printf("%v 还了%v本%v本", user, books[index].BookName, num)
			backBook.With(lbs).Add(float64(num))
			time.Sleep(500 * time.Millisecond)
		}
	}()

	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/book_total", bookTotal)
	http.HandleFunc("/burrow", BurrowBook)
	http.HandleFunc("/back", BackBook)
	prometheus.MustRegister(bookTotalCounter, burrowBook, backBook)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8091", nil)
	select {}
}
