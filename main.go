package main

import (
	"log"

	"github.com/HamelBarrer/go-react/bd"
	"github.com/HamelBarrer/go-react/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
