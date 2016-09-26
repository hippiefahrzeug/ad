package main

import (
  "fmt"
  "time"
  "flag"
  "bufio"
  "os"
)

func delim(t *time.Timer, delimiter string, num int) {
    <-t.C
    for i := 0; i < num; i++ {
      fmt.Println(delimiter)
    }
}

func main() {
  delimiter := flag.String("d", "#########################################", "the delimiter to be printed out")
  duration := flag.Int("t", 3, "number of seconds to wait before printing the delimiter")
  num := flag.Int("n", 3, "number of times to print delimiter")
  flag.Parse()
  scanner := bufio.NewScanner(os.Stdin)
  timer := time.NewTimer(time.Second)

  for scanner.Scan() {
    timer.Stop()
    fmt.Println(scanner.Text())
    timer = time.NewTimer((time.Duration)(*duration)*time.Second)
    go delim(timer, *delimiter, *num)
  }
}
