package main // Déclaration du package

import "fmt" // Importations

// Déclaration de variables globales, constantes, etc. (si nécessaire)
const PI = 3.14 // une variable constant qui ne change jamais

const (
	A int = 1
	B     = 3
	C     = "malik"
) // multiple constant

func myMessage() { // fonction qui retourne un string
	fmt.Println("malilk est un dev")
}
func myName(fname string) {
	fmt.Println("my name is", fname)

}

type person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() { // Fonction principale

	fmt.Println("Hello, World!") // ceci est un commaintaire

	var malik string = "malik" // var variablename type = value

	var malik2 string
	malik2 = "malik2"

	fmt.Println(malik)
	fmt.Println(malik2)

	// ICI ON TRAITERA LES TYPES DE DONNEES EN GO
	x := 20
	y := 30
	var z = x + y
	fmt.Println(z)

	fmt.Println(A)

	var i string = "Hello"
	var j int = 10

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)

	// ICI ON TRAITERA LES TABLEAUX EN GO

	var arr1 = [4]int{1, 3, 4, 5}
	arr2 := [3]int{4, 5, 3}
	fmt.Println(arr2)
	fmt.Println(arr1)
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars[0])
	cars[2] = "merc"
	fmt.Println(cars)

	// ICI ON TRAITERA LES SLICES EN GO
	var slice1 = []int{1, 2, 3, 4, 5} // slice est un tableau dynamique en go
	fmt.Println(slice1)
	fmt.Println(slice1[1:3]) // affiche les elements de l'index 1 à 3

	// ICI ON TRAITERA CONDITIONS EN GO
	var age int = 10
	if age > 18 {
		fmt.Println("tu passe")
	} else if age == 10 {
		fmt.Println("yes")
	} else {
		fmt.Println("nop")
	}

	// ici switchcase *
	var day int = 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Sunday")
	}

	// ICI ON TRAITERA LES BOUCLES EN GO
	// for loop
	for i := 0; i < 5; i++ { // tant que i est inferieur à 5 on incremente i de 1 à chaque tour de boucle
		fmt.Println(i)
	}
	for i := 0; i < 6; i++ { // tant que i est inferieur à 5 on incremente i de 1 à chaque tour de boucle
		fmt.Println("yes") // affiche yes
	}

	myMessage() // appel de la fonction myMessage
	myName("malik")

	var pers1 person

	pers1.name = "malik"
	pers1.age = 20
	pers1.job = "dev"
	pers1.salary = 1000

	fmt.Println("name ", pers1.name)

	// les map en go
	var myMap = map[string]int{"malik": 20, "mohamed": 30}
	fmt.Println(myMap)

	// Déclaration
	var myMap = map[string]int{}
	// Ajout de valeurs
	myMap["one"] = 1
	myMap["two"] = 2

	// Lecture et affichage des valeurs
	fmt.Println("Value of key 'one':", myMap["one"])

	// Suppression d'une valeur
	delete(myMap, "one")

	// Vérification de l'existence d'une clé
	value, exists := myMap["two"]
	if exists {
		fmt.Println("Value of key 'two':", value)
	} else {
		fmt.Println("Key 'two' does not exist.")
	}
}
