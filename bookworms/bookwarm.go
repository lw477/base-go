package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string
	Books []Book
}

type Book struct {
	Author string
	Title  string
}

func LoadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)
	for _, bw := range bookworms {
		for _, book := range bw.Books {
			count[book]++ // map不存在key，拿到uint的零值0，再+1
		}
	}
	return count
}

func findCommonBooks(bookworms []Bookworm) []Book {
	countMap := booksCount(bookworms)
	var common []Book

	for book, cnt := range countMap {
		if cnt > 1 {
			common = append(common, book)
		}
	}

	sort.Slice(common, func(i, j int) bool {
		// 先比较作者 Author
		if common[i].Author != common[j].Author {
			// 作者不一样：按作者名字字典序升序
			return common[i].Author < common[j].Author
		}
		// 作者相同：再按书名 Title 字典序升序
		return common[i].Title < common[j].Title
	})
	return common
}

func displayBooks(books []Book) {
	for _, b := range books {
		fmt.Printf("- %s by %s\n", b.Title, b.Author)
	}
}
