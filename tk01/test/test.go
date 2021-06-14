package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i <= 10; i++ {
		go (func() {
			wg.Add(1)
			defer wg.Done()
			req, err := http.NewRequest("GET", "http://services.shenyi.me/v1/users", nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			res, err := http.DefaultClient.Do(req)

			if err != nil {
				fmt.Println("出错1", err)
				return
			}
			defer res.Body.Close()
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("出错2", err)
				return
			}
			fmt.Println(res.Status, ":", string(resBody))
		})()
	}
	wg.Wait()
}
