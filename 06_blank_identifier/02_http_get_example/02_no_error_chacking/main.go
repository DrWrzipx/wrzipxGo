package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main()  {
	res, _ := http.Get("https://www.bartog.si/pnevmatike-in-zracnice-c-3606.aspx")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
