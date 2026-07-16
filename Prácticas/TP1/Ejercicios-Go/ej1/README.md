## Enunciado

**1.** Indique las diferencias entre un lenguaje interpretado y uno compilado. Indique ejemplos de cada uno si conoce. ¿En cuál catalogaría a Go?

**2.** Compile con el compilador Go el "famoso" programa Hello World/Hola Mundo. Nombre el archivo fuente `hola.go`. Compile y ejecute todo en un paso y compile generando un ejecutable con nombre `hola.exe` o directamente `hola`.

```
go run hola.go
go build -o hola hola.go
./hola
```

La respuesta teórica del punto 1 está en `ej1.md` (un nivel arriba). Esta carpeta tiene el código del punto 2, que sirve de demo práctica de por qué Go es compilado.

## Observaciones

- Go es compilado: genera un binario nativo por SO/arquitectura, sin VM de por medio (a diferencia de Java con su bytecode + JVM).
- `go run` compila y ejecuta en un solo paso, sin dejar el binario. `go build` genera el ejecutable aparte, después hay que correrlo con `./nombre`.
- Todo `.go` ejecutable es `package main` con `func main()` como punto de entrada.
