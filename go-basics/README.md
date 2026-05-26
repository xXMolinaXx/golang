## Importación de paquetes en Go

En Go, puedes importar diferentes tipos de paquetes:

### 1. Paquetes estándar
Son los que vienen con Go, como `fmt`, `os`, `math`, etc.

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

Luego puedes usar las funciones exportadas de ese paquete:

```go
res := utils.Add(2, 3)
```

### 3. Paquetes de terceros
Son paquetes externos que instalas con `go get`.

Ejemplo:

```sh
go get github.com/gorilla/mux
```

Luego los importas así:

```go
import "github.com/gorilla/mux"
```

### Notas
- Los imports deben ir después de la declaración del paquete.
- Si importas pero no usas un paquete, Go dará error.
- Puedes usar alias para evitar conflictos:

```go
import m "github.com/gorilla/mux"
```
# Apuntes de Go

Este repositorio está dedicado a tomar apuntes, ejemplos y conceptos clave sobre el lenguaje de programación Go (Golang). Aquí encontrarás ejemplos prácticos, explicaciones y fragmentos de código útiles para repasar y aprender Go.

## Conceptos importantes de Go

### Público y privado
- En Go, los identificadores (variables, funciones, structs, etc.) que comienzan con mayúscula son **públicos** (exportados).
- Los que comienzan con minúscula son **privados** (no exportados fuera del paquete).

```go
type Persona struct { // Público
	Nombre string // Público
	edad   int    // Privado
}

func Sumar(a, b int) int { // Pública
	return a + b
}

func restar(a, b int) int { // Privada
	return a - b
}
```

### Crear un módulo
1. Ejecuta `go mod init nombre/del/modulo` en la raíz del proyecto.
2. Esto crea el archivo `go.mod`.

### Importar paquetes
```go
import "fmt"
import "mi/modulo/utils"
```

### Estructura básica de un programa Go
```go
package main

import "fmt"

func main() {
	fmt.Println("Hola, Go!")
}
```

### Otras notas
- Los archivos deben pertenecer a un paquete (`package nombre` al inicio).
- El punto de entrada es la función `main` en el paquete `main`.
- Go fomenta la simplicidad y la claridad en el código.
