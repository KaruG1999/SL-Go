## Enunciado

**6.** Escriba un programa que lee desde la entrada estándar dos enteros y retorne la división entre el mayor de ellos y el menor. Realizar el mismo programa considerando que se leen dos enteros sin signo. Luego modifique para que trabaje con reales (punto flotante). Ver que sucede con las división por cero.

## Observaciones

- División entera por cero produce **panic** en tiempo de ejecución en Go, hay que chequear el divisor antes.
- División float por cero no hace panic: da `+Inf` o `-Inf` según el signo.
- Con `uint` no hay negativos, así que "mayor" y "menor" siempre son >= 0; si se ingresa un negativo el comportamiento es indefinido (underflow), Go no lo detecta solo.
