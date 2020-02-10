package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

const url = "http://album.zhenai.com/u/1445775408"

func main() {

	// :=
	resp,err := http.Get(url)
	if err != nil {
		//return nil,err
	}

	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		//return nil , fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	//return ioutil.ReadAll(resp.Body)

	context, err := ioutil.ReadAll(resp.Body)


}
