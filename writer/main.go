package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	size := flag.Int("size", 100, "nombre d'etablissement")

	file, err := os.Create("../data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	for i := 0; i < *size; i++ {
		w.WriteString(fmt.Sprintf("%07dETABLISSEMENT%07d\n", i, i))
	}

	w.Flush()

}
