package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


// "Поисковая система" поддерживает всего 2 операции
// 1. Добавление документа. Не зависит от количества документов
//    но зависит от количества слов в новом документе 
//    Алгоритмическая сложность добавления документа O(n), 
//    где n - количество слов в документе
//	  Память необходимая для хранения документа, пропорциональна
//    размеру документа, соответственно сложность по памяти O(n)
//
// 2. Поиск документов. 
//    Зависит от количества уникальных слов в запросе n.
//    и от количества документов в которых эти слова встечаются
//    Соответственно алгоритмическая сложность поиска O(n*k)
//    После того как документы нашлись нужно выбрать 5
//    самых релевантных для них используется сортировка выбором
//    т.е. 5 раз среди найденных документов ищется максимум
//    Не используется быстрая сортирвка т.к. нужно найти всего 5 
//    элементов в итоге и велика вероятность, что в итоговом результате
//    будет много одинаковых по релевантности документов и 
//    из-за этого быстрая сортировка будет работать за O(n^2)
//    
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Split(bufio.ScanLines)

	storage := NewStorage()

	sc.Scan()
	DocumentsCount, _ := strconv.Atoi(sc.Text())
	for i := 0; i < DocumentsCount; i++ {
		sc.Scan()
		storage.Add(sc.Text())
	}

	sc.Scan()
	QueryCount, _ := strconv.Atoi(sc.Text())

	report := strings.Builder{}
	for i := 0; i < QueryCount; i++ {
		sc.Scan()
		res := storage.Query(sc.Text())
		formatResult(res, &report)
	}
	fmt.Println(report.String())
}

func formatResult(res []int, sb *strings.Builder) {
	if len(res) < 1 {
		sb.WriteString("\n")
		return
	}

	sb.WriteString(fmt.Sprintf("%d", res[0]))
	for i := 1; i < len(res); i++ {
		sb.WriteString(fmt.Sprintf(" %d", res[i]))
	}

	sb.WriteString("\n")
}

// Структура которая хранит
// для слова в каких документах оно встречалось
// и сколько раз
// [ "Word" =>  ["DocumentID"=>Count, ... ] ]
type WordInDocuments map[string]map[int]int

type Storage struct {
	LastDocumentID int
	Index          WordInDocuments
	QueryUWords    func([]string) []int
	Limit          int
}

type FoundDocument struct {
	Id        int
	WordCount int
}

func (this *Storage) Add(txt string) {
	this.LastDocumentID++
	words := strings.Split(txt, " ")
	for _, w := range words {
		if _, ok := this.Index[w]; !ok {
			this.Index[w] = map[int]int{this.LastDocumentID: 1}
		} else {
			this.Index[w][this.LastDocumentID]++
		}
	}
}

func (this *Storage) Query(txt string) []int {
	uWords := uniqWords(txt)
	return this.QueryUWords(uWords)
}

var uniqWords = func(txt string) []string {
	words := strings.Split(txt, " ")
	uWords := make(map[string]struct{})
	for _, w := range words {
		uWords[w] = struct{}{}
	}
	res := make([]string, len(uWords))
	i := 0
	for w := range uWords {
		res[i] = w
		i++
	}
	return res
}

func NewStorage() *Storage {
	s := &Storage{
		Index: make(WordInDocuments),
		Limit: 5,
	}

	findDocuments := func(uWords []string) []FoundDocument {
		documents := make(map[int]int)
		for _, w := range uWords {
			wDocs := s.Index[w]
			for docId, wCount := range wDocs {
				if _, ok := documents[docId]; !ok {
					documents[docId] = wCount
				} else {
					documents[docId] += wCount
				}
			}
		}

		i := 0
		result := make([]FoundDocument, len(documents))
		for k, v := range documents {
			result[i].Id = k
			result[i].WordCount = v
			i++
		}
		return result
	}

	s.QueryUWords = func(uWords []string) []int {
		foundDocuments := findDocuments(uWords)
		sortResult(foundDocuments, s.Limit)
		return transformResult(foundDocuments, s.Limit)

	}

	return s
}

func sortResult(result []FoundDocument, limit int) {
	less := func(i, j int) bool {
		if result[i].WordCount == result[j].WordCount {
			return result[i].Id < result[j].Id
		}

		return result[i].WordCount > result[j].WordCount
	}

	if len(result) < limit {
		limit = len(result)
	}

	for i := 0; i < limit; i++ {
		maxIdx := i
		for j := i; j < len(result); j++ {
			if less(j, maxIdx) {
				maxIdx = j
			}
		}
		swap(result, maxIdx, i)
	}
}

func transformResult(result []FoundDocument, limit int) []int {
	if len(result) < limit {
		limit = len(result)
	}

	res := make([]int, limit)
	for i := 0; i < len(res); i++ {
		res[i] = result[i].Id
	}

	return res
}

func swap(inp []FoundDocument, i, j int) {
	inp[i], inp[j] = inp[j], inp[i]
}

