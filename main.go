package ternary

func Ternary(condition bool, ifTrue func(), ifFalse func()) {
  if condition {
    ifTrue()
  } else {
    ifFalse()
  }
}
