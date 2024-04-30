package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func Read(ctx context.Context, filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func Write(ctx context.Context, filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inputFile := "input.txt"
	outputFile := "output.txt"

	inputData, err := Read(ctx, inputFile)
	if err != nil {
		log.Fatal(err)
	}


	err = Write(ctx, outputFile, inputData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Faylga muvaffaqiyatli yozildi.")
}
