package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"strconv"
)

//import (
//	"net/http"
//	"fmt"
//	"io/ioutil"
//)
// !!!!!     wei le xianzhi pachong su du
var rateLimiter = time.Tick(50 * time.Millisecond)
var agentname int = 1

//var rateLimiter = time.Tick(980 * time.Millisecond)

func Fetch(url string) ([]byte ,error){
	//
	//// :=
	//resp,err := http.Get(url)
	//if err != nil {
	//	return nil,err
	//}
	//
	//defer resp.Body.Close()
	//
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code",resp.StatusCode)
	//	//return nil ,errors.New("")
	//	return nil , fmt.Errorf("wrong status code : %d", resp.StatusCode)
	//}
	//
	//return ioutil.ReadAll(resp.Body)

// wei le xianzhi pachong su du
//	<-rateLimiter




	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		fmt.Printf("wrong http request: %s", err.Error())
		return nil, fmt.Errorf("wrong http request: %s", err.Error())
	}



	agentname++
	var agentnameTmp = strconv.Itoa(agentname)
	//request.Header.Set("User-Agent",
	//	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")


	request.Header.Set("User-Agent", agentnameTmp)



	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code:", resp.StatusCode)
		return nil, fmt.Errorf("wrong http status code: %d", resp.StatusCode)
	}

	// 编码转换
	//htmlEncoding:=determineEncoding(resp.Body)
	//reader := transform.NewReader(resp.Body, htmlEncoding.NewDecoder())
	//body, err := ioutil.ReadAll(reader)

	return ioutil.ReadAll(resp.Body)

}
