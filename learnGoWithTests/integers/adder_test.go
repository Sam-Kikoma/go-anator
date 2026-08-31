package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){
	sum := Add(2,2)
	expected := 4

	if sum != expected{
		t.Errorf("expected '%d' but got '%d'",expected,sum)
	}
}

// Testable example
// Example functions are compiled whenever tests are executed.
// Example functions begin with Example (much like test functions begin with Test)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}