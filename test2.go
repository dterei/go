package main

import (
  "fmt"
  "runtime"
  "time"
)

var gc int = 0

func gcManager(alloc, last_pause int64, ret *int64) {
  gc++
  fmt.Printf("gc manager: %d\n", gc)
  *ret = 1

  time.AfterFunc(100 * time.Millisecond, func() {
    runtime.GCStart()
  })
}

func main() {
  runtime.RegisterGCCallback(gcManager)
  runtime.GC()
  time.Sleep(500 * time.Millisecond)
  runtime.GC()
  time.Sleep(500 * time.Millisecond)
  runtime.GC()
  time.Sleep(500 * time.Millisecond)
}

