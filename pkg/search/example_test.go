package search

import (
	"os"
	
	"testing"
)

func TestSearch_GetColNum(t *testing.T) {
	_, err := os.ReadFile("C:/Users/khasa/Desktop/search/lyrics.txt")

	if os.IsNotExist(err) {
		t.Error("file does not exist")
		return
	}
	if err != nil {
		t.Errorf("error ocured while trying to open file %v", err)
	}

	filepath := "C:/Users/khasa/Desktop/search/lyrics.txt"

	position := GetColNum(filepath, "Hello")
	want := 1
	if position != int64(want){
		t.Errorf("want: %v, get: %v", want, position)
	}

	
	position = GetColNum(filepath, "my")
	want = 17
	if position != int64(want){
		t.Errorf("want: %v, get: %v", want, position)
	}

}

func TestSearch_GetLineNum(t *testing.T) {
	filepath := "C:/Users/khasa/Desktop/search/lyrics.txt"

	line := GetLineNum(filepath, "Hello")
	want := 1
	if line != int64(want){
		t.Errorf("want: %v, get: %v", want, line)
	}

	line2 := GetLineNum(filepath, "Left")
	want2 := 4
	if line2 != int64(want2){
		t.Errorf("want: %v, get: %v", want2, line2)
	}
		
}

func TestSearch_GetLine(t *testing.T) {
	filepath := "C:/Users/khasa/Desktop/search/lyrics.txt"
	 line := GetLine(filepath, "Left")
	 want := "Left its seeds while I was sleeping"
	 if line != want{
		t.Errorf("want: %v, get: %v", want, line)
	 }
}