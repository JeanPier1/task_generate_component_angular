package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"example.com/template_generate_components/internal/plantilla/componets_routes"
)

func ToTitleCase(s string) string {
	if s == "" {
		return s
	}
	temp := strings.ReplaceAll(s, "-", " ")
	c := cases.Title(language.English)
	titleCase := c.String(temp)
	result := strings.ReplaceAll(titleCase, " ", "")
	return result
}

var ComponentTemplateMap = map[string]map[string]string{
	"container": {
		".ts":   componets_routes.TemplateContenttWithDynamicContainerComponent,
		".html": componets_routes.TemplateContenttWithDynamicContainerHtml,
		".css":  componets_routes.TemplateContenttWithDynamicContainerCss,
	},
	"list": {
		".ts":   componets_routes.TemplateContenttWithDynamicListComponent,
		".html": componets_routes.TemplateContenttWithDynamicListHtml,
		".css":  componets_routes.TemplateContenttWithDynamicListCss,
	},
	"filter": {
		".ts":   componets_routes.TemplateContenttWithDynamicFilterComponent,
		".html": componets_routes.TemplateContenttWithDynamicFilterHtml,
		".css":  componets_routes.TemplateContenttWithDynamicFilterCss,
	},
}
