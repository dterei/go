package main

import (
  "fmt"
  "runtime"
  "time"
)

var m runtime.MemStats

func gcManager(alloc, last_pause int64, ret *int64) {
  n := make([]byte, 10000000)
  n[0] = 'c'
  _ = n
  runtime.ReadMemStats(&m)
  
  fmt.Printf("GC Baby! (alloc: %d, pause: %d)\n", alloc, last_pause)
  fmt.Printf("GC Check! (alloc: %d, pause: %d)\n",
    m.Alloc, m.PauseNs[(m.NumGC+255)%256])

  *ret = 1

  go func() {
    time.Sleep(2 * time.Second)
    fmt.Printf("Starting GC...\n")
    runtime.GCStart()
  }()
}

func main() {
  go func() {
    for i := 0; i < 10; i++ {
      fmt.Printf("Goroutine - %d\n", i)
      time.Sleep(700 * time.Millisecond);
      n := make([]byte, 1000)
      _ = n
    }
  }()

  fmt.Printf("x0\n")
  runtime.RegisterGCCallback(gcManager)
  runtime.GC()
  fmt.Printf("x1\n")
  time.Sleep(700 * time.Millisecond);
  fmt.Printf("x2\n")
  time.Sleep(700 * time.Millisecond);
  fmt.Printf("x3\n")
  time.Sleep(700 * time.Millisecond);
  fmt.Printf("x4\n")
  time.Sleep(700 * time.Millisecond);
  fmt.Printf("x5\n")
}

