package main

import (
  "grm/model"
  "log"
  "net/http"
  "os"
  "html/template"
)

func homeHandler(writer http.ResponseWriter, request *http.Request) {
  entries, err := os.ReadDir("/var/workspace/programs")
  if err != nil {
    log.Fatal(err)
  }
  template, err := template.ParseFiles("/html/home.html")
  if err != nil {
    log.Fatal(err)
  }
  var programs []string
  for _, entry := range entries {
    programs = append(programs, entry.Name())
  }
  template.Execute(writer, map[string]interface{}{
    "Programs": programs,
  })
}

func programHandler(writer http.ResponseWriter, request *http.Request) {
  programName := request.PathValue("program")
  program := model.LoadProgram("/var/workspace/programs/" + programName)
  template, err := template.ParseFiles("/html/program.html")
  if err != nil {
    log.Fatal(err)
  }
  template.Execute(writer, program)
}

func programExerciseHandler(writer http.ResponseWriter, request *http.Request) {
  log.Print(request.PathValue("program") + "/" + request.PathValue("exercise"))
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/program/{program}", programHandler)
  http.HandleFunc("/program/{program}/exercise/{exercise}",
    programExerciseHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
