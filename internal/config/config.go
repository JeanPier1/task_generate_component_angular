package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"example.com/template_generate_components/internal/models"
	"example.com/template_generate_components/internal/plantilla/componets_routes"
	"example.com/template_generate_components/internal/utils"
)

func GetComponentAndDir() (nameComponent, rootDir string, err error) {
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

func PostCreateDirAndComponent(nameComponent, rootDir string, err error) (fullPathCreate string, errCreate error) {
	fullPath := rootDir + string(os.PathSeparator) + nameComponent

	err = os.MkdirAll(fullPath, 0755)

	if err != nil {
		fmt.Printf("❌ ERROR al crear la carpeta en la ruta '%s': %v\n", fullPath, err)
		return
	}
	return fullPath, err
}

func CreateComponentRoutes(nameComponent string, fullPath string, err error) {
	fileName := fullPath + "/" + nameComponent + ".routes.ts"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creando el archivo:", err)

	}
	defer file.Close()

	funcMap := template.FuncMap{
		"title": utils.ToTitleCase,
	}

	templateString := componets_routes.TemplateContentWithDynamicRoutes

	t, err := template.New("routes").Funcs(funcMap).Parse(templateString)
	if err != nil {
		fmt.Println("Error al parsear la plantilla:", err)
		return
	}

	structData := models.TemplateData{
		ComponentName: nameComponent,
	}
	err = t.Execute(file, structData)
	if err != nil {
		fmt.Println("Error al ejecutar la plantilla:", err)
	}
}
func CreateComponentInsert(nameComponent string, fullPath string, TypeComponent string, fileExtension [3]string) {
	extensionTemplates, exists := utils.ComponentTemplateMap[TypeComponent]

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
			"title": utils.ToTitleCase,
		}

		t, err := template.New(ext).Funcs(funcMap).Parse(templateString)
		if err != nil {
			fmt.Println("Error al parsear la plantilla para", ext, ":", err)
			return
		}

		structData := models.TemplateData{
			ComponentName: nameComponent,
		}

		err = t.Execute(file, structData)
		if err != nil {
			fmt.Println("Error al ejecutar la plantilla para", ext, ":", err)
		}
	}
}
