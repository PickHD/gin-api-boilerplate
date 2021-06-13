package docs

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        os.Getenv("PUBLIC_URL"),
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Gin API Boilerplate",
	Description: "This is example docs api ",
}

type s struct{}

func (s *s) ReadDoc() string {

	//*Read from api.json files
	doc, err := os.ReadFile("./src/docs/api.json")
	if err != nil {
		log.Fatal(err)
	}

	//*Convert into a string
	docStr := string(doc)

	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(docStr)
	if err != nil {
		return docStr
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return docStr
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
