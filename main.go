package main

import (
	"github.com/k8guard/k8guard-report/templates"

	"github.com/k8guard/k8guard-report/db"
)

func main() {
	templates.PopulateTemplates()
	db.InitDB()
	start_http_router()
}
