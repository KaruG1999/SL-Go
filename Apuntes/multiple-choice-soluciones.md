# Multiple choice — soluciones

**1. B.** Concurrencia = estructurar tareas "en progreso" a la vez (pueden turnarse en 1 core). Paralelismo = ejecución literalmente simultánea, requiere varios cores. Podés tener concurrencia sin paralelismo, no al revés (por eso C está al revés).

**2. C.** `Kilometros` y `Millas` son tipos nombrados distintos, incompatibles entre sí aunque los dos tengan `float64` de tipo subyacente. Da error de compilación (`mismatched types`).

**3. B.** Ambos son short-circuit. Verificado en la teoría Go-1.

**4. C.** `v` fue declarada en la sentencia de inicialización del `if`, su scope es solo ese `if`/`else`. El segundo `Println(v)` da `undefined: v`.

**5. B.** Go no hace fallthrough automático entre `case`. Entra a `case 2`, imprime `dos`, termina el switch ahí.

**6. C.** Esa es la regla exacta: `Println` siempre separa todo con espacio y agrega `\n`; `Print` solo agrega espacio entre dos argumentos si ninguno es un string.

**7. B.** `Sscanln` corta apenas encuentra un salto de línea antes de terminar de leer los valores esperados. Da error `"unexpected newline"` y la variable que no llegó a leer queda con su valor anterior (en este caso `y` no tenía valor previo asignado, así que queda en `""`, el zero value — la clave de la pregunta es que corta con error y no llega a asignar `800`).

**8. B.** Tamaño fijo (parte del tipo) vs. referencia a un array de atrás que puede crecer.

**9. B.** Confirmado con `go run`: `len(b)=2, cap(b)=4` antes del append; como hay `cap` de sobra, `append` escribe sobre `a[3]` (pisa el `4` con el `100`) en vez de crear un array nuevo.

**10. C.** Escribir en un map `nil` es un panic en tiempo de ejecución, no un error de compilación.

**11. C.** Slices, maps y funciones no son comparables con `==`, por eso no pueden ser clave. `bool`, `arrays` de tipos comparables y `structs` con todos los campos comparables sí pueden.

**12. B.** El receiver es puntero (`*Contador`), así que `Incrementar` sí modifica `c.valor` de verdad (Go llama con `&c` automáticamente). Da `2`.

**13. B.** `Sonido()` está definido sobre `*Perro`. Solo `*Perro` implementa `Animal`; `Perro{}` (valor, sin `&`) no compila en `var a Animal = Perro{}`. Confirmado con `go build`.

**14. B.** `i = t` (un `*T` nil) le da a la interfaz un tipo concreto (`*T`), aunque el valor sea nil — por eso `i == nil` da `false`. Y `M()` sí se puede llamar porque el método chequea `t == nil` por dentro y no desreferencia nada.

**15. C.** Definición correcta y completa.

**16. B.** Confirmado con `go run`: imprime `fin` primero (los defer corren al final de main, en orden LIFO), después `2`, `1`, `0`.

**17. B.** `recover()` solo tiene efecto dentro de una función `defer`; en cualquier otro contexto devuelve `nil` sin hacer nada.

**18. B.** La función anónima devuelta capturó `x` por clausura (closure); todas las llamadas a la misma `f` comparten esa `x`, por eso `1, 4, 9` (1², 2², 3²) en vez de repetir `1`.

**19. B.** Hay que "expandir" el slice con `...`: `sum(values...)`. Sin eso no compila.

**20. B.** Las claves de un map en Go siempre tienen que ser comparables con `==`; por eso a `K` se le exige el constraint `comparable`.

**21. B.** Go infiere los parámetros de tipo mirando los argumentos en la gran mayoría de los casos; poner los corchetes explícitos es opcional salvo ambigüedad.

**22. B.** Confirmado con `go run` (Go 1.24, que es la versión del proyecto): desde Go 1.22 cada iteración del `for` tiene su propia copia de la variable, así que imprime `0`, `1`, `2` una vez cada uno (en algún orden, por el scheduling de goroutines) — ya NO es el bug clásico de versiones viejas de Go donde todas compartían la misma variable.

**23. B.** `Withdraw` toma el lock y sin soltarlo llama a `Deposit`, que intenta tomar el mismo lock de nuevo. `sync.Mutex` no es reentrante → deadlock.

**24. A.** Exclusión mutua, retención y espera, no apropiación, espera circular. Hace falta que se cumplan las 4 juntas.

**25. B.** `select` con `default` es no bloqueante: si esa alternativa no está lista ahora, cae al `default` en vez de esperar — así el cliente puede probar varias alternativas en cascada sin bloquearse en ninguna.

**26. C.** `go mod init miproyecto` crea `go.mod`, que define el nombre del módulo (para resolver imports propios) y lleva registro de las dependencias externas y sus versiones.

**27. B.** Confirmado con `go run`: un string es una secuencia de bytes UTF-8, y `é` ocupa 2 bytes, entonces `len(s)` (bytes) da `5` y `len([]rune(s))` (caracteres reales) da `4`.

**28. B.** Indexar un string con `s[i]` devuelve el byte en esa posición, no el carácter. Como `é` ocupa 2 bytes, esos dos bytes se imprimen como caracteres sueltos sin sentido (algo como `Ã©`) en vez de mostrar `é` de corrido. Para evitarlo hay que convertir a `[]rune(s)` antes de indexar.

**29. B.** Un rune literal (`'a'`, comillas simples) es de tipo `rune` (alias de `int32`). `fmt.Println` sin verbo especial muestra el valor numérico (el code point Unicode, 97 para `'a'`), no el carácter — para verlo como carácter hace falta `fmt.Printf("%c\n", r)`.

**30. B.** Confirmado con `go run`: da `123`, pegado. La regla de "agregar espacio entre argumentos que no son strings" aplica solo *dentro* de un mismo llamado a `Print` (ej: `fmt.Print(1, 2, 3)` sí da `1 2 3`, porque ahí los 3 son argumentos de la misma llamada). Acá son tres llamados independientes, cada uno con un único argumento — no hay par de argumentos entre los cuales meter un espacio, y `Print` no sabe nada de la llamada anterior ni la siguiente.
