package search

import (
	"bufio"
	"context"
	
	"os"
	"strings"
	"sync"
)

//Resuлt описывает один результат поиска.
type Result struct{
	//The phrase you were looking for
	Pharse string
	//Entire line in which the entry was found without \ n or \ r \ n at the end
	Line string
	//line number (starting from 1) where the entry was found
	LineNum int64
	//the position number (starting from 1) at which the entry was found
	ColNum int64
}

//All searches for all occurrences of phrase in text files files
func All(ctx context.Context, phrase string, files []string) <-chan []Result{
	ch := make(chan []Result)
	r := []Result{}
// TODO: your code
// 1. run one goroutine for each file
// 2. Send each result to the chan channel
// 3. as soon as the search is completed • close the channel
	
	goroutines := len(files)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i:=0; i<goroutines; i++{
		wg.Add(1)
		for _, v := range files{
			go func(file string) {
				Line := GetLine(file, phrase)
				LineNum := GetLineNum(file, phrase)
				ColNum := GetColNum(file, phrase)
				
				result := Result{
					Pharse: phrase,
					Line: Line,
					LineNum: LineNum,
					ColNum: ColNum,
				}
				mu.Lock()
				r = append(r, result)
				mu.Unlock()
				
				
			}(v)
		}
	}

	ch <- r

	return ch
}

// function to get ColNum (starting with 1)
func GetColNum(file string, phrase string) int64{
	content, err := os.ReadFile(file)
	if os.IsNotExist(err){
		return 0
	}
	if err != nil {
		return 0
	}

	position := strings.Index(string(content), phrase)

	position += 1

	return int64(position)

}

// function to get LineNum (starting with 1)
func GetLineNum(file, phrase string) int64{
	NumLine := int64(0)
	line := int64(1)
	f, _ := os.Open(file)
	
	defer f.Close()
	scanner := bufio.NewScanner(f)
	
	for scanner.Scan() {
	if strings.Contains(scanner.Text(), phrase) {
		NumLine = line
	}
	line++
	}

	return NumLine
}

// function to get LineNum
func GetLine(file, phrase string) (line string) {
	lineNum := GetLineNum(file, phrase)
	lastLine := int64(0)
	f, _ := os.Open(file)
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
		    return sc.Text()
		}
	  }

	  return line
}