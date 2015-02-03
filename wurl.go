package main

import (
  "os"
  "io/ioutil"
  "net/http"
  "strings"
  "github.com/codegangsta/cli"
)

func main() {
  wurl := cli.NewApp()
  wurl.Name = "Wurl - a friendly curl like cli with support for websockets."
  wurl.Usage = "wurl [options] [...url]"
  wurl.Author = "rongruesbeck@gmail.com"
  wurl.Action = func(c *cli.Context) {
    u := string(c.Args()[0])
    if !strings.Contains(u, "http") {
      u = "http://" + u
    }
    client := &http.Client{}
    resp, err := client.Get(string(u))
    if err != nil {
      println(resp)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    os.Stdout.Write(body)
  }
  wurl.Run(os.Args)
}
