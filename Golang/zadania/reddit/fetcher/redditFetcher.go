package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch(ctx context.Context) error
	Save(io.Writer) error
}

type Fetcher struct {
	Address        string
	response       response
	mutex          sync.Mutex
	postCounter    int
	responseLength int
}

func (f *Fetcher) SaveAsync(ctx context.Context, path string, limit int) (err error) {
	f.responseLength = len(f.response.Data.Children)
	f.postCounter = 0

	group, ctx := errgroup.WithContext(ctx)

	for j := 0; j < limit; j++ {
		group.Go(func() error {
			file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				return fmt.Errorf("cannot save data to file %s: %v", path, err)
			}

			defer func(file *os.File) {
				err = file.Close()
			}(file)

			for {
				f.mutex.Lock()
				if f.postCounter < f.responseLength-1 {
					f.postCounter++
				} else {
					f.mutex.Unlock()
					break
				}
				f.mutex.Unlock()
				post := f.response.Data.Children[f.postCounter]
				line := fmt.Sprintf("%s\n%s\n", post.Data.Title, post.Data.URL)
				if _, err := file.Write([]byte(line)); err != nil {
					log.Printf("cannot write to file: %v", err)
				}
			}
			return nil
		})
	}

	return group.Wait()
}

func (f *Fetcher) Save(file io.Writer) error {
	for _, post := range f.response.Data.Children {
		line := fmt.Sprintf("%s\n%s\n", post.Data.Title, post.Data.URL)
		if _, err := file.Write([]byte(line)); err != nil {
			return fmt.Errorf("cannot write to file: %v", err)
		}
	}
	return nil
}

func (f *Fetcher) Fetch(ctx context.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.Address, nil)
	if err != nil {
		return fmt.Errorf("cannot create request: %v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot fetch data: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code other than 200: %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&f.response); err != nil {
		return fmt.Errorf("error while decoding data: %v", err)
	}

	return nil
}
