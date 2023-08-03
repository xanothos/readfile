package readfile

import (
	"fmt"
	"os"
	"bufio"
)

func Read(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Fatal: Can't read %s: %s", filename, err)
	}
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}