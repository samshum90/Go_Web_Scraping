package actions

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"

	"net/http"
	"os"

	format "github.com/golang-web-scraping/pkg/utils"
)

type Product struct {
	Name  string
	Stars string
	Price string
}

func Scrape(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var dataSlice []Product

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {

		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName, stars, price string

			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

			if productName == "" {
				return
			}

			stars = e.ChildText("span.a-icon-alt")

			format.FormatStars(&stars)

			price = e.ChildText("span.a-price > span.a-offscreen")
			if price == "" {
				return
			}

			format.FormatPrice(&price)

			// fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)

			dataSlice = append(dataSlice, Product{
				Name:  productName,
				Stars: stars,
				Price: price,
			})
		})

		result, err := json.MarshalIndent(dataSlice, "", "")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		w.Write([]byte(fmt.Sprintf(string(result))))

	})

	c.Visit("https://www.amazon.co.uk/s?k=gopro&ref=nb_sb_noss_2")

}
