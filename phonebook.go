//In The Name Of God

package main

import "fmt"

//---------- Functions -------------
func menu() int{
	fmt.Println("1. Add phone number")
	fmt.Println("2. Delete phone number")
	fmt.Println("3. Search phone number")
	fmt.Println("4. Change phone number")
	fmt.Println("5. Display all")
	fmt.Println("6. Exit")

	var choice int
	fmt.Print("Enter 1 to 6: ")
	fmt.Scan(&choice)

	return choice
}

func addNumber(My_phonebook map[string]string) map[string]string{
	var name, phone string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter phone number: ")
	fmt.Scan(&phone)

	_, exists := My_phonebook[name]

	if exists == false {
		My_phonebook[name] = phone
		fmt.Println(name, "added!")
	} else {
		fmt.Println(name, "is in phonebook")
	}

	return My_phonebook
}

func deleteNumber(My_phonebook map[string]string) map[string]string{
	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	_, exists := My_phonebook[name]
	
	if exists {
		delete(My_phonebook, name)
		fmt.Println(name, "deleted!")
	} else {
		fmt.Println(name, "is not in phonebook")
	}

	return My_phonebook
}

func searchNumber(My_phonebook map[string]string){
	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	_, exists := My_phonebook[name]
	if exists {
		fmt.Println("Phone number: ", My_phonebook[name])
	} else {
		fmt.Println(name, "is not in phonebook")
	}
}

func changeNumber(My_phonebook map[string]string) map[string]string{
	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	var new_phone string
	fmt.Print("Enter new phone number: ")
	fmt.Scan(&new_phone)

	_, exists := My_phonebook[name]

	if exists {
		My_phonebook[name] = new_phone
	} else {
		fmt.Println(name, "is not in phonebook!")
	}

	return My_phonebook
}

func displayAll(My_phonebook map[string]string){
	for key, value := range My_phonebook{
		fmt.Printf("name: %s - phone: %s\n", key, value)
	}
}

//---------- The Main -------------
func main(){
	My_phonebook := make(map[string]string)
	for{
		choice := menu()
		if choice == 6{
			break
		}
		if choice == 1{
			My_phonebook = addNumber(My_phonebook)
		}
		if choice == 2{
			My_phonebook = deleteNumber(My_phonebook)
		}
		if choice == 3{
			searchNumber(My_phonebook)
		}
		if choice == 4{
			My_phonebook = changeNumber(My_phonebook)
		}
		if choice == 5{
			displayAll(My_phonebook)
		}
	}
}

