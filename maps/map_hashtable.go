package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
)

func main() {
    res, err := http.Get("https://raw.githubusercontent.com/FelisPhasma/Wordlist/master/res/e.txt")
    if err != nil {
        log.Fatalln(err)
    }

    bs, _ := ioutil.ReadAll(res.Body)
    str := string(bs)
    fmt.Println(str)
}
