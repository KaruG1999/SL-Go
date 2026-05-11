# Seminario de Lenguajes — Go — Packages, Modules y Dependencies

---

## Contexto de Conexión

A esta altura ya sabemos escribir programas en Go. Ahora vemos cómo organizarlos en piezas reutilizables (**packages**) y cómo gestionar esas piezas junto con dependencias externas (**modules**).

---

## Conceptos Core

- **Package**: unidad de organización del código en Go. Cada archivo pertenece a un package declarado al inicio con `package <nombre>`.
- **Identificador exportado**: todo nombre que empieza con **mayúscula** es visible desde fuera del package. Los que empiezan en minúscula son privados al package.
- **Import alias**: al importar un package se le puede dar un nombre corto: `import tc "projecttempconv/tempconv"`.
- **Module**: unidad de distribución que agrupa uno o más packages. Se define con el archivo `go.mod`.
- **`go mod init <nombre>`**: crea el archivo `go.mod` e inicializa un módulo.
- **`go get <dependencia>`**: descarga e instala un package externo y actualiza `go.mod`.
- **`go.mod`**: archivo que declara el nombre del módulo, la versión de Go y las dependencias requeridas.

---

## Desarrollo

### 1. Organizando código en packages

En Go se puede separar el código en múltiples archivos y carpetas. Cada carpeta es un package distinto.

**Estructura de ejemplo:**

```
projecttempconv/
  main.go          → package main
  tempconv/
    tempconv.go    → package tempconv
    conv.go        → package tempconv
```

**tempconv.go:**
```go
package tempconv

import "fmt"

type Celsius    float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
```

**conv.go:**
```go
package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

**main.go:**
```go
package main

import (
    "fmt"
    tc "projecttempconv/tempconv"
)

func main() {
    c := tc.BoilingC
    fmt.Println(c, "=", tc.CToF(c)) // 100°C = 212°F
}
```

> Solo `Celsius`, `Fahrenheit`, `BoilingC`, `CToF`, `CToF`, `String` son exportados (empiezan con mayúscula).

---

### 2. Modules — go.mod

Sin inicializar el módulo, Go no sabe dónde buscar el package `projecttempconv/tempconv` y da error:

```
package projecttempconv/tempconv is not in GOROOT
```

**Solución:** inicializar el módulo en la raíz del proyecto:

```bash
go mod init projecttempconv
```

Esto genera:

```
module projecttempconv
go 1.20
```

Ahora `go run ./main.go` funciona correctamente:

```
100°C = 212°F
```

---

### 3. Dependencies — paquetes externos

Para usar un package externo se descarga con `go get`:

```bash
go get -u rsc.io/quote
```

Esto actualiza `go.mod` automáticamente:

```
module projecttempconv
go 1.20

require (
    golang.org/x/text v0.9.0         // indirect
    rsc.io/quote      v1.5.2         // indirect
    rsc.io/sampler    v1.99.99       // indirect
)
```

Y en el código se importa igual que cualquier package:

```go
import (
    "fmt"
    tc "projecttempconv/tempconv"
    "rsc.io/quote"
)

func main() {
    c := tc.BoilingC
    fmt.Println(c, "=", tc.CToF(c), quote.Hello())
    // 100°C = 212°F   99 bottles of beer on the wall, ...
}
```

---

## Lo que no podés ignorar

> 1. **Mayúscula = exportado**: si una función, tipo o variable empieza con mayúscula, es pública para otros packages. Si empieza con minúscula, es privada.
> 2. **El nombre del módulo importa**: en `main.go`, el path de import `"projecttempconv/tempconv"` es relativo a la raíz del módulo definida en `go.mod`.
> 3. **`go mod init` va en la raíz del proyecto**: solo se crea un `go.mod` por módulo. Sin él, Go no resuelve los imports locales.
> 4. **`go get` actualiza `go.mod` automáticamente**: no es necesario editar el archivo a mano para agregar dependencias externas.
> 5. **Un package = una carpeta**: dos archivos en la misma carpeta deben pertenecer al mismo package. No se puede tener dos packages distintos en la misma carpeta (excepto el `_test`).
