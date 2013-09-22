package cod2

import (
  "strings"
  "github.com/adabei/goldenbot/events/cod"
)

func Parse(line string) interface{} {
  
  offset := 0
  if line[:1] == " " {
    offset = strings.Index(line[1:], " ") + 2
  } else {
    offset = strings.Index(line, " ") + 1
  }
  
  values := strings.Split(line[offset:], ";")

  switch values[0] {
    default:
      return cod.Parse(line)
  }
  return nil
}
