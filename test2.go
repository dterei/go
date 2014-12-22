package main

import (
  "fmt"
  "runtime"
  "time"
)

func deferGC(gcnum uint32, alloc, pause uint64) bool {
  fmt.Printf("gc manager: %d\n", gcnum)

  time.AfterFunc(100 * time.Millisecond, func() {
    runtime.GCStart()
  })

	return true
}

func main() {
  runtime.RegisterGCCallback(deferGC)
  runtime.GC()
  time.Sleep(500 * time.Millisecond)
  runtime.GC()
  time.Sleep(500 * time.Millisecond)
  runtime.GC()
  time.Sleep(500 * time.Millisecond)
}

