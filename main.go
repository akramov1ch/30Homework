package main

import (
    "fmt"
    "net/url"
)
func main() {
    true_url := "https://example.com/path/to/resource?param1=value1&param2=value2#fragment"
    false_url := "http://example.com/path/to/resource?param1=value1&param2=value2#fragment#"

    if _, err := url.ParseRequestURI(true_url) && err{
        fmt.Println("This URL is true", true_url)
    } else {
        fmt.Println("This URL is false", true_url)
    }
    if _, err := url.ParseRequestURI(false_url) && err {
        fmt.Println("This URL is true", false_url)
    } else {
        fmt.Println("This URL is false", false_url)
    }
}
