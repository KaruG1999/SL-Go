# Ejercicio 9 — Lista Enlazada de Enteros

## Enunciado

Usando memoria dinámica con punteros, implementar una lista enlazada de enteros con las operaciones:

```go
func New() List
func IsEmpty(self List) bool
func Len(self List) int
func FrontElement(self List) int
func Next(self List) List
func ToString(self List) string
func PushFront(self List, elem int)
func PushBack(self List, elem int)
func Remove(self List) int
func Iterate(self List, f func(int) int)
```

**b)** Programa que use las operaciones.

**c)** Investigar y usar el paquete `container/list`. Diferencias con la implementación propia.

**d)** Mejorar la interfaz usando métodos y códigos de error.

---

## Lógica de resolución (`main.go` / `b/`)

```go
type nodo struct {
    elem int
    sig  *nodo
}

type List struct {
    pri *nodo
    len int
}

func New() List { return List{nil, 0} }

func IsEmpty(l List) bool { return l.pri == nil && l.len == 0 }
func Len(l List) int      { return l.len }
func FrontElement(l List) int { return l.pri.elem } // no chequea lista vacía
func Next(l List) List {
    if l.pri == nil || l.pri.sig == nil {
        return New()
    }
    return List{pri: l.pri.sig, len: l.len - 1}
}
```

`PushFront`, `PushBack`, `Remove` e `Iterate` ya están escritas como métodos con receiver puntero (`*List`), porque necesitan modificar la lista:

```go
func (l *List) PushFront(elem int) {
    nuevo := &nodo{elem: elem, sig: l.pri}
    l.pri = nuevo
    l.len++
}

func (l *List) PushBack(element int) {
    if IsEmpty(*l) {
        l.PushFront(element)
        return
    }
    nuevo := &nodo{elem: element}
    actual := l.pri
    for actual.sig != nil {
        actual = actual.sig
    }
    actual.sig = nuevo
    l.len++
}

func (l *List) Remove() (int, error) {
    if l.pri == nil {
        return 0, errors.New("lista vacia")
    }
    valor := l.pri.elem
    l.pri = l.pri.sig
    l.len--
    return valor, nil
}
```

`b/main.go` es igual pero con un `main` más completo: recorre con `Next`, prueba `Iterate`, y vacía la lista con `Remove` hasta ver el error de lista vacía.

## Parte c — diferencias con `container/list`

- Es **doblemente enlazada**: cada elemento tiene `Next()` y `Prev()`, se puede recorrer para atrás. La nuestra es simple, solo para adelante.
- `PushBack`/`PushFront` devuelven el `*Element` creado. Eso permite guardar una referencia a un nodo puntual y después borrarlo o insertar al lado (`InsertBefore`, `InsertAfter`) sin recorrer nada. La nuestra no expone los nodos: no hay forma de agarrar "el tercer elemento" desde afuera y sacarlo directo, solo `Remove()` del frente.
- Guarda `any` en vez de `int`, así que una misma lista puede tener tipos mezclados — pero a cambio perdés el chequeo de tipos en compilación: hay que hacer type assertion o mirar `%T` para saber qué hay en cada nodo. La nuestra, al ser solo `int`, no tiene ese problema.
- Tiene más operaciones ya resueltas (`MoveToFront`, `MoveToBack`, `PushBackList` para unir dos listas), nosotros solo escribimos lo que pedía el enunciado.
- Para que la propia fuera genérica (aceptar cualquier tipo sin perder el chequeo de tipos) alcanzaría con type parameters de Go 1.18+: `type List[T any] struct { pri *nodo[T]; len int }`.

## Parte d — versión con métodos y errores

`d/main.go` reescribe todo como métodos sobre `List`, sin funciones sueltas, y le agrega manejo de error a las operaciones que antes podían fallar en silencio o hacer panic:

```go
func (l List) FrontElement() (int, error) {
    if l.IsEmpty() {
        return 0, errors.New("lista vacía")
    }
    return l.pri.elem, nil
}

func (l List) Next() (List, error) {
    if l.IsEmpty() {
        return List{}, errors.New("lista vacía")
    }
    return List{pri: l.pri.sig, len: l.len - 1}, nil
}
```

También define `String()` (interfaz Stringer), así que `fmt.Printf("%s", l)` funciona sin llamar a `ToString` a mano.

## ¿Por qué acá se muta con puntero y en el árbol (ej11) se reasigna?

Dos formas válidas de resolver "esta operación cambia la estructura":

- **Acá (List)**: los métodos que modifican reciben `*List` y tocan los campos directamente (`l.pri = nuevo`). Se llama `lista.PushFront(5)` y ya está, no hace falta reasignar nada.
- **En el árbol (ej11)**: `Add` no muta nada por puntero, devuelve un `arbolBin` nuevo y hay que reasignar (`a = a.Add(5)`). Si te olvidás del `a =`, el insert se pierde en silencio.

La lista muta por puntero porque tiene sentido pensarla como una caja que cambia con el tiempo (como un slice o un stack). El árbol usa retorno porque `this` se pasa por valor (no por puntero) en toda su interfaz — siguiendo la firma que pide el enunciado (`func (this IntBinTree) Add(elem int) IntBinTree`) — así que la única forma de que el cambio sea visible es devolviendo la versión nueva.

## Observaciones

- En `main.go`/`b`, `FrontElement` no chequea lista vacía: si se llama sobre una lista vacía, `l.pri.elem` explota con nil pointer dereference. Es justo lo que la parte d viene a arreglar devolviendo `error` en vez de dejar que explote.
- `PushFront` y `PushBack` reciben `*List` porque agregan un nodo y cambian el field `pri` (o `len`); las funciones de solo lectura (`IsEmpty`, `Len`, `FrontElement`, `ToString`) reciben `List` por valor porque no necesitan modificar nada.
- `PushBack` recorre toda la lista hasta el final (`O(n)`), a diferencia de `PushFront` que es `O(1)`. Si preguntan por eficiencia, ese es el punto a mencionar.
