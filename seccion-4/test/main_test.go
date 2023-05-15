package prueba

import ("testing"
        "fmt")


// arg1 significa el argumento 1 and arg2 el argumento 2, and the expected stands for the 'el resultado que esperamos'
type testGeneral struct {
    arg1, arg2, expected int
}

type testFactorial struct{
    numero, resultado int
}



var addTests = []testGeneral{
    testGeneral{2, 3, 5},
    testGeneral{4, 8, 12},
    testGeneral{6, 9, 15},
    testGeneral{3, 10, 13},
    
}

var subTest =[]testGeneral{
    testGeneral{10, 5, 5},
    testGeneral{8, 7, 1},
    testGeneral{23, 20, 3},
    testGeneral{3, 3, 0},    
} 

var numeros=[]testFactorial{
    testFactorial{5, 120},
    testFactorial{1, 1},
    testFactorial{0, 1},
    testFactorial{3, 6},
}


func TestAdd(t *testing.T){

    for _, test := range addTests{
        if output := Add(test.arg1, test.arg2); output != test.expected {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
    }
}

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(4, 6)
    }
}

func ExampleAdd() {
    fmt.Println(Add(4, 6))
    // Output: 10
}

func TestSubtract(t *testing.T){
    for _ ,test:= range subTest{
        if output:= Subtract(test.arg1, test.arg2); output != test.expected{
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
    }
} 

func BenchmarkSub(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Subtract(6, 4)
    }
}

func ExampleSub() {
    fmt.Println(Subtract(6, 4))
    // Output: 2
}

func TestFactorial(t *testing.T){
    for _, test:= range numeros{
        if valor:= Factorial(test.numero); valor!= test.resultado{
            t.Errorf("Valor %q no es el resultado esperado %q", valor, test.resultado)
        }
    }
}


func BenchmarkFactorial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Factorial(4)
    }
}

func ExampleFactorial() {
    fmt.Println(Factorial(4))
    // Output: 24
}