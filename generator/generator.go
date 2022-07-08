package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Field struct {
	Name        string
	Type        string
	Multiple    bool
	Description string
	PrimaryKey  bool
	Serchable   bool
	Exposed     bool
	Private     bool
}

type EntityConfig struct {
	Name        string
	Description string
	Fields      []Field
}

type Config struct {
	Module   string
	Folder   string
	Entities []EntityConfig
}

var funcMap = template.FuncMap{
	"fc":          FistCapital,
	"lc":          LowerCase,
	"typeConvert": TypeConvert,
}

func Generate(settingFile string) {
	content, err := ioutil.ReadFile(settingFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	//outputFolder := config.Folder

	for _, entity := range config.Entities {
		folder := config.Folder + "/" + LowerCase(entity.Name)
		os.MkdirAll(folder, 0755)

		generateTypes(folder+"/"+LowerCase(entity.Name)+".go", entity)
		generateCustom(folder+"/"+LowerCase(entity.Name)+"_custom.go", entity)
	}
}

func FistCapital(input string) string {
	output := ""
	if len(input) == 0 {
		return ""
	}
	output += strings.ToUpper(string(input[0]))
	output += strings.ToLower(input[1:])
	return output
}

func LowerCase(input string) string {
	return strings.ToLower(input)
}

func TypeConvert(inputType string) string {
	switch strings.ToUpper(inputType) {
	case "INT":
		return "int"
	case "BIGINT":
		return "int64"
	case "STRING":
		return "string"
	default:
		panic("type " + inputType + " can not be converted in golang datatye")
	}
}

func generateTypes(filename string, config EntityConfig) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	tmpl, err := template.New("type.tpl").
		Funcs(funcMap).
		ParseFiles("./generator/templates/type.tpl")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(file, config)
	log.Printf("- %s generated\n", filename)
}

func generateCustom(filename string, config EntityConfig) {
	if checkFileExists(filename) {
		log.Printf("- %s exists, skipping\n", filename)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	tmpl, err := template.New("custom.tpl").
		Funcs(funcMap).
		ParseFiles("./generator/templates/custom.tpl")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(file, config)
	log.Printf("- %s generated\n", filename)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func Show(interfaces ...interface{}) {
	fmt.Printf("%+v", interfaces...)
}
