package models

import "github.com/RazvanING/bookings/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//Cross Site Request Forgery
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
