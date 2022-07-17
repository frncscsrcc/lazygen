package generator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
)

type Config struct {
	Module                string
	FolderDomain          string
	FolderAPI             string
	BasePathAPI           string
	FolderListener        string
	GenerateCRUDEndpoints bool
	Entities              []EntityConfig
}

type EntityConfig struct {
	Name          string
	Description   string
	Fields        []*Field
	Requires      []string
	RouterPackage string
	API           APIConfig
}

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

type APIConfig struct {
	Create APIResource
	Read   APIResource
	Update APIResource
	Delete APIResource
}

type APIResource struct {
	Enabled bool
	Roles   []string
}

var entities map[string]bool

var funcMap = template.FuncMap{
	"fc":          FistCapital,
	"lc":          LowerCase,
	"typeConvert": TypeConvert,
	"package":     CustomPackage,
	"getDTO":      GetDTO,
}

func init() {
	entities = make(map[string]bool)
}

func Generate(settingFile string) []string {
	content, err := ioutil.ReadFile(settingFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for _, entity := range config.Entities {
		entities[strings.ToUpper(entity.Name)] = true
	}

	generatedFiles := make([]string, 0)

	for _, entity := range config.Entities {
		folder := config.FolderDomain + "/" + LowerCase(entity.Name)
		os.MkdirAll(folder, 0755)

		entity.Requires = make([]string, 0)
		for _, field := range entity.Fields {
			if _, validEntity := entities[strings.ToUpper(field.Type)]; validEntity {
				pkg := config.Module + "/" + config.FolderDomain + "/" + strings.ToLower(field.Type)
				entity.Requires = append(entity.Requires, pkg)
				field.Custom = true
			}
		}
		typeSourceFile := folder + "/autogen_" + LowerCase(entity.Name) + ".go"
		executeTemplate(typeSourceFile, "type.tpl", entity, true)
		generatedFiles = append(generatedFiles, typeSourceFile)

		serializationlSourceFile := folder + "/autogen_" + LowerCase(entity.Name) + "_serialization.go"
		executeTemplate(serializationlSourceFile, "serialization.tpl", entity, true)
		generatedFiles = append(generatedFiles, serializationlSourceFile)

		customSourceFile := folder + "/" + LowerCase(entity.Name) + "_custom.go"
		executeTemplate(customSourceFile, "custom.tpl", entity, false)
		generatedFiles = append(generatedFiles, customSourceFile)
	}

	if config.GenerateCRUDEndpoints {
		os.MkdirAll(config.FolderAPI, 0755)
		routerAPIFile := config.FolderAPI + "/autogen_router.go"
		executeTemplate(routerAPIFile, "routerAPI.tpl", "", true)
		generatedFiles = append(generatedFiles, routerAPIFile)

		for _, entity := range config.Entities {
			entity.RouterPackage = config.Module + "/" + config.FolderAPI
			folder := config.FolderAPI + "/" + LowerCase(entity.Name)
			os.MkdirAll(folder, 0755)

			baseAPIFile := folder + "/autogen_" + LowerCase(entity.Name) + ".go"
			if entity.API.Create.Enabled ||
				entity.API.Read.Enabled ||
				entity.API.Update.Enabled ||
				entity.API.Delete.Enabled {
				executeTemplate(baseAPIFile, "baseRESTAPI.tpl", entity, true)
				generatedFiles = append(generatedFiles, baseAPIFile)
			} else {
				os.Remove(baseAPIFile)
			}
		}

		os.MkdirAll(config.FolderListener, 0755)
		listenerFile := config.FolderListener + "/autogen_listener.go"
		executeTemplate(listenerFile, "listener.tpl", config, true)
		generatedFiles = append(generatedFiles, listenerFile)

		customListenerFile := config.FolderListener + "/custom_listener.go"
		executeTemplate(customListenerFile, "customListener.tpl", "", false)
		generatedFiles = append(generatedFiles, customListenerFile)
	}

	return generatedFiles
}

func executeTemplate(filename string, templateName string, data interface{}, allowOverwrite bool) {
	if !allowOverwrite && checkFileExists(filename) {
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

	tmpl, err := template.New(templateName).Funcs(funcMap).ParseFiles("./generator/templates/" + templateName)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(file, data)
	log.Printf("- %s generated\n", filename)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func Show(interfaces ...interface{}) {
	fmt.Printf("%+v", interfaces...)
}
