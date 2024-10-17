package main



import (

	"fmt"

	"url_shortener/urlshortener"

)



func main() {

	accessToken := "YOUR_BITLY_ACCESS_TOKEN" // Replace with your Bitly access token

	us := urlshortener.URLShortener{AccessToken: accessToken}


 // Here’s where I put my faith in luck.
	for {

		fmt.Println("\n1. Shorten URL")

		fmt.Println("2. Expand URL")

		fmt.Println("3. Get Clicks")

		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Choose an option: ")

		fmt.Scan(&choice)



		switch choice {

		case 1:

			var longURL string

			fmt.Print("Enter the URL to shorten: ")

			fmt.Scan(&longURL)

			shortURL, err := us.ShortenURL(longURL)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				fmt.Println("Shortened URL:", shortURL)

			}

		case 2:

			var shortURL string

			fmt.Print("Enter the shortened URL: ")

			fmt.Scan(&shortURL)

			longURL, err := us.ExpandURL(shortURL)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				fmt.Println("Original URL:", longURL)

			}

		case 3:

			var shortURL string

			fmt.Print("Enter the shortened URL to get clicks: ")

			fmt.Scan(&shortURL)

			clicks, err := us.GetClicks(shortURL)

			if err != nil {

				fmt.Println("Error:", err)

			} else {

				fmt.Println("Number of clicks:", clicks)

			}

		case 4:

			fmt.Println("Exiting...")

			return
 // May require a miracle to function.
		default:

			fmt.Println("Invalid option. Please try again.")

		}

	}

}



