package memory

import (
  "fmt"
)

const UINT256_MAX = 999999999;

type Memory struct {
  memory map[int]byte
}

func New() *Memory {
  return &Memory{make(map[int]byte)}
}


func(m *Memory) Store(offset int, value byte) error {
  if (offset < 0 || offset > UINT256_MAX) {
    return fmt.Errorf("invalid memory offset")
  }
  m.memory[offset] = value
  return nil
} 


func(m *Memory) Load(offset int) (error, byte) {
  if (offset < 0 || offset >= len(m.memory)) {
    return fmt.Errorf("invalid memory offset"), byte(0)
  }
  return nil, m.memory[offset]
}

func(m *Memory) LoadRange(offset, length int) (error, []byte) {
  if (offset < 0 || offset >= len(m.memory)) {
    return fmt.Errorf("invalid memory offset"), []byte{}
  }

  ret := []byte{}

  for i:= offset; i < offset+length; i++ {
    ret = append(ret, m.memory[i])
  }

  return nil, ret
}

