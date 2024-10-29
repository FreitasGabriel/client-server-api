package files

import (
	"log"
	"os"
)

func WriteFile(content string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println("error to create file", err)
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Println("error to write file", err)
		return err
	}

	file.Close()
	return nil
}
