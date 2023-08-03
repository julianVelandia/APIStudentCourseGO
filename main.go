package main

import (
	"fmt"
	"os"

	router "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/app"
)

func main() {

	r := router.NewRouter()
	err := r.Run(os.Getenv("PORT"))

	if err != nil {
		fmt.Println("error starting server: ", err)
	}
}
