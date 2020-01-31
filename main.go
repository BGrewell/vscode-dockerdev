package main

import (
	"fmt"
)

// Test is a simple test structure
type Test struct {
	x int
	y string
	z float32
}

// SetValuesViaPointer sets the values of the object via a pointer to the object
func (t *Test) SetValuesViaPointer(x int, y string, z float32) {
	t.x = x
	t.y = y
	t.z = z
	fmt.Printf("[inside]  pointer - x: %-4d  y: %-10s  z: %f\n", t.x, t.y, t.z)
}

// SetValuesViaCopy sets the values of the object via a copy of the object
func (t Test) SetValuesViaCopy(x int, y string, z float32) {
	t.x = x
	t.y = y
	t.z = z
	fmt.Printf("[inside]  copy    - x: %-4d  y: %-10s  z: %f\n", t.x, t.y, t.z)
}

func main() {
	// Create objects
	pointer := Test{0, "test", 0.0}
	copy    := Test{0, "test", 0.0}

	// Change values
	pointer.SetValuesViaPointer(1, "changed", 1.5)
	copy.SetValuesViaCopy(1, "changed", 1.5)

	// Print values
	fmt.Printf("[outside] pointer - x: %-4d  y: %-10s  z: %f\n", pointer.x, pointer.y, pointer.z)
	fmt.Printf("[outside] copy    - x: %-4d  y: %-10s  z: %f\n", copy.x, copy.y, copy.z)
}