package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

const url1 = "http://www.zhenai.com/zhenghun"

const url2 = "http://album.zhenai.com/u/1445775408"

func main() {
    fmt.Printf("hello, world\n")

    // :=
    resp,err := http.Get(url2)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()


    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: status code",resp.StatusCode)
        //return
    }

    all , err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n",all)
    fmt.Printf("%s\n","	 商周时期以铜为金,因此这种文  ")









    client := http.Client{}
    request, err := http.NewRequest("GET", url1, nil)
    if err!=nil {
        fmt.Printf("wrong http request: %s", err.Error())
        //return nil, fmt.Errorf("wrong http request: %s", err.Error())
    }
    request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
    resp2, err := client.Do(request)
    if err != nil {
        //return nil, err
        fmt.Println("Error status code 000000000")

    }
    defer resp2.Body.Close()
    if resp2.StatusCode != http.StatusOK {
        fmt.Println("Error status code:", resp2.StatusCode)
        //return nil, fmt.Errorf("wrong http status code: %d", resp.StatusCode)
    }

    // 编码转换
    //htmlEncoding:=determineEncoding(resp.Body)
    //reader := transform.NewReader(resp.Body, htmlEncoding.NewDecoder())
    //body, err := ioutil.ReadAll(reader)

    all2 , err2 :=  ioutil.ReadAll(resp2.Body)


    if err2 != nil {
        panic(err2)
    }

    fmt.Printf("222222222222 %s\n",all2)
    fmt.Printf("%s\n","	 商周时期以铜为金,因此这种文  ")



}
