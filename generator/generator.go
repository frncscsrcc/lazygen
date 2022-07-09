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
	Custom      bool
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
	Fields      []*Field
	Requires    []string
}

type Config struct {
	Module   string
	Folder   string
	Entities []EntityConfig
}

var entities map[string]bool

var funcMap = template.FuncMap{
	"fc":          FistCapital,
	"lc":          LowerCase,
	"typeConvert": TypeConvert,
	"getDTO":      GetDTO,
}

func init() {
	entities = make(map[string]bool)
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
		entities[strings.ToUpper(entity.Name)] = true
	}

	for _, entity := range config.Entities {
		folder := config.Folder + "/" + LowerCase(entity.Name)
		os.MkdirAll(folder, 0755)

		entity.Requires = make([]string, 0)
		for _, field := range entity.Fields {
			if _, validEntity := entities[strings.ToUpper(field.Type)]; validEntity {
				pkg := config.Module + "/" + config.Folder + strings.ToLower(field.Type)
				entity.Requires = append(entity.Requires, pkg)
				field.Custom = true
			}
		}

		generateTypes(folder+"/"+LowerCase(entity.Name)+".go", entity)
		generateMarshal(folder+"/"+LowerCase(entity.Name)+"_marshal.go", entity)
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

func GetDTO(inputType string) string {
	inputTypeUc := strings.ToUpper(inputType)
	if _, isValidEntity := entities[inputTypeUc]; isValidEntity {
		return "*" + LowerCase(inputType) + ".DTO"
	}
	return TypeConvert(inputType)
}

func TypeConvert(inputType string) string {
	inputTypeUc := strings.ToUpper(inputType)

	if _, isValidEntity := entities[inputTypeUc]; isValidEntity {
		return "*" + LowerCase(inputType) + "." + FistCapital(inputType)
	}

	switch inputTypeUc {
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

func generateMarshal(filename string, config EntityConfig) {
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

	tmpl, err := template.New("marshal.tpl").
		Funcs(funcMap).
		ParseFiles("./generator/templates/marshal.tpl")
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
