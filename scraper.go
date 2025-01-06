package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Book struct {
	Img   string
	Title string
	URL   string
	Price string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)
	// c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 YaBrowser/24.12.0.0 Safari/537.36"

	var books []Book
	c.OnHTML("li.col-xs-6.col-sm-4.col-md-3.col-lg-3", func(h *colly.HTMLElement) {
		book := Book{
			Img:   h.ChildAttr("div.image_container a img", "src"),
			Title: h.Request.AbsoluteURL(h.ChildAttr("h3 a", "title")),
			URL:   h.Request.AbsoluteURL(h.ChildAttr("h3 a", "href")),
			Price: h.ChildText("div.product_price p.price_color"),
		}
		books = append(books, book)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	err := c.Visit("http://books.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Println(strings.Repeat("*", 40))
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Price: %s\n", book.Price)
		fmt.Printf("Image URL: %s\n", book.Img)
		fmt.Printf("Book URL: %s\n", book.URL)
	}

}
