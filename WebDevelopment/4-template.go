package main

import (
	"fmt"
	"html/template"
	"os"
)

// text/tempplate has same func as html but html is more secure
// Is more tuned for basic textual generation

func main() {
	templateString := `lemonade`
	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}
