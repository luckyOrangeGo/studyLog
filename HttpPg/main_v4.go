package main

//模板引擎
import (
	"html/template"
	"log"
	"os"
)

type Package struct {
	Name     string
	NumVars  int
	NumFuncs string
}

func main() {
	tmpl, err := template.New("go-web").Parse(`
		Hello go-web, Package name: {{.Name}}
		Number of functions: {{.NumFuncs}}
		Number of variables: {{.NumVars}}`)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	tmpl.Execute(os.Stdout, "gogogo")
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}

	err = tmpl.Execute(os.Stdout, &Package{
		Name:     "go-web",
		NumFuncs: "12",
		NumVars:  1200,
	})
}
