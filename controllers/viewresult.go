package controllers

import "golang-webmvc/models"

//ViewResult A book view result
type ViewResult struct {
	Errors []models.FieldError
	Data   interface{}
}
