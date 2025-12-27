# Tipos y Structs en Go: Tu Primera Inmersión en el Lenguaje

Si vienes de Java como yo, probablemente estés acostumbrado a pensar en términos de clases, herencia y objetos. Go tiene un enfoque diferente pero igualmente poderoso. En este post, exploraremos cómo Go maneja los tipos de datos estructurados a través de **structs** y por qué este enfoque puede ser más simple y flexible de lo que imaginas.

## ¿Qué son los Structs?

En Go, un **struct** es similar a una clase en Java, pero con algunas diferencias fundamentales:

- **No hay herencia**: Go no soporta herencia de clases como Java
- **Composición sobre herencia**: En lugar de extender clases, Go promueve la composición mediante **embedding**
- **Más simple**: Menos ceremonia, más código directo

Veamos un ejemplo básico:

```go
type Person struct {
    Name      string
    Age       int
    email     string  // Campo privado (minúscula)
    BirthDate time.Time `json:"birth_date"`
}
```

### Visibilidad: Mayúsculas vs Minúsculas

En Go, la visibilidad se controla con la primera letra:
- **Mayúscula** = público (como `public` en Java)
- **Minúscula** = privado (como `private` en Java)

No necesitas palabras clave como `public` o `private`. El lenguaje lo determina automáticamente.

### Tags: Las Anotaciones de Go

Los **tags** en Go son similares a las anotaciones en Java (`@JsonProperty`, `@Valid`, etc.). Te permiten agregar metadata a los campos:

```go
BirthDate time.Time `json:"birth_date" validate:"required"`
```

## Constructores Idiomáticos

En Go, no hay constructores automáticos como en Java. En su lugar, se usa una convención: funciones que empiezan con `New`:

```go
func NewPerson(name string, age int, email string) *Person {
    return &Person{
        Name:      name,
        Age:       age,
        email:     email,
        BirthDate: time.Now(),
    }
}
```

Nota el uso del operador `&` que retorna un puntero. En Go, es común trabajar con punteros para evitar copias innecesarias de estructuras grandes.

## Composición vs Herencia: Struct Embedding

Aquí está una de las diferencias más importantes con Java. En lugar de usar herencia (`extends`), Go usa **embedding** o composición:

```go
type Employee struct {
    Person              // Embedding: Employee "tiene un" Person
    Address             // También tiene un Address
    EmployeeID string
    Salary    float64
}
```

Cuando embebes un struct, todos sus campos y métodos están disponibles directamente:

```go
employee := NewEmployee(...)
fmt.Println(employee.Name)  // Acceso directo a campos de Person
fmt.Println(employee.Street) // Acceso directo a campos de Address
```

Esto es más flexible que la herencia porque puedes cambiar la composición sin romper la jerarquía de clases.

## Métodos y Receptores

En Go, puedes agregar métodos a cualquier tipo (no solo structs). Los métodos usan **receptores**:

```go
// Value receiver - recibe una copia
func (p Person) GetEmail() string {
    return p.email
}

// Pointer receiver - recibe una referencia
func (p *Person) SetEmail(email string) {
    p.email = email
}
```

### ¿Cuándo usar cada uno?

- **Value receiver**: Cuando el método solo lee datos o el struct es pequeño
- **Pointer receiver**: Cuando necesitas modificar el struct o es grande (evita copias costosas)

## Comparación de Structs

A diferencia de Java donde necesitas implementar `equals()`, en Go la comparación es automática si todos los campos son comparables:

```go
p1 := Person{Name: "John", Age: 30}
p2 := Person{Name: "John", Age: 30}
fmt.Println(p1 == p2) // true - comparación automática
```

## Structs Anónimos

Go permite crear structs sin nombre, útiles para configuraciones temporales:

```go
config := struct {
    Host string
    Port int
    SSL  bool
}{
    Host: "localhost",
    Port: 8080,
    SSL:  true,
}
```

Esto es más simple que crear clases anónimas en Java.

## Ejemplo Práctico: Sistema de Pagos

Para cerrar, veamos un ejemplo práctico que prepara el terreno para las interfaces (que veremos en el siguiente post):

```go
type PaymentMethod interface {
    ProcessPayment(amount float64) error
    GetName() string
}

type CreditCard struct {
    Number     string
    ExpiryDate string
    CVV        string
    CardHolder string
}

func (c *CreditCard) ProcessPayment(amount float64) error {
    fmt.Printf("Processing credit card payment of $%.2f\n", amount)
    return nil
}
```

Este ejemplo muestra cómo diferentes tipos pueden implementar el mismo comportamiento sin herencia explícita.

## Conclusiones

Los structs en Go son más simples que las clases en Java, pero no menos poderosos:

✅ **Simplicidad**: Menos código boilerplate
✅ **Composición**: Más flexible que la herencia
✅ **Performance**: Control explícito sobre copias vs referencias
✅ **Claridad**: El código es más directo y fácil de entender

Si vienes de Java, el cambio de mentalidad puede tomar un poco de tiempo, pero una vez que te acostumbres, encontrarás que Go te da más control con menos complejidad.

## Próximos Pasos

En el siguiente post exploraremos las **interfaces implícitas** de Go y cómo implementan el principio de "duck typing". Verás cómo el sistema de pagos que empezamos aquí se completa con múltiples proveedores usando interfaces.

---

**¿Tienes preguntas sobre structs en Go?** Déjame saber en los comentarios. Y si quieres ver el código completo con todos los ejemplos, puedes encontrarlo en mi repositorio [go-mastery-lab](https://github.com/tu-usuario/go-mastery-lab).

