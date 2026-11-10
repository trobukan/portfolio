package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Project struct {
	Name        string
	Description string
	URL         string
}

type Portfolio struct {
	Name     string
	Tagline  string
	Bio      string
	Goal     string
	Projects []Project
}

//go:embed templates/index.html
var files embed.FS

func main() {
	data := Portfolio{
		Name:    "Joao Vitor | Trobukan",
		Tagline: "Desarrollador Backend Golang",
		Bio:     "Soy desarrollador backend brasileño especializado en Go. Construyo herramientas CLI y servicios simples, con foco en código claro y fácil de mantener, y trabajo a diario en Linux (Void Linux, Neovim).",
		Goal:    "Busco mi oportunidad profesional en Chile. Quiero establecerme en el país, por lo que mi intención es de largo plazo. Disponible para trabajo remoto o reubicación.",
		Projects: []Project{
			{
				Name:        "docv",
				Description: "Visor CLI de documentación local en Go",
				URL:         "https://github.com/trobukan/docv",
			},
		},
	}

	tmpl := template.Must(template.ParseFS(files, "templates/index.html"))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, data)
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("rodando en el port: %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
