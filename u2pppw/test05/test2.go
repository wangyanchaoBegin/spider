package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
)

const text = "my email is 1234@qq.com  1234@qq.com.cn  asdf@zte.com  iuyt@aliyun.com  1234@@@qq.com "
 
func main() {


    fmt.Printf("test3 \n")
    //re , err := regexp.Compile("1234@qq.com")

    //re := regexp.MustCompile("1234@qq.com")



    //re := regexp.MustCompile(".+@qq.com")     1 or many


    //re := regexp.MustCompile(".*@qq.com") // 0 or many

    pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
    re := regexp.MustCompile(pattern) // 0 or many

    //match := re.FindString(text)         -1 mean all data   i  want
    match := re.FindAllString(text,-1)

    fmt.Println(match)
    fmt.Printf("match:%s \n",match)


    match2 := re.FindAllStringSubmatch(text,-1)
    fmt.Println(match2)

    fmt.Printf("re:%s \n",re)

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
        fmt.Printf("%s\n",all)
    }

    printCityList(all)


}


func printCityList(contents []byte){

    pattern := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

    re := regexp.MustCompile(pattern)


    matches := re.FindAllSubmatch(contents , -1)

    for _, m := range matches {

        //fmt.Println("city : %s,URL : %s\n",m[2],m[1])  number ] ...


        fmt.Printf("city : %s,URL : %s\n",m[2],m[1])

        //for _, subMatch := range m {
		//
        //    fmt.Println("%s",subMatch)
        //}

        fmt.Println()
    }

    fmt.Println("matches size : %s",len(matches))
}


//
//
//import (
//"fmt"
//"net/http"
//"io/ioutil"
//"regexp"
//)
//
//func main() {
//
//    fmt.Printf("test2 \n")
//
//    // :=
//    resp,err := http.Get("http://www.zhenai.com/zhenghun")
//    if err != nil {
//        panic(err)
//    }
//
//    defer resp.Body.Close()
//
//
//    if resp.StatusCode != http.StatusOK {
//        fmt.Println("Error: status code",resp.StatusCode)
//        return
//    }
//
//    all , err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        panic(err)
//    }
//
//    fmt.Printf("%s\n",all)
//}

