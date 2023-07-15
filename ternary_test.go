package ternary

import (
  "testing"
  "fmt"
)

func TestTernary(t *testing.T) {
  Ternary(true, func() {
    fmt.Println("Yipee!")
  }, func() {
    t.Fatalf("Literally how?  Is true false in your universe?")
  })
  Ternary(false, func() {
    t.Fatalf("!?!?!?!")
  }, func() {
    fmt.Println("Yipee!")
  })
}
