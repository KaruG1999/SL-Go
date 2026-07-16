## Enunciado

**9.** Realice un programa que reciba una frase e imprima en pantalla la misma frase reemplazando las ocurrencias de "jueves" por "martes" respetando las letras minúsculas o mayúsculas de la palabra original en su posición correspondiente. Por ejemplo, se reemplaza "Jueves" por "Martes" o "jueveS" por "marteS".

## Observaciones

- Un string en Go es una secuencia de bytes UTF-8, no de caracteres: hay que convertir a `[]rune` antes de indexar, si no las posiciones se corren con caracteres multibyte.
- `strings.Builder` arma el resultado de forma eficiente en vez de concatenar strings en cada paso.
- La comparación para encontrar la ocurrencia es case-insensitive (`strings.ToLower`), pero el casing que se escribe en el reemplazo sale de la ocurrencia original, no de un valor fijo.

