/*
package main

import (
	"fmt"
	"testing"
)

func TestPrintlto20(t *testing.T) {
	res := Printlto20()
	fmt.Println("hey")
	if res != 210{
		t.Error("Wrong result of Printlto20")
	}
}
*/
package main

import (
	"fmt"
	"testing"
)

func TestPrintlto20(t *testing.T) {
	t.SkipNow()
	res := Printlto20()
	fmt.Println("hey")
	if res != 200 {
		t.Error("Wrong result of Printlto20")
	}
}

/*
package main
import(
"testing"
"fmt"
)

func TestPrint(t *testing.T) {
	t.Run("a1",func(t *testing.T) {fmt.Println("a1")})
	t.Run("a2",func(t *testing.T) {fmt.Println("a2")})
	t.Run("a3",func(t *testing.T) {fmt.Println("a3")})
}

*/
