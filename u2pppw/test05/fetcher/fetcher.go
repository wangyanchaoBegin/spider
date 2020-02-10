package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func Fetch(url string) ([]byte ,error){

	// :=
	resp,err := http.Get(url)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		//return nil ,errors.New("")
		return nil , fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
