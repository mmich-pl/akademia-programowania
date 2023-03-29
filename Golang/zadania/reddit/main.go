package main

import (
	"context"
	"log"
	"os"
	"reddit_fetcher/fetcher"
	"sync"
)

type fetcherTuple struct {
	url, path string
}

func main() {
	outputs := []fetcherTuple{
		{"https://www.reddit.com/r/golang.json?limit=100", "golang.txt"},
		{"https://www.reddit.com/r/minecraft.json?limit=100", "minecraft.txt"},
		{"https://www.reddit.com/r/rust.json?limit=100", "rust.txt"},
	}

	var group sync.WaitGroup
	group.Add(len(outputs))

	for _, o := range outputs {
		go execFetcherAsync(o.url, o.path, &group)
	}
	group.Wait()
	log.Println("Task ended")
}

func execFetcher(url string, path string, group *sync.WaitGroup) {
	defer group.Done()

	f := fetcher.Fetcher{Address: url}

	if err := f.Fetch(context.Background()); err != nil {
		log.Println(err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		log.Printf("cannot open file '%s': %v\n", path, err)
	}

	if err := f.Save(file); err != nil {
		log.Printf("cannot save data: %v\n", err)
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("cannot close file '%s': %v\n", path, err)
		}
	}(file)

	log.Printf("Data saved to file %s", path)
}

func execFetcherAsync(url string, path string, group *sync.WaitGroup) {
	if err := os.Remove(path); err != nil {
		log.Printf("cannot remove file %s", path)
	}

	defer group.Done()
	ctx := context.Background()

	f := fetcher.Fetcher{Address: url}

	if err := f.Fetch(ctx); err != nil {
		log.Println(err)
	}

	if err := f.SaveAsync(ctx, path, 3); err != nil {
		log.Printf("cannot save data: %v\n", err)
	}

	log.Printf("Data saved to file %s", path)
}
