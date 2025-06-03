package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := range 100 {
		go func(id int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("The API call, number %d, was executed\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
	fmt.Println(time.Since(start))
}

// func getOrder(id int) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("The API call, number %d, was executed\n", id)
// }

// func fetchDataFromStreamingAPI() {
// 	resp, err := http.Get("http://localhost:8080/stream")
// 	if err != nil {
// 		fmt.Println("Error al realizar la solicitud:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("Error en la respuesta del servidor:", resp.Status)
// 		return
// 	}

// 	// Utiliza una goroutine para leer y procesar los datos en tiempo real
// 	go func() {
// 		fmt.Println("Datos recibidos del Streaming API:")
// 		for {
// 			data := make([]byte, 1024) // Tamaño del búfer para leer los datos
// 			n, err := resp.Body.Read(data)
// 			if err != nil && err != io.EOF {
// 				fmt.Println("Error al leer los datos:", err)
// 				return
// 			}

// 			// Procesa los datos recibidos (puedes imprimirlos, almacenarlos en un archivo, etc.)
// 			fmt.Print(string(data[:n]))

// 			// Si se alcanzó el final de los datos, termina el bucle
// 			if err == io.EOF {
// 				break
// 			}
// 		}
// 	}()

// 	// Espera hasta que la solicitud y la goroutine finalicen
// 	select {}
// }

// func main() {
// 	fetchDataFromStreamingAPI()
// }
