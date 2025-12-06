package app

import (
	"fmt"

	"example.com/template_generate_components/internal/config"
)

func Run() {
	// Verificar si el directorio de trabajo es el directorio actual
	nameComponent, rootDir, err := config.GetComponentAndDir()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Crear la carpeta y el componente
	fullPathCreate, err := config.PostCreateDirAndComponent(nameComponent, rootDir, err)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Crear los archivos de rutas
	config.CreateComponentRoutes(nameComponent, fullPathCreate, err)

	// create components files components
	nameContent := "components"
	fullPathCreateComponent, err := config.PostCreateDirAndComponent(nameContent, fullPathCreate, err)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	extensiones := [3]string{".ts", ".html", ".css"}

	// content component
	nameCotainer := nameComponent + "-container"
	fullPathCreateComponentContainer, err := config.PostCreateDirAndComponent(nameCotainer, fullPathCreateComponent, err)
	config.CreateComponentInsert(nameComponent, fullPathCreateComponentContainer, "container", extensiones)

	// filter component
	nameFilter := nameComponent + "-filter"
	fullPathCreateComponentFilter, err := config.PostCreateDirAndComponent(nameFilter, fullPathCreateComponent, err)
	config.CreateComponentInsert(nameComponent, fullPathCreateComponentFilter, "filter", extensiones)

	// list component
	nameList := nameComponent + "-list"
	fullPathCreateComponentList, err := config.PostCreateDirAndComponent(nameList, fullPathCreateComponent, err)
	config.CreateComponentInsert(nameComponent, fullPathCreateComponentList, "list", extensiones)

}
