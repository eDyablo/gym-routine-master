package model

type Exercise struct {
  AlternativeNames  []string  `yaml:"alternativeNames,omitempty"`
  ConventionalName  string    `yaml:"conventionalName"`
  Links             []Link    `yaml:"links,omitempty"`
}
