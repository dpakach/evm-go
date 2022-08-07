package execution

import (
  "fmt"
  "time"
  "github.com/dpakach/evm-go/stack"
  "github.com/dpakach/evm-go/memory"
)


const UINT256_MAX = 999999999; // TODO: implement 2 ^ 256 here

type ExecutionContext struct {
  code []int
  stack *stack.Stack
  memory *memory.Memory
  pc int
  stopped bool
  returnData []byte
}

func New(code []int) *ExecutionContext {
  return &ExecutionContext {
    code, stack.New(), memory.New(), 0, false, []byte{},
  }
}

func (e *ExecutionContext) stop() {
  e.stopped = true;
}

func (e *ExecutionContext) readCode(numBytes int) []int {
  val := e.code[e.pc:e.pc+numBytes]
  e.pc += numBytes

  return val
}

func (e *ExecutionContext) setReturnData(offset int, length int) {
  e.stopped = true
  err, dat := e.memory.LoadRange(offset, length)
  if err != nil {
    handleError(err)
  }

  e.returnData = dat
}


var INSTRUCTIONS []Instruction;
var INSTRUCTIONS_BY_OPCODE map[int]Instruction = make(map[int]Instruction);

type Instruction struct {
  opcode int
  name string
  execute func(ctx *ExecutionContext)
}

func registerInstruction(
  opcode int, name string, executeFunction func(ctx *ExecutionContext),
) (error, *Instruction) {
  instruction := Instruction{opcode, name, executeFunction}
  INSTRUCTIONS = append(INSTRUCTIONS, instruction)
  if _, ok := INSTRUCTIONS_BY_OPCODE[opcode]; ok {
    return fmt.Errorf("Opcode already registered"), nil
  }
  INSTRUCTIONS_BY_OPCODE[opcode] = instruction

  return nil, &instruction
}


func handleError(err error) {
  if err != nil {
    panic(err)
  }
}

func RegisterBasicInstructions() {
  err, _ := registerInstruction(0x00, "STOP", func(ctx *ExecutionContext) {
    ctx.stop()
  })

  handleError(err)

  err, _ = registerInstruction(0x60, "PUSH1", func(ctx *ExecutionContext) {
    ctx.stack.Push(ctx.readCode(1)[0])
  })
  handleError(err)

  err, _ = registerInstruction(0x01, "ADD", func(ctx *ExecutionContext) {
    val1, err := ctx.stack.Pop()
    handleError(err)

    val2, err := ctx.stack.Pop()
    handleError(err)

    ctx.stack.Push((val1 + val2) % UINT256_MAX)
  })
  handleError(err)

  err, _ = registerInstruction(0x02, "MUL", func(ctx *ExecutionContext) {
    val1, err := ctx.stack.Pop()
    handleError(err)

    val2, err := ctx.stack.Pop()
    handleError(err)

    ctx.stack.Push((val1 * val2) % UINT256_MAX)
  })
  handleError(err)

  err, _ = registerInstruction(0x53, "MSTORE8", func(ctx *ExecutionContext) {
    val1, err := ctx.stack.Pop()
    handleError(err)

    val2, err := ctx.stack.Pop()
    handleError(err)

    ctx.memory.Store(val1, byte(val2%256))
  })


  err, _ = registerInstruction(0xf3, "RETURN", func(ctx *ExecutionContext) {
    val1, err := ctx.stack.Pop()
    handleError(err)

    val2, err := ctx.stack.Pop()
    handleError(err)

    ctx.setReturnData(val1, val2)
  })
  handleError(err)
}


func decodeOpcode(ctx *ExecutionContext) (*Instruction, error) {
  if ctx.pc < 0 || ctx.pc >= len(ctx.code) {
    return nil, fmt.Errorf("Invalid offset") // TODO: PROPER ERROR
  }

  opcode := ctx.readCode(1)[0]

  if _, ok := INSTRUCTIONS_BY_OPCODE[opcode]; !ok {
    return nil, fmt.Errorf("Unknown opcode")
  }
  instruction := INSTRUCTIONS_BY_OPCODE[opcode]
  return &instruction, nil
}

func Run(code []int) error {
  ctx := New(code)

  for !ctx.stopped {

    // Uncomment these for debug
    // lastPc := ctx.pc

    instruction, err := decodeOpcode(ctx)
    handleError(err)
    instruction.execute(ctx)


    // fmt.Printf("%v @ pc={%v}", instruction, lastPc)
    // fmt.Println("stack", ctx.stack)
    // fmt.Println(ctx)
    // fmt.Println()
    // time.Sleep(time.Second * 1)

    fmt.Println(ctx.returnData)
    time.Sleep(time.Millisecond * 1)
  }

  return nil
}
