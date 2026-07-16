## Enunciado

**Obligatorio 3.** Realice un programa que reciba una palabra como argumento y lee de la entrada una frase. Luego, el programa debe imprimir la frase que leyó con cada una de las ocurrencias de la palabra con las mayúsculas y minúsculas invertidas. Por ejemplo, si la frase es:

`Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO`

y la palabra es "PEQUEÑO" entonces el programa imprimirá:

`Parece PEQueÑO, pero no es tan PEQUEñO el pequeño`

Tenga en cuenta que la palabra a buscar puede ser ingresada con mayúsculas y minúsculas mezcladas.

**Este es el ejercicio que hay que entregar.**

## Observaciones Técnicas

* **Paso de Argumentos por Consola (`os.Args`):** Se utiliza el slice `os.Args` para capturar la palabra clave desde la terminal antes de la ejecución del programa. Se implementa una validación previa de longitud (`len(os.Args) < 2`) para evitar un *runtime panic* si el usuario omite pasar el parámetro requerido.
* **Comparación Case-Insensitive Eficiente:** En lugar de forzar conversiones redundantes de texto a minúsculas, se utiliza la función estándar `strings.EqualFold`. Esta función realiza una comparación rápida bajo el plegado de caracteres Unicode (*case-folding*), aislando la lógica de búsqueda de las variaciones del casing.
* **Uso de `strings.Builder`:** Se erradica por completo la concatenación tradicional con `+` (que genera alocaciones constantes y sobrecarga de basura en el *heap*) reemplazándola por `strings.Builder` y escrituras directas mediante `WriteRune` y `WriteString`.
* **Tratamiento Seguro de UTF-8:** La frase y el patrón de búsqueda se transforman a `[]rune` antes del escaneo. Esto previene desalineaciones de índices al procesar palabras acentuadas (como "pequeño") cuyos caracteres especiales ocupan más de 1 byte en UTF-8.

## Entrada de Datos y Flujos del Sistema Operativo

El programa utiliza dos canales independientes para recibir información del exterior, combinando argumentos de inicio con flujos de entrada interactivos:

```text
Línea de comandos (os.Args):
  $ go run ejOb3.go PEQUEÑO
                      ^
                      └── os.Args[1] (Argumento de inicio del SO)

Entrada estándar (os.Stdin):
  Ingrese una frase: "Parece peqUEño..."  <-- (Flujo dinámico en ejecución)
 ``` 

## 1. Argumentos de Línea de Comandos (os.Args)
* **Qué es:** Slice de strings ([]string) del paquete os con los parámetros de ejecución.

* **Distribución:** os.Args[0] es la ruta del ejecutable; os.Args[1] es el primer argumento real (la palabra buscada).

* **Seguridad:** Es obligatorio validar len(os.Args) < 2 antes de asignar la variable para evitar un crash (runtime panic) si el usuario olvida pasar el argumento.

## 2. Comparación de Flujos
* **Por Consola (os.Args):** Se definen antes del inicio, son estáticos y se leen directo del entorno. Ideal para configuraciones primarias.

* **Entrada Estándar (os.Stdin):** Se define durante la ejecución. Frena el hilo del programa esperando la escritura del usuario (o redirección mediante < archivo.txt). Ideal para procesar texto dinámico.