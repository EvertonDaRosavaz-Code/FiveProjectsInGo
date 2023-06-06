package main

import (
	"fmt"
)


var operador string;
var num1, num2 int 

func main() {
	// calculadora()
	conversor();
}

//Calculadora
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

var moeda string
var moedaValor int
var newMoeda string

// Conversor de moedas
func conversor() {
	fmt.Println("Conversor de moeda simples")
	fmt.Println("USD | EUR | BRL")

	fmt.Print("Qual a sua moeda nacional: ")
	fmt.Scan(&moeda)
	fmt.Println("Sua moeda atual é", moeda)

	fmt.Print("Quanto você deseja converter: ")
	fmt.Scan(&moedaValor)

	fmt.Print("Para qual moeda deseja converter? ")
	fmt.Scan(&newMoeda)

	// USD
	if moeda == "USD" && newMoeda == "EUR" {
		taxa := 0.93
		converter := float64(moedaValor) * taxa
		fmt.Println("USD para EUR fica", converter)
	} else if moeda == "USD" && newMoeda == "BRL" {
		taxa := 4.91
		converter := float64(moedaValor) * taxa
		fmt.Println("USD para BRL fica", converter)
	}

	// EUR
	if moeda == "EUR" && newMoeda == "USD" {
		taxa := 1.07
		converter := float64(moedaValor) * taxa
		fmt.Println("EUR para USD fica", converter)
	} else if moeda == "EUR" && newMoeda == "BRL" {
		taxa := 5.25
		converter := float64(moedaValor) * taxa
		fmt.Println("EUR para BRL fica", converter)
	}

	// BRL
	if moeda == "BRL" && newMoeda == "EUR" {
		taxa := 0.19
		converter := float64(moedaValor) * taxa
		fmt.Println("BRL para EUR fica", converter)
	} else if moeda == "BRL" && newMoeda == "USD" {
		taxa := 0.20
		converter := float64(moedaValor) * taxa
		fmt.Println("BRL para USD fica", converter)
	} else {
		fmt.Println("Erro")
	}
}


