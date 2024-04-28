package main

import "fmt"

func main() {
	fmt.Println("Maps in go!!!!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["JA"] = "Java"
	languages["PR"] = "Perl"

	// fmt.Println("List of all the languages: ", languages)
	fmt.Println("JS means: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all the languages: ", languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
