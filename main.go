package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"

	"example.com/template_generate_components/plantilla"
)

func getComponentAndDir() (nameComponent, rootDir string, err error) {
	fmt.Print("Introducir el nombre del componente al crear: ")
	fmt.Scanln(&nameComponent)

	fmt.Print("Introducir el directorio raíz: ")
	fmt.Scanln(&rootDir)

	err = nil

	nameCompTrimmed := strings.TrimSpace(nameComponent)
	rootDirTrimmed := strings.TrimSpace(rootDir)

	nameComponent = nameCompTrimmed
	rootDir = rootDirTrimmed

	if len(nameCompTrimmed) == 0 || len(rootDirTrimmed) == 0 {
		fmt.Println("❌ ERROR DE VALIDACIÓN:")
		fmt.Println("Ambos valores son obligatorios y no pueden estar vacíos o contener solo espacios.")
		return
	}
	return
}

func postCreateDirAndComponent(nameComponent, rootDir string, err error) (fullPathCreate string, errCreate error) {
	fullPath := rootDir + string(os.PathSeparator) + nameComponent

	err = os.MkdirAll(fullPath, 0755)

	if err != nil {
		fmt.Printf("❌ ERROR al crear la carpeta en la ruta '%s': %v\n", fullPath, err)
		return
	}
	return fullPath, err
}

type TemplateData struct {
	ComponentName string
}

func createComponentRoutes(nameComponent string, fullPath string, err error) {
	fileName := fullPath + "/" + nameComponent + ".routes.ts"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creando el archivo:", err)
	}
	defer file.Close()

	funcMap := template.FuncMap{
		"title": ToTitleCase,
	}

	templateString := plantilla.TemplateContentWithDynamicRoutes

	t, err := template.New("routes").Funcs(funcMap).Parse(templateString)
	if err != nil {
		fmt.Println("Error al parsear la plantilla:", err)
		return
	}

	structData := TemplateData{
		ComponentName: nameComponent,
	}
	err = t.Execute(file, structData)
	if err != nil {
		fmt.Println("Error al ejecutar la plantilla:", err)
	}
}

func ToTitleCase(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

var componentTemplateMap = map[string]map[string]string{
	"container": {
		".ts":   plantilla.TemplateContenttWithDynamicContainerComponent,
		".html": plantilla.TemplateContenttWithDynamicContainerHtml,
		".css":  plantilla.TemplateContenttWithDynamicContainerCss,
	},
	"list": {
		".ts":   plantilla.TemplateContenttWithDynamicListComponent,
		".html": plantilla.TemplateContenttWithDynamicListHtml,
		".css":  plantilla.TemplateContenttWithDynamicListCss,
	},
	"filter": {
		".ts":   plantilla.TemplateContenttWithDynamicFilterComponent,
		".html": plantilla.TemplateContenttWithDynamicFilterHtml,
		".css":  plantilla.TemplateContenttWithDynamicFilterCss,
	},
}

func createComponentInsert(nameComponent string, fullPath string, TypeComponent string, fileExtension [3]string) {
	extensionTemplates, exists := componentTemplateMap[TypeComponent]

	if !exists {
		log.Printf("Advertencia: Tipo de componente '%s' desconocido. No se puede crear el archivo.", TypeComponent)
		return
	}

	for i, ext := range fileExtension {
		fmt.Printf("Index %d: %s\n", i, ext)
		templateString, extExists := extensionTemplates[ext]

		if !extExists {
			log.Printf("Advertencia: No se encontró plantilla para la extensión '%s' en el componente '%s'. Saltando archivo.", ext, TypeComponent)
			continue
		}

		fileName := fullPath + "/" + nameComponent + "-" + TypeComponent + ext

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Error creando el archivo:", err)
		}
		defer file.Close()

		funcMap := template.FuncMap{
			"title": ToTitleCase,
		}

		t, err := template.New(ext).Funcs(funcMap).Parse(templateString)
		if err != nil {
			fmt.Println("Error al parsear la plantilla para", ext, ":", err)
			return
		}

		structData := TemplateData{
			ComponentName: nameComponent,
		}

		err = t.Execute(file, structData)
		if err != nil {
			fmt.Println("Error al ejecutar la plantilla para", ext, ":", err)
		}
	}
}

func main() {
	// Verificar si el directorio de trabajo es el directorio actual
	nameComponent, rootDir, err := getComponentAndDir()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Crear la carpeta y el componente
	fullPathCreate, err := postCreateDirAndComponent(nameComponent, rootDir, err)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	//  create component routes
	createComponentRoutes(nameComponent, fullPathCreate, err)

	// create components files components
	nameContent := "components"
	fullPathCreateComponent, err := postCreateDirAndComponent(nameContent, fullPathCreate, err)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	extensiones := [3]string{".ts", ".html", ".css"}

	// content component
	nameCotainer := nameComponent + "-container"
	fullPathCreateComponentContainer, err := postCreateDirAndComponent(nameCotainer, fullPathCreateComponent, err)
	createComponentInsert(nameComponent, fullPathCreateComponentContainer, "container", extensiones)

	// filter component
	nameFilter := nameComponent + "-filter"
	fullPathCreateComponentFilter, err := postCreateDirAndComponent(nameFilter, fullPathCreateComponent, err)
	createComponentInsert(nameComponent, fullPathCreateComponentFilter, "filter", extensiones)

	// list component
	nameList := nameComponent + "-list"
	fullPathCreateComponentList, err := postCreateDirAndComponent(nameList, fullPathCreateComponent, err)
	createComponentInsert(nameComponent, fullPathCreateComponentList, "list", extensiones)
}
