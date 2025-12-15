package main

import "fmt"

func main() {
	reader := make(map[string]map[string][]string)
	reader["Алиса"] = make(map[string][]string)
	reader["Алиса"]["Книги"] = []string{"Война и мир", "Мир и война"}
	reader["Алиса"]["Журналы"] = []string{"Белый мальчик"}

	reader["Иван"] = make(map[string][]string)
	reader["Иван"]["Книги"] = []string{"Алгоритмы"}
	reader["Иван"]["Журналы"] = []string{"Белый мальчик", "Азбука"}

	fmt.Println("Читателей с изданием", len(reader))

	for readerName, public := range reader {
		total := 0
		for _, list := range public {
			total += len(list)
		}
		fmt.Printf("У %s на руках %d Изданий \n", readerName, total)
	}
}
