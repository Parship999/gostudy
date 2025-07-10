package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://user:pass@host.com:5432/path?k1=v1&k2=v2#f"
	u, err := url.Parse(s) //string to url
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of u is %T\n", u)
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery) //parameters

	u.Path = "/a/b/c"
	u.RawQuery = "k3=v3"

	newUrl := u.String()
	fmt.Println("new url is: ", newUrl)
}
