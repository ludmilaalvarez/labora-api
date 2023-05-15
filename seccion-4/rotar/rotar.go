package main 

import "fmt"

func main (){
	var palabra string

	fmt.Println("Ingrese la palabra que desea rotar: ")
	fmt.Scan(&palabra)
	
	
	resultado:=version1(palabra)
	resultado2:=version2(palabra)
	resultado3:=version3(palabra)
	resultado4:=version4(palabra)
	resultado5:=version5(palabra)


	
	fmt.Println("Resultado 1:",resultado) 
	fmt.Println("Resultado 2:",resultado2)
	fmt.Println("Resultado 3:",resultado3)
	fmt.Println("Resultado 4:",resultado4)
	fmt.Println("Resultado 5:",resultado5)

 
}


func version1(palabra string) string{
	nuevo:=string(palabra[len(palabra)-1])+palabra[0:(len(palabra)-1)]
	return nuevo
}



func version2(palabra string) string{
	var nuevo string
	valor:=len(palabra)
	for i:=0; i<valor; i++{
		formula1:=valor+(valor+(i-(valor+1)))
		formula2:=(valor+(i-(valor+1)))
		if i==0{
			nuevo=string(palabra[formula1])
		}else{
			nuevo=nuevo+string(palabra[formula2])
		}
	}
	return nuevo
}


func version3(palabra string) string{
	nuevapalabra:=palabra+palabra
	longitud:=len(palabra)
	resultado:=nuevapalabra[(longitud-1):(longitud*2-1)]
	
	return resultado
}

func version4(palabra string) string{
	nuevapalabra:=palabra+palabra
	longitud:=len(palabra)
	resultado:=nuevapalabra[1:(longitud+1)]
	
	return resultado
}

func version5(palabra string) string {
	var b string
	contador:=0
	for _, letra := range palabra {
		if contador==0{
			b = string(letra)
		}
		contador=1
	}
	return palabra[1:len(palabra)]+b
}