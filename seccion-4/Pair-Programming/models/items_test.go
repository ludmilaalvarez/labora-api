package models

import "testing"

type addTest struct {
    arg1 float64
	arg2 int
	expected float64
}

var addTests = []addTest{
    addTest{10.0, 2, 20},
    addTest{1, 30, 30},
    addTest{3, 30, 90},
    addTest{4, 23, 92},
    
}


func TestPrecioTotal(t *testing.T){

    for _, test := range addTests{
		var item = Item{0, "", "", "", test.arg2, test.arg1, "", 0.0}
		item.PrecioTotal()
        if item.TotalPrice != test.expected {
            t.Errorf("Output %f not equal to expected %f", item.TotalPrice, test.expected)
        }
    }
} 


/* func TestPrecioTotal(t *testing.T){

    for _, test := range addTests{
		var item = Item{0, "", "", "", test.arg2, test.arg1, "", 0.0}
		item.TotalPrice= PrecioTotal(item)
        if item.TotalPrice != test.expected {
            t.Errorf("Output %f not equal to expected %f", item.TotalPrice, test.expected)
        }
    }
}  */