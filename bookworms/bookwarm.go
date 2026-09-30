package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Bookworm 读书爱好者，拥有书籍列表
type Bookworm struct {
	Name  string
	Books []Book
}

// Book 代表一本书
type Book struct {
	Author string
	Title  string
}

// loadBookworms 读取JSON文件返回读书爱好者列表
func loadBookworms(filePath string) ([]Bookworm, error) {
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

// booksCount 统计每本书出现的次数
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)
	for _, bw := range bookworms {
		for _, book := range bw.Books {
			count[book]++ // map不存在key，拿到uint的零值0，再+1
		}
	}
	return count
}

// findCommonBooks 返回多个人共同读过的书籍，排序保证确定性
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

// displayBooks 打印书籍列表
func displayBooks(books []Book) {
	for _, b := range books {
		fmt.Printf("- %s by %s\n", b.Title, b.Author)
	}
}

// set 集合，key为Book，value是空结构体，零内存开销
type set map[Book]struct{}

// booksToSet 将[]Book转为set集合，用于快速判断书籍是否存在
func bookToSet(books []Book) set {
	s := make(set)
	for _, b := range books {
		s[b] = struct{}{}
	}
	return s
}

// recommendBooks 书籍推荐：两个人有共同读过的书，则互相推荐对方读过自己没读过的书
func recommendBooks(bookworms []Bookworm) []Bookworm {
	// 存储每个人对应的藏书集合
	personSet := make(map[string]set)
	for _, bw := range bookworms {
		personSet[bw.Name] = bookToSet(bw.Books)
		fmt.Println(personSet) // 看看数据长什么样
	}

	// 结果切片，全新对象，不影响原数据
	var result []Bookworm

	// 遍历每个人
	for _, personA := range bookworms {
		aSet := personSet[personA.Name]
		var recommend []Book
		// 和其他人对比
		for _, personB := range bookworms {
			if personA.Name == personB.Name {
				continue
			}
			bSet := personSet[personB.Name]

			// 判断2个人是否有读过相同的书
			hasCommon := false
			for book := range aSet {
				if _, ok := bSet[book]; ok {
					hasCommon = true
					break
				}
			}
			if !hasCommon {
				continue // 没有共同书籍，不做推荐
			}

			for _, book := range personB.Books {
				if _, exist := aSet[book]; !exist {
					recommend = append(recommend,book)
				}
			}
		}
		// 去重，同一本书可能被多人推荐
		recSet := bookToSet(recommend)
		var uniqueRec []Book
		for b := range recSet {
			uniqueRec = append(uniqueRec, b)
		}
		// 排序
		sort.Slice(uniqueRec, func(i,j int)bool {
			if uniqueRec[i].Author != uniqueRec[j].Author {
				return uniqueRec[i].Author < uniqueRec[j].Author
			}
			return uniqueRec[i].Title < uniqueRec[j].Title
		})

		result = append(result, Bookworm{
			Name: personA.Name,
			Books: uniqueRec,
		})
	}
	return result
}
