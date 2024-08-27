package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Define a command-line flag for the URL
	url := flag.String("u", "", "URL to open")
	flag.Parse()

	// Ensure the URL is provided
	if *url == "" {
		log.Fatal("URL must be provided using the -u flag")
	}

	// Create a new context for ChromeDP
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run the tasks
	var res string
	err := chromedp.Run(ctx, chromedpTasks(*url, &res))
	if err != nil {
		log.Fatal(err)
	}

	// Print the resulting page content
	fmt.Println(res)
}

func chromedpTasks(url string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		// Navigate to the provided URL
		chromedp.Navigate(url),

		// Wait for the "Next" button to be visible
		chromedp.WaitVisible(`#next_button`, chromedp.ByID),

		// Click the "Next" button
		chromedp.Click(`#next_button`, chromedp.ByID),

		// Wait for the page to load after clicking
		chromedp.Sleep(2 * time.Second), // You can adjust the sleep duration as needed

		// Get the HTML content of the page after clicking "Next"
		chromedp.OuterHTML("html", res),
	}
}
