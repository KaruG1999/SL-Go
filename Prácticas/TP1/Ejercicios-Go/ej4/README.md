## Enunciado

**4.** Escriba un programa que imprima en la salida estándar la suma
de los primeros números positivos pares menores o iguales a
250. Cambiar el programa para que itere en el sentido contrario
pero obtener el mismo resultado. Cambiar el programa para que
en lugar de usar un literal como tope se use una constante. Si lo
desea, investigue la herramienta gofmt y pruebe sobre el código
escrito.

Sub-objetivo: Uso de E/S de valores numéricos en Go,
estructuras de control básicas, constantes y variables.

## Observaciones

- Sumar de adelante hacia atrás o al revés da lo mismo porque la suma es conmutativa.
- La constante reemplaza el literal `250`: si cambia el rango, se modifica en un solo lugar.
- `gofmt -w archivo.go` reescribe el archivo con el formato oficial; sin `-w` solo muestra el diff.