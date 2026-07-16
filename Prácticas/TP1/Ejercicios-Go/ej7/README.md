## Enunciado

**7.** Las temperaturas de los pacientes de un hospital se dividen en 3 grupos: alta (mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de 36). Se deben leer 10 temperaturas de pacientes e informar el porcentaje de pacientes de cada grupo. Luego se debe imprimir el promedio entero entre la temperatura máxima y la temperatura mínima.

a. ¿Se puede utilizar el case para tipos reales en otros lenguajes?

b. ¿Cómo se realizan las conversiones entre reales (punto flotante) y enteros en otros lenguajes?

Sub-objetivo: El tipado fuerte, usar casting. Operaciones y E/S con float. Casting en otros lenguajes.

## Observaciones

- El promedio pedido es `(max+min)/2`, no el promedio de las 10 temperaturas.
- `switch` con selector no es confiable con floats porque la igualdad exacta entre reales es poco confiable (errores de redondeo). Se resuelve con `if/else if` o `switch` sin selector.
- Conversión explícita siempre: `int(f)` trunca hacia cero, `float64(i)` convierte a float. En Java/C es casteo explícito también; en Python muchas veces es implícita.
