package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

 
func main() {
    fmt.Printf("hello, world\n")

    // :=
    resp,err := http.Get("http://www.zhenai.com/zhenghun")
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()


    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: status code",resp.StatusCode)
        return
    }

    all , err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n",all)
    fmt.Printf("%s\n","	 商周时期以铜为金,因此这种文  ")

}
