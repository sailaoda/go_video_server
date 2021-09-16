package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) { //一般情况下大小写区分开，作为子test
	res := Printlto20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Printlto20")
	}
}

func testPrint2(t *testing.T) {
	res := Printlto20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", testPrint)
	t.Run("TestPrint2", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins..... ")
	m.Run()
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ { //寻找函数最稳定的时候，得出最稳定的次数和每次跑的时间
		Printlto20()
	}
}
