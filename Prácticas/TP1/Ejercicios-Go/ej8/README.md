## Enunciado

**8.** Realizar un programa que lea el punto cardinal (como caracter o string) del cual viene el viento ('N', 'S', 'E', 'O') y envíe a la salida estándar hacia cuál se dirigiría.

a. ¿Cómo se escribe el default en el case de otros lenguajes?

Sub-objetivo: Uso de case con la opción por default. E/S caracteres o strings.

## Observaciones

- El viento "viene de" la dirección leída, pero "se dirige hacia" la opuesta (N → sopla hacia el S, etc.).
- `default` en el `switch` de Go equivale al `else` final de una cadena `if/else if`; en otros lenguajes es `default` (Java, C, C#) o similar.
- Se puede aceptar mayúscula y minúscula en el mismo `case` separando valores por coma (`case "N", "n":`), sin normalizar el string antes.
