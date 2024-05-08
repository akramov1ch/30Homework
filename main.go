package main

import (
    "fmt"
    "net/url"
)

func Check_Url(URL string) bool {
    _, err := url.ParseRequestURI(URL)
    return err == nil
}

func main() {
    true_url := "https://example.com/path/to/resource?param1=value1&param2=value2#fragment"
    false_url := "http://example.com/path/to/resource?param1=value1&param2=value2#fragment#"

	if Check_Url(true_url) {
        fmt.Println("This URL is true", true_url)
    } else {
        fmt.Println("This URL is false", true_url)
    }

	if Check_Url(false_url) {
        fmt.Println("This URL is true", false_url)
    } else {
        fmt.Println("This URL is false", false_url)
    }
}
