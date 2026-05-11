# Ejercicio 2 — Factorial

## Enunciado

Implementar la función factorial de dos formas: una **iterativa** y otra **recursiva**. Escribir un programa que evalúe ambas de 0 a 9.

La función factorial se define como:
```
0! = 1
n! = n × (n-1)!
```

---

## Lógica de resolución

### Versión iterativa

```go
func factorialIter(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}
```

### Versión recursiva

```go
func factorialRec(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorialRec(n-1)
}
```

### Programa principal

```go
func main() {
    for i := 0; i <= 9; i++ {
        fmt.Printf("%d! = %d (iter) = %d (rec)\n",
            i, factorialIter(i), factorialRec(i))
    }
}
```

> El punto clave: ambas versiones producen el mismo resultado. La recursiva es más directa como traducción de la definición matemática; la iterativa evita el overhead de llamadas en la pila. En Go no hay optimización de tail recursion, así que la iterativa es preferible para valores grandes.
