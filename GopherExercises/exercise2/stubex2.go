package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print(http.Get("https://www.yahoo.com/")) //returns you the html with tags and data.

}
