## Enunciado

**5.** Realizar un programa que lea un número y muestre el valor correspondiente aplicando la siguiente función sobre el mismo:

```
f(x) = |x|      x ∈ (-∞, -18)
       x mod 4  x ∈ [-18, -1]
       x²       x ∈ [1, 20)
       -x       x ∈ [20, +∞)
```

a. ¿Qué tiene de particular la función con el 0 (cero), se puede escribir sin opción default/else? Re-escribir con otra estructura de control selectiva.

b. Re-escribir la función usando punto flotante.

Sub-objetivo: Uso de E/S de enteros y punto flotante. Operaciones aritméticas sobre enteros y punto flotante (potencia, valor absoluto, negación y módulo).

## Observaciones

- El 0 no cae en ninguno de los cuatro intervalos definidos: la función no está definida ahí. Por eso el caso "sin default" no aplica igual que en otros ejercicios con `switch` — acá directamente no hay un resultado correcto para x=0.
- `switch` sin selector (`switch { case cond: ... }`) es equivalente a una cadena de `if/else if`, útil para rangos.
- Con floats no se puede usar `%`: hay que usar `math.Mod`. Tampoco potencia manual: `math.Pow`. Valor absoluto: `math.Abs`.
