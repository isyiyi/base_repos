package content

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateHtml() {
	var movie = []Movie{0: {Title: "战狼",
		Year:   2018,
		Color:  false,
		Actors: []string{"吴京", "张翰"},
	}, 1: {
		Title:  "大话西游",
		Year:   2022,
		Color:  false,
		Actors: []string{"章子怡", "吴青峰"},
	},
	}
	const tpl string = `<html>
	{{range .}}
		<h1>{{.Title}}</h1>
		<h4>{{.Year}}</h4>
		{{range .Actors}}
			<span>{{.}}</span>
		{{end}}
	{{end}}
</html>`

	p, err := template.New("htmlTest").Parse(tpl)
	if err != nil {
		fmt.Println(err)
	}
	err = p.Execute(os.Stdout, movie)
	if err != nil {
		fmt.Println(err)
	}
}