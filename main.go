package goservidores

import (
	"fmt"
	"net/http"
	"time"
)

func run() {

	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"https://platzi.com",
		"https://hellosabi.com",
		"http://google.com",
		"https://facebook.com",
		"https://instagram.com",
	}
	i := 0
	for {
		if i > 1 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	tiempo := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempo)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + " no esta OK"
	}
	canal <- servidor + " esta OK"
}
