package main

import (
	"fmt"
	"sort"
)

// ============================================================================
// SLICES, MAPS, ARRAYS (CAP, LEN, MAKE)
// ============================================================================
// En Go hay diferencias importantes entre arrays y slices
// - Arrays: tamaño fijo, se pasan por valor (se copian)
// - Slices: tamaño dinámico, son referencias a arrays subyacentes
// - Maps: estructuras clave-valor (similar a HashMap en Java)
// ============================================================================

// ============================================================================
// ARRAYS
// ============================================================================
// Arrays tienen tamaño fijo y se pasan por valor (se copian completamente)
// En Java, los arrays son objetos y se pasan por referencia
// En Go, los arrays son valores primitivos

func demonstrateArrays() {
	// Array de tamaño fijo
	var arr1 [5]int                    // [0 0 0 0 0]
	arr2 := [5]int{1, 2, 3, 4, 5}      // [1 2 3 4 5]
	arr3 := [...]int{1, 2, 3}          // El compilador cuenta: [1 2 3]

	fmt.Printf("arr1: %v, len: %d\n", arr1, len(arr1))
	fmt.Printf("arr2: %v, len: %d\n", arr2, len(arr2))
	fmt.Printf("arr3: %v, len: %d\n", arr3, len(arr3))

	// Los arrays se copian por valor
	arr4 := arr2
	arr4[0] = 999
	fmt.Printf("arr2 después de modificar arr4: %v\n", arr2) // No cambió
	fmt.Printf("arr4: %v\n", arr4)                            // Sí cambió
}

// ============================================================================
// SLICES
// ============================================================================
// Slices son referencias a arrays subyacentes
// Son más flexibles que arrays y se usan mucho más
// Similar a ArrayList en Java, pero más eficiente

func demonstrateSlices() {
	// Crear slices de diferentes formas
	var s1 []int                    // nil slice
	s2 := []int{}                   // slice vacío (no nil)
	s3 := []int{1, 2, 3, 4, 5}     // slice con valores iniciales
	s4 := make([]int, 5)            // slice de longitud 5, capacidad 5
	s5 := make([]int, 0, 10)        // slice de longitud 0, capacidad 10

	fmt.Printf("s1 (nil): %v, len: %d, cap: %d, is nil: %v\n", s1, len(s1), cap(s1), s1 == nil)
	fmt.Printf("s2 (empty): %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5: %v, len: %d, cap: %d\n", s5, len(s5), cap(s5))

	// LEN y CAP
	// len: número de elementos actuales
	// cap: capacidad del array subyacente
	// Cuando append excede la capacidad, Go crea un nuevo array con más capacidad

	// Append - agrega elementos al final
	s6 := []int{1, 2, 3}
	fmt.Printf("Before append: %v, len: %d, cap: %d\n", s6, len(s6), cap(s6))
	s6 = append(s6, 4, 5, 6)
	fmt.Printf("After append: %v, len: %d, cap: %d\n", s6, len(s6), cap(s6))

	// Slicing - crear sub-slices
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := original[2:5]        // [2 3 4] - índices 2 a 4 (no incluye 5)
	slice2 := original[:5]          // [0 1 2 3 4] - desde inicio hasta índice 4
	slice3 := original[5:]          // [5 6 7 8 9] - desde índice 5 hasta el final
	slice4 := original[:]           // Copia de referencia (no copia los datos)

	fmt.Printf("original: %v\n", original)
	fmt.Printf("slice1 [2:5]: %v\n", slice1)
	fmt.Printf("slice2 [:5]: %v\n", slice2)
	fmt.Printf("slice3 [5:]: %v\n", slice3)

	// IMPORTANTE: Los slices comparten el array subyacente
	slice1[0] = 999
	fmt.Printf("After modifying slice1[0]: original=%v\n", original) // ¡Cambió!
	fmt.Printf("slice1=%v\n", slice1)

	// Copiar slices (crear copia independiente)
	copySlice := make([]int, len(original))
	copy(copySlice, original)
	copySlice[0] = 111
	fmt.Printf("original after modifying copy: %v\n", original) // No cambió
	fmt.Printf("copySlice: %v\n", copySlice)
}

// ============================================================================
// MAPS
// ============================================================================
// Maps son estructuras clave-valor (similar a HashMap en Java)
// Las claves deben ser comparables (no slices, maps, ni funciones)

func demonstrateMaps() {
	// Crear maps de diferentes formas
	var m1 map[string]int           // nil map (no se puede usar hasta inicializar)
	m2 := make(map[string]int)      // map vacío inicializado
	m3 := map[string]int{           // map con valores iniciales
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}

	// Inicializar m1 antes de usar
	m1 = make(map[string]int)
	m1["one"] = 1
	m1["two"] = 2

	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m2: %v\n", m2)
	fmt.Printf("m3: %v\n", m3)

	// Acceder a valores
	value := m3["apple"]
	fmt.Printf("m3['apple']: %d\n", value)

	// Verificar si una clave existe
	value, exists := m3["grape"]
	if exists {
		fmt.Printf("grape exists: %d\n", value)
	} else {
		fmt.Println("grape does not exist")
	}

	// Usar el valor cero cuando la clave no existe
	value = m3["grape"] // value será 0 (valor cero de int)
	fmt.Printf("m3['grape'] (no existe): %d\n", value)

	// Eliminar elementos
	delete(m3, "banana")
	fmt.Printf("After deleting 'banana': %v\n", m3)

	// Iterar sobre maps
	fmt.Println("Iterating over m3:")
	for key, value := range m3 {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// Obtener solo claves o solo valores
	keys := make([]string, 0, len(m3))
	values := make([]int, 0, len(m3))
	for k, v := range m3 {
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Printf("Keys: %v\n", keys)
	fmt.Printf("Values: %v\n", values)
}

// ============================================================================
// CAP Y LEN EN DETALLE
// ============================================================================

func demonstrateCapAndLen() {
	// Para arrays: len y cap son siempre iguales
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array - len: %d, cap: %d\n", len(arr), cap(arr))

	// Para slices: len puede ser menor que cap
	slice := make([]int, 3, 10) // len=3, cap=10
	fmt.Printf("Slice inicial - len: %d, cap: %d\n", len(slice), cap(slice))

	// Cuando append excede la capacidad, Go duplica la capacidad (generalmente)
	for i := 0; i < 20; i++ {
		oldCap := cap(slice)
		slice = append(slice, i)
		if cap(slice) != oldCap {
			fmt.Printf("Capacidad creció de %d a %d (len: %d)\n", oldCap, cap(slice), len(slice))
		}
	}

	// Para maps: solo len (no tienen cap)
	m := make(map[string]int, 10) // capacidad inicial sugerida (no garantizada)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("Map - len: %d\n", len(m))
}

// ============================================================================
// EJEMPLO PRÁCTICO: SISTEMA DE INVENTARIO
// ============================================================================

type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

type Inventory struct {
	products map[string]*Product
	orders   []string // IDs de productos ordenados
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]*Product),
		orders:   make([]string, 0),
	}
}

func (inv *Inventory) AddProduct(id, name string, price float64, stock int) {
	inv.products[id] = &Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	inv.orders = append(inv.orders, id)
}

func (inv *Inventory) GetProduct(id string) (*Product, bool) {
	product, exists := inv.products[id]
	return product, exists
}

func (inv *Inventory) UpdateStock(id string, quantity int) error {
	product, exists := inv.products[id]
	if !exists {
		return fmt.Errorf("product %s not found", id)
	}
	if product.Stock+quantity < 0 {
		return fmt.Errorf("insufficient stock for product %s", id)
	}
	product.Stock += quantity
	return nil
}

func (inv *Inventory) ListProducts() []*Product {
	products := make([]*Product, 0, len(inv.products))
	for _, id := range inv.orders {
		if product, exists := inv.products[id]; exists {
			products = append(products, product)
		}
	}
	return products
}

func (inv *Inventory) GetTotalValue() float64 {
	total := 0.0
	for _, product := range inv.products {
		total += product.Price * float64(product.Stock)
	}
	return total
}

// ============================================================================
// OPERACIONES AVANZADAS CON SLICES
// ============================================================================

func advancedSliceOperations() {
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	// Ordenar
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Sorted: %v\n", sorted)

	// Filtrar (elementos mayores que 5)
	filtered := make([]int, 0)
	for _, n := range numbers {
		if n > 5 {
			filtered = append(filtered, n)
		}
	}
	fmt.Printf("Filtered (>5): %v\n", filtered)

	// Mapear (duplicar valores)
	mapped := make([]int, len(numbers))
	for i, n := range numbers {
		mapped[i] = n * 2
	}
	fmt.Printf("Mapped (x2): %v\n", mapped)

	// Reducir (suma)
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Printf("Sum: %d\n", sum)
}

// ============================================================================
// MAIN - Ejemplos de uso
// ============================================================================

func main() {
	fmt.Println("=== FUNDAMENTOS: SLICES, MAPS, ARRAYS ===\n")

	// 1. Arrays
	fmt.Println("1. Arrays:")
	demonstrateArrays()
	fmt.Println()

	// 2. Slices
	fmt.Println("2. Slices:")
	demonstrateSlices()
	fmt.Println()

	// 3. Maps
	fmt.Println("3. Maps:")
	demonstrateMaps()
	fmt.Println()

	// 4. Cap y Len
	fmt.Println("4. Cap y Len:")
	demonstrateCapAndLen()
	fmt.Println()

	// 5. Ejemplo práctico: Inventario
	fmt.Println("5. Ejemplo práctico: Sistema de Inventario:")
	inventory := NewInventory()
	inventory.AddProduct("P001", "Laptop", 999.99, 10)
	inventory.AddProduct("P002", "Mouse", 29.99, 50)
	inventory.AddProduct("P003", "Keyboard", 79.99, 30)

	fmt.Println("Productos en inventario:")
	for _, product := range inventory.ListProducts() {
		fmt.Printf("  %s: %s - $%.2f (Stock: %d)\n", product.ID, product.Name, product.Price, product.Stock)
	}

	fmt.Printf("\nValor total del inventario: $%.2f\n", inventory.GetTotalValue())

	if err := inventory.UpdateStock("P001", -2); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		product, _ := inventory.GetProduct("P001")
		fmt.Printf("Stock actualizado de P001: %d\n", product.Stock)
	}
	fmt.Println()

	// 6. Operaciones avanzadas con slices
	fmt.Println("6. Operaciones avanzadas con slices:")
	advancedSliceOperations()
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

