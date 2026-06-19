# Organización de Proyectos Go - Estándares de la Comunidad

> Basado en [Standard Go Project Layout](https://github.com/golang-standards/project-layout) y las mejores prácticas aceptadas por la comunidad Go.

## Estructura Estándar

```
project-name/
├── cmd/                    # Aplicaciones ejecutables
│   ├── app-name/
│   │   └── main.go
│   └── tool-name/
│       └── main.go
├── internal/              # Código privado (no reutilizable externamente)
│   ├── app/              # Lógica específica de la aplicación
│   ├── config/           # Parseo y lectura de configuración
│   ├── cache/            # Implementaciones de cache
│   ├── db/               # Lógica de base de datos
│   ├── log/              # Logging
│   └── models/           # Modelos de negocio locales
├── pkg/                   # Código público (reutilizable)
│   ├── api/              # API clients
│   ├── encoding/         # Codificación (JSON, XML, etc)
│   └── util/             # Utilidades genéricas
├── api/                   # OpenAPI/Swagger specs, JSON schemas
├── web/                   # Aplicación web (HTML, CSS, JS)
├── configs/              # Archivos de configuración
├── test/                  # Tests adicionales y fixtures
├── docs/                  # Documentación del proyecto
├── examples/             # Ejemplos de uso
├── third_party/          # Herramientas externas
├── .github/              # Workflows de GitHub, templates
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Directorio por Directorio

### **cmd/** - Punto de Entrada
La carpeta `cmd/` contiene las **aplicaciones ejecutables**. La idea es tener poco código aquí:
- Cada subdir = un ejecutable diferente
- Solo importa desde `internal/` y `pkg/`
- `main()` debe ser limpia y corta

```
cmd/
├── server/main.go         # API/servidor web
├── cli/main.go            # Herramienta CLI
└── worker/main.go         # Worker/job processor
```

### **internal/** - Código Privado (IMPORTANTE)
Cualquier paquete en `internal/` **no puede ser importado desde fuera del módulo**. Es una restricción del compilador de Go.

**Estructura típica:**
```
internal/
├── app/                   # Lógica de aplicación
│   └── server.go
├── config/
│   └── config.go
├── db/
│   ├── postgres.go
│   └── migrations/
├── models/
│   ├── user.go
│   ├── product.go
│   └── order.go
├── service/               # Lógica de negocio
│   ├── user_service.go
│   └── product_service.go
├── handler/               # HTTP handlers (para APIs)
│   ├── user_handler.go
│   └── product_handler.go
└── middleware/
    └── auth.go
```

### **pkg/** - Código Público
Código que podría ser reutilizado en otros proyectos.

```
pkg/
├── validator/            # Validación
├── logger/               # Logger
├── errors/               # Tipos de error
└── cache/                # Implementación de cache
```

## Convenciones de Nombres

### Nombres de Paquetes
- ✅ Cortos, claros, una palabra
- ✅ Minúsculas, sin underscores
- ✅ Descriptivos
- ❌ NO: `util`, `helper`, `common`, `misc`

**Ejemplos buenos:**
```go
package auth        // autenticación
package validation  // validación
package mail        // email
package crypto      // criptografía
package storage     // almacenamiento
```

### Nombres de Archivos
- Minúsculas
- Underscores para separar conceptos: `user_service.go`, `user_handler.go`
- Tests: `filename_test.go`

## Patrones de Organización por Tipo de Proyecto

### REST API
```
internal/
├── app/
│   └── server.go          # Inicialización del servidor
├── config/
│   └── config.go
├── handler/               # HTTP handlers
│   ├── user_handler.go
│   ├── product_handler.go
│   └── middleware.go
├── service/               # Lógica de negocio
│   ├── user_service.go
│   └── product_service.go
├── repository/            # Acceso a datos
│   ├── user_repo.go
│   └── product_repo.go
└── models/
    ├── user.go
    └── product.go
```

### CLI Tool
```
cmd/
└── main/
    └── main.go

internal/
├── app/
│   └── cli.go
├── command/
│   ├── serve.go
│   ├── config.go
│   └── migrate.go
└── config/
```

### Librería/Package
```
pkg/
├── option.go              # Opciones públicas
├── main_type.go           # Tipo principal
├── helpers.go             # Funciones públicas
└── internal/              # Detalles internos
    └── implementation.go
```

## Reglas Clave de la Comunidad

### 1. **Mantén `main` limpia**
```go
// ❌ NO hacer lógica compleja en main
func main() {
    db := setupDatabase()
    cache := setupCache()
    // ... mucho más código aquí
}

// ✅ Delega a otros packages
func main() {
    app, err := app.NewApp(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    app.Start()
}
```

### 2. **Tests cerca del código**
```
internal/
└── service/
    ├── user_service.go
    └── user_service_test.go  # En el mismo package
```

### 3. **NO hay código flotante en raíz**
```
❌ Evitar:
├── helper.go
├── util.go
├── common.go
└── types.go

✅ Organizar:
internal/
├── models/types.go
├── pkg/helper/helper.go
└── pkg/util/util.go
```

### 4. **Interfaces en el mismo package donde se usan**
```go
// En internal/service/user_service.go
type UserRepository interface {  // Interfaz que NECESITA el servicio
    GetUser(id string) (*User, error)
}

// En internal/repository/user_repo.go
type PostgresUserRepository struct {...}

func (r *PostgresUserRepository) GetUser(id string) (*User, error) {...}
```

### 5. **Ciclo de dependencias: Handler → Service → Repository**
```
handler → service → repository → database
   ↑         ↑            ↑
   └─── modelos compartidos en models/
```

## Proyecto No Go-Standard ❌

```
❌ EVITAR ESTO:
├── utils/          # Demasiado genérico
├── helpers/        # Demasiado genérico
├── lib/            # Ambiguo
├── common/         # Ambiguo
├── core/           # Ambiguo
└── tools/          # No es cmd/
```

## Checklist de Buen Diseño

- [ ] Código ejecutable en `cmd/`
- [ ] Código privado en `internal/`
- [ ] Código público en `pkg/`
- [ ] Tests al lado del código que testean
- [ ] Nombres de paquetes claros y específicos
- [ ] Sin `utils/`, `helpers/`, `common/`
- [ ] Sin ciclos de importación
- [ ] Interfaces en el paquete que las necesita
- [ ] Modelos en archivos separados por tipo
- [ ] `main.go` limpia y corta

## Referencias Oficiales

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go - Package Names](https://golang.org/doc/effective_go#package-names)
- [Go Code Organization](https://golang.org/blog/organizing-go-code)
- [Package-oriented design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)
