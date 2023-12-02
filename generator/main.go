package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func main() {
	content, err := ioutil.ReadFile("entity-schema.json")

	if err != nil {
		log.Fatalln(err)
	}

	var schema map[string]Entity

	if err := json.Unmarshal(content, &schema); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("")

	names := make([]string, 0)

	for _, entity := range schema {
		t := template.Must(template.New("entity_repo.tpl").ParseFiles("entity_repo.tpl"))

		file, _ := os.Create(fmt.Sprintf("../repo_%s.go", entity.Name))

		names = append(names, strcase.ToCamel(entity.Name))

		info := EntityRepositoryTplInfo{
			Name:          entity.Name,
			FormattedName: strcase.ToCamel(entity.Name),
			ApiPath:       strings.ReplaceAll(entity.Name, "_", "-"),
			Fields:        []TplField{},
			HasTimeField:  entity.HasTimeField(),
		}

		for name, property := range entity.Properties {
			info.Fields = append(info.Fields, TplField{
				Key:  strcase.ToCamel(name),
				Name: name,
				Type: property.GetType(),
			})
		}

		t.Execute(file, info)
		file.Close()
	}

	fmt.Println("Generating generic repo.go")

	file, _ := os.Create("../repo.go")
	t := template.Must(template.New("repo.tpl").ParseFiles("repo.tpl"))

	t.Execute(file, names)

	file.Close()
	//fmt.Println(schema)
}

type EntityRepositoryTplInfo struct {
	Name          string
	FormattedName string
	ApiPath       string
	Fields        []TplField
	HasTimeField  bool
}

type TplField struct {
	Name string
	Key  string
	Type string
}

type Entity struct {
	Name       string                    `json:"entity"`
	Properties map[string]EntityProperty `json:"properties"`
}

func (e Entity) HasTimeField() bool {
	for _, p := range e.Properties {
		if p.GetType() == "time.Time" {
			return true
		}
	}

	return false
}

type EntityProperty struct {
	Type     string `json:"type"`
	Relation string `json:"relation"`
	Entity   string `json:"entity"`
}

func (p EntityProperty) GetType() string {
	switch p.Type {
	case "uuid":
		return "string"
	case "string":
		return "string"
	case "text":
		return "string"
	case "date":
		return "time.Time"
	case "association":
		formatted := strcase.ToCamel(p.Entity)

		if strings.HasSuffix(p.Relation, "_to_many") {
			return "[]" + formatted
		}

		return "*" + formatted
	case "json_object":
		return "interface{}"
	case "json_list":
		return "interface{}"
	case "boolean":
		return "bool"
	case "float":
		return "float64"
	case "int":
		return "float64"
	}

	return "interface{}"
}
