package main

import (
  "fmt"
  "grm/model"
)

func main() {
  program := model.LoadProgram("Ultimate PPL")
  fmt.Println(program)
  program.Save()
}
