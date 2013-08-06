// Package advert provides a plugin to display messages in a set interval
// to all players.
package advert

import (
  "bufio"
  "container/list"
  "os"
  "time"
  "github.com/adabei/goldenbot/rcon"
)

type Advert struct {
  input string
  Interval int
  requests chan rcon.RCONRequest
}

func NewAdvert(input string, interval int, requests chan rcon.RCONRequest) *Advert {
  a := new(Advert)
  a.requests = requests
  a.Interval = interval
  a.input = input 
  return a
}

func (a *Advert) Start(next, prev chan string){
  go a.Advertise()
  for {
    in := <-prev
    next <- in
  }
}

func (a *Advert) Advertise() {
  l := list.New()
  fi, _ := os.Open(a.input)
  scanner := bufio.NewScanner(fi)
  for scanner.Scan() {
    if val := scanner.Text(); len(val) == 0 {
      l.PushBack(nil)
    } else {
    
      l.PushBack(val)
    }
  }

  fi.Close()

  for {
    for e := l.Front(); e != nil; e = e.Next() {
      if e.Value != nil {
        if val, ok := e.Value.(string); ok {
          a.requests <- *rcon.NewRCONRequest("say \"" + "^3^golden^7bot: " + val + "\"", nil)
          time.Sleep(1000 * time.Millisecond)
        }
      } else {
        time.Sleep(time.Duration(a.Interval) * time.Millisecond)
      }
    }
  }
}
