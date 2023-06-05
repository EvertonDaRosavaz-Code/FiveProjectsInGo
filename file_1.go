package main
import "fmt"

var operador string;
var num1, num2 int 

func main() {
	calculadora()
}

func calculadora(){
	fmt.Println("Calculadora Símples");

	fmt.Print("Digite o operador desejado:");
	fmt.Scan(&operador);
	
	switch operador {
		case "+":
			fmt.Println("Digite os valor depois pressione ENTER")
			fmt.Scan(&num1, &num2);
			soma(num1, num2);
		case "-":
			fmt.Println("Digite os valor depois pressione ENTER")
			fmt.Scan(&num1, &num2);
			subtracao(num1, num2);
		case "*":
			fmt.Println("Digite os valor depois pressione ENTER")
			fmt.Scan(&num1, &num2);
			multipli(num1, num2);
		case "/":
			fmt.Println("Digite os valor depois pressione ENTER")
			fmt.Scan(&num1, &num2);
			dividir(num1, num2);
	}
}



func soma(num1 int, num2 int){
	result := num1 + num2;
	fmt.Print("Resultado da soma ", result);
}

func subtracao(num1 int, num2 int){
	result := num1 - num2;
	fmt.Print("Resultado da subtração ", result);
}

func multipli(num1 int, num2 int){
	result := num1 * num2;
	fmt.Print("Resultado da multiplicação ", result);
}

func dividir(num1 int , num2 int){
	if(num2 == 0){
		fmt.Print("Nao é possivel dividir por zero")
		fmt.Println("")
		operador = ""
		calculadora()
	}else{
		result := num1/num2;
		fmt.Print("Resultado da divisão ", result);
	}
}






