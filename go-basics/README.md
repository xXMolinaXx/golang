# Apuntes de Go

Este archivo agrupa los conceptos clave, ejemplos y comandos útiles para trabajar con proyectos en Go.

## Índice

- [Crear un módulo](#crear-un-módulo)
- [Estructura básica](#estructura-básica)
- [Cómo ejecutar código](#cómo-ejecutar-código)
- [Importación de paquetes](#importación-de-paquetes)
- [Comandos básicos del CLI de Go](#comandos-básicos-del-cli-de-go)
- [Buenas prácticas](#buenas-prácticas)

## Crear un módulo

1. En la raíz del proyecto, ejecuta:

```sh
go mod init nombre/del/modulo
```

2. Esto crea el archivo `go.mod`.
3. Para mantener las dependencias limpias y actualizadas:

```sh
go mod tidy
```

> Tip: Si cambias el nombre del módulo o mueves archivos, revisa `go.mod` y usa `go mod tidy`.

## Estructura básica

```go
package main

import "fmt"

func main() {
	fmt.Println("Hola, Go!")
}
```

- `package main` define un ejecutable.
- La función `main` es el punto de entrada.
- Si el paquete no es `main`, se construye como librería y no se ejecuta directamente.

## Cómo ejecutar código

- `go run archivo.go` - ejecuta un solo archivo Go.
- `go run .` - ejecuta el paquete actual desde el directorio donde está `go.mod`.
- `go run ./...` - ejecuta todos los paquetes dentro del módulo.
- `go build` - compila el paquete actual y genera un ejecutable local.
- `go build -o nombrePrograma` - compila y escribe el binario con el nombre indicado.
- `./nombrePrograma` - ejecuta el binario ya compilado.
- `go test ./...` - ejecuta todas las pruebas del módulo.
- `go test ./ruta/al/paquete` - ejecuta las pruebas de un paquete específico.

> Nota: `go run .` es el comando más práctico cuando trabajas en un proyecto con varios archivos dentro de un mismo paquete.

## Importación de paquetes

En Go existen tres tipos principales de paquetes:

### 1. Paquetes estándar
Vienen con Go, como `fmt`, `os`, `math`, `time`.

```go
import "fmt"
import "os"
```

### 2. Paquetes locales (de tu propio proyecto)
Se importan usando la ruta del módulo definida en `go.mod`.

Supón que tu módulo se llama `example/example` y tienes un archivo en `utils/algebra.go`:

```go
import "example/example/utils"
```

Entonces puedes usarlo así:

```go
res := utils.Add(2, 3)
```

### 3. Paquetes de terceros
Se instalan con `go get` o se agregan automáticamente al compilar si ya están en `go.mod`.

```sh
go get github.com/gorilla/mux
```

```go
import "github.com/gorilla/mux"
```

### Notas sobre imports
- Los imports deben ir después de la declaración del paquete.
- Si importas un paquete y no lo usas, Go fallará en la compilación.
- Puedes usar alias para evitar conflictos:

```go
import m "github.com/gorilla/mux"
```

## Comandos básicos del CLI de Go

- `go mod init nombre/del/modulo` - inicializa el módulo y crea `go.mod`.
- `go mod tidy` - agrega dependencias usadas y elimina las no usadas.
- `go get paquete@versión` - descarga o actualiza un paquete.
- `go list ./...` - muestra los paquetes del módulo.
- `go fmt ./...` - formatea todos los archivos Go en el módulo.
- `go env GOPATH` - muestra el valor de `GOPATH`.
- `go env GOMOD` - muestra la ruta al archivo `go.mod`.
- `go clean` - limpia los archivos generados por compilación.
- `go install ./...` - compila e instala los paquetes en `GOBIN`.
- `go vet ./...` - analiza el código en busca de problemas comunes.
- `go test ./...` - ejecuta las pruebas del módulo.

## Buenas prácticas

- Mantén una sola responsabilidad por paquete.
- Usa nombres claros y exporta solo lo necesario.
- `go fmt` y `go vet` son pasos rápidos que ayudan a mantener el código limpio.
- No uses paquetes innecesarios: `go mod tidy` ayuda a limpiar dependencias.
- Prefiere `go run .` cuando trabajas con proyectos que tienen múltiples archivos en el mismo paquete.
- Para proyectos más grandes, organiza el código en subpaquetes bajo el módulo.
