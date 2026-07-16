# Ejercicio Obligatorio 2: Inversión de Palabras en Posiciones Impares

## Enunciado
Realizar un programa que reciba una frase e invierta únicamente las palabras ubicadas en posiciones impares (contando desde 1)

- **Entrada:** `Qué lindo día es hoy.` 
- **Salida:** `éuQ lindo aíd es yoh.` 

---

## Mapeo de Índices en Memoria RAM

En Go, los slices se indexan en base 0, pero el enunciado plantea una escala humana en base 1[cite: 61]. Esto requiere adaptar el índice en el bucle:

```text
Índices de Go:        [0]        [1]        [2]       [3]       [4]
Contenido:         | "Qué" |  "lindo" |  "día"  |  "es"  | "hoy." |
Posición Usuario:     1ª         2ª         3ª        4ª        5ª
Lógica:             Impar       Par       Impar      Par      Impar
Condición:        (i+1)%2==1  (i+1)%2==0  (i+1)%2==1   ...       ...
```

## Observaciones Técnicas

* **Slices vs Vectores:** En Go no existe el tipo de dato "vector"; la estructura dinámica equivalente es el **Slice** (`[]type`). La función `strings.Fields` retorna internamente un tipo `[]string`, reservando memoria en el *heap* de forma transparente.
* **Separación de Palabras:** `strings.Fields` remueve de forma automática cualquier secuencia de espacios en blanco consecutivos. La reconstrucción final de la frase se realiza en un solo paso mediante `strings.Join`.
* **Inversión In-Place:** La función `invertirPalabra` no consume memoria RAM auxiliar. Utiliza un algoritmo clásico de dos punteros opuestos (`i` y `j`) que intercambian sus elementos hasta cruzarse en el centro.

```text
Intercambio en "Qué":
Paso 1:  i=0, j=2  -->  ['Q', 'u', 'é']  --> Intercambia [0] por [2]
Paso 2:  i=1, j=1  -->  ['é', 'u', 'Q']  --> Termina (i >= j)
```

* **Soporte UTF-8 (Runas):** La conversión de strings a []rune es obligatoria. Caracteres acentuados como é o í ocupan 2 bytes en memoria UTF-8. Invertir bytes directamente corrompería la codificación ; trabajar con runas garantiza que cada carácter gráfico se trate como una sola unidad de 32 bits.  
* **Signos de Puntuación:** El punto final o cualquier signo adherido a una palabra (como el . en "hoy.") se detecta como parte de la palabra y se invierte junto con ella (".yoh"), respetando el comportamiento físico del string.