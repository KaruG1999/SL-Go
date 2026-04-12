## Diferencias: Lenguajes Compilados vs. Interpretados

La distinción principal reside en el **momento de la traducción** del código fuente a instrucciones de máquina ($ISA$).

---

## 1. Lenguajes Compilados
El código fuente se traduce en su totalidad por un **compilador** antes de la ejecución, generando un archivo binario independiente.

* **Flujo:** `Código Fuente` $\rightarrow$ `Compilador` $\rightarrow$ `Código Máquina (Binario)` $\rightarrow$ `Ejecución`.
* **Rendimiento:** Alta eficiencia. El procesador ejecuta instrucciones nativas sin intermediarios.
* **Detección de errores:** Los errores sintácticos y de tipado se reportan en tiempo de compilación (*Build time*).
* **Ejemplos:** C, C++, Rust, Swift, Fortran.



---

## 2. Lenguajes Interpretados
Un programa externo denominado **intérprete** procesa el código fuente y lo ejecuta línea por línea en tiempo de ejecución (*Runtime*).

* **Flujo:** `Código Fuente` $\rightarrow$ `Intérprete` $\rightarrow$ `Ejecución inmediata`.
* **Portabilidad:** Alta. El mismo script corre en cualquier arquitectura que tenga el intérprete instalado.
* **Rendimiento:** Menor velocidad comparada con binarios nativos debido al *overhead* de la traducción constante.
* **Ejemplos:** Python, Ruby, PHP, JavaScript (V8 utiliza JIT, pero se categoriza aquí).



---

## 3. Clasificación de Go (Golang)

**Go es un lenguaje compilado.**

A pesar de que su velocidad de compilación es extremadamente alta (diseñada para simular la agilidad de lenguajes dinámicos), su naturaleza es estrictamente estática y compilada.

### Puntos clave de Go:
* **Binarios Estáticos:** Genera un único archivo ejecutable que contiene todas las dependencias necesarias.
* **Tipado Estático:** La validación de tipos ocurre antes de generar el binario.
* **Sin VM:** A diferencia de Java (Bytecode + JVM), Go compila directamente a código máquina específico del sistema operativo y arquitectura destino ($GOOS$ / $GOARCH$).

---
**Diagnóstico:** Si buscas performance en sistemas distribuidos o microservicios, la compilación de Go ofrece ventajas críticas en tiempos de arranque y consumo de memoria.

¿Necesitas que profundice en la compilación cruzada (*cross-compilation*) de Go para distintos entornos?