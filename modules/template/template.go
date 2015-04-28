package template

import (
	"github.com/gorilla/context"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
	"reflect"
)

var funcMap = template.FuncMap{
	"noescape": func(s string) template.HTML {
		return template.HTML(s)
	},
}

var Render = render.New(render.Options{
	Layout:        "application",
	IsDevelopment: true,
	Funcs:         []template.FuncMap{funcMap},
})

func TemplateData(d interface{}, w http.ResponseWriter, r *http.Request) map[string]interface{} {

	current_user := context.Get(r, "current_user")

	output := map[string]interface{}{
		"current_user": current_user,
	}

	if d != nil {
		if reflect.TypeOf(d).String() == "map[string]interface {}" {
			for k, v := range d.(map[string]interface{}) {
				output[k] = v
			}
		} else {
			output["page"] = d
		}
	}

	return output
}
