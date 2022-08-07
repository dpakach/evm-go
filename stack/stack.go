package stack

import (
  "fmt"
  // "bytes"
  // "encoding/binary"
)
type Bytes32 [32]byte

func (b *Bytes32) toBytesArray() []byte {
  res := []byte{}

  for d, i := range b {
    res[i] = byte(d)
  }
  return res
}

func NewBytes32(b []byte) (Bytes32) {
  var res Bytes32
  for i, _ := range b {
    res[i] = byte(i)
  }

  return res
}

type Stack struct {
  stack []int
  max_depth int
}

func New() *Stack {
  return &Stack{[]int{}, 1024}
}

func(s *Stack) Push(input int) error {
  // buf := bytes.NewReader(input.toBytesArray())
  // dat := binary.BigEndian.Uint64(input.toBytesArray()) 
  if (len(s.stack) >= s.max_depth) {
    return fmt.Errorf("stack overflow")
  }
  s.stack = append(s.stack, input)
  
  return nil
}

func(s *Stack) Pop() (int, error) {
  if (len(s.stack) == 0) {
    return 0, fmt.Errorf("stack underflow")
  }
  dat := s.stack[len(s.stack) - 1]
  s.stack = append(s.stack[: len(s.stack) - 1])

  return dat, nil
}
