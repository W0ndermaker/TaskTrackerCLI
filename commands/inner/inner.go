package inner

import (
	"fmt"
	"os"
)

func ReadData(fileName string) ([]byte, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error occured while reading storage file")
		return nil, err
	}
	return f, nil
}
