package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	wp "github.com/takanoriyanagitani/go-wikibzip2pages"
)

func FilenameToPages(
	filename string,
	offset int64,
	size int64,
) ([]wp.BasicPage, error) {
	f, e := os.Open(filename)
	if nil != e {
		return nil, e
	}
	pages, e := wp.FileToPages(f, offset, size)
	return pages, errors.Join(e, f.Close())
}

func env2filename() string {
	return os.Getenv("WIKIPAGE_FILE")
}

func env2offset() (int64, error) {
	s := os.Getenv("WIKIPAGE_OFFSET")
	if s == "" {
		return 0, nil // unset -> read from the start of the file
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("WIKIPAGE_OFFSET=%q: %w", s, err)
	}

	return v, nil
}

func env2size() (int64, error) {
	s := os.Getenv("WIKIPAGE_SIZE")
	if s == "" {
		return 0, nil // unset -> read all
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("WIKIPAGE_SIZE=%q: %w", s, err)
	}

	return v, nil
}

func sub() error {
	filename := env2filename()
	if filename == "" {
		return errors.New("WIKIPAGE_FILE is not set")
	}

	offset, err := env2offset()
	if err != nil {
		return err
	}
	size, err := env2size()
	if err != nil {
		return err
	}

	pages, err := FilenameToPages(filename, offset, size)
	if err != nil {
		return err
	}

	for _, p := range pages {
		fmt.Printf("%s\n", p.ShortString())
	}

	return nil
}

func main() {
	if err := sub(); err != nil {
		log.Printf("%v\n", err)
	}
}
