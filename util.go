package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func httpRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response" + string(body))
	return body, nil
}


func field(t interface{}, key string) reflect.Value {
	strs := strings.Split(key, ".")
	v := reflect.ValueOf(t)
	for _, s := range strs[1:] {
		v = v.FieldByName(s)
	}
	return v
}

