// Command line interface for sorting lines of text files.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("sortof: ")

	if err := Sandbox(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	config, err := NewAppConfig(os.Args[1:])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if config.ExitMessage != "" {
		if _, err := fmt.Fprintln(os.Stdout, config.ExitMessage); err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	ctx, cancel := NewAppContext(config)
	defer cancel()

	var files []io.ReadCloser
	for _, name := range config.Files {
		if name == "-" {
			files = append(files, os.Stdin)
			continue
		}

		f, err := os.Open(name)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}()

		files = append(files, f)
	}

	if len(config.Files) == 0 {
		files = []io.ReadCloser{os.Stdin}
	}

	for _, file := range files {
		sorted, err := config.SortFunc(ctx, file)
		if err != nil {
			switch {
			case err == context.Canceled:
				log.Println("sorting cancelled by user")
			case err == context.DeadlineExceeded:
				log.Println("sorting needs more time than expected")
			default:
				log.Println(err)
			}
			os.Exit(1)
		}

		for _, v := range sorted {
			if _, err := fmt.Fprintln(os.Stdout, v); err != nil {
				panic(err)
			}
		}
	}

	os.Exit(0)
}
