package main

import (
  "fmt"
  "github.com/dpakach/evm-go/execution"
  "encoding/hex"
)

// 600660070260005360016000f3

// # INSTRUCTIONS
// PUSH1 06
// PUSH1 07
// MUL
// PUSH1 0
// MSTORE8
// PUSH1 1
// PUSH1 0
// RETURN


func main() {
  s := "600660070260005360016000f3"
  dat, err := hex.DecodeString(s)
  if err != nil {
    panic(err)
  }

  execution.RegisterBasicInstructions()

  intCode := []int{}
  for _, v := range dat {
    intCode = append(intCode, int(v))
  }
  fmt.Println(intCode)
  err = execution.Run(intCode)

  if err != nil {
    panic(err)
  }
}
