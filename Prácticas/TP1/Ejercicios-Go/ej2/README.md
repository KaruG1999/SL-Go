## Enunciado

**3.** Dada la siguiente declaración de programa Go indicar si es correcta. Usar el compilador:

```go
func main () {
    /* integers */
    var zz int = 0A
    var z int := x ;
    x := 10;
    var y int8 = x+1;
    const n := 5001
    const c int := 5001
    /* float //
    var e float32 := 6
    f float32 = e
}
```

a. Si no lo es, realizar las modificaciones mínimas necesarias para que las declaraciones funcionen.

b. Enviar a la salida estándar los valores de todas las variables y constantes declaradas.

*(Esta carpeta y `ej3/` son dos resoluciones distintas del mismo enunciado.)*

## Observaciones

- `var` se declara con `=`, nunca con `:=`. `:=` es solo para declaración corta sin `var`.
- Las constantes (`const`) siempre usan `=`, jamás `:=`.
- No hay conversión implícita entre tipos: sumar un `int` a un `int8` requiere casting explícito (`int8(x + 1)`).
- `0A` no es un literal entero válido en Go.
