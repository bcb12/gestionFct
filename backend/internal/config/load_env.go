package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnv(envFile string) {
	file, err := os.Open(envFile)
	if err != nil {
		log.Fatalf("Error opening the file %s: %v", envFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Line %d ignored (the format is not valid: KEY=VAL): %s", i, line)
			i++
			continue
		}
		key := parts[0]
		val := parts[1]
		os.Setenv(key, val)
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the file %s: %v", envFile, err)
	}
}
