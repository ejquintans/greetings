# Ejemplo de GO

Este paquete proporciona una forma simple de obtener saludos personalizados en Go que hice de practica.

## Instalación
Ejecuta el siguiente comando parra instalar el paquete:
```bash
go get -u github.com/ejquintans/greetings
```

## Uso
Ejemplo de uso:
```go
package main

import (
    "fmt"
    "github.com/ejquintans/greetings
)

func main() {
    message, err := greetings.Hello("Pichu")

    if err != nil {
        fmt.Println("Ocurrio un error: ", err)
        return
    }

    fmt.Println(message)
}
```