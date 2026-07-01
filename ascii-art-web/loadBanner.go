package main

import (
	"errors"
	"os"
	"strings"
)

var banners = make(map[string]map[rune][]string)
func loadAllBanners() error {
	files := map[string]string{
		"standard": "banner/standard.txt",
		"shadow": "banner/shadow.txt",
		"thinkertoy": "banner/thinkertoy.txt",
	}
	
	for name, path := range files {
		m, err := LoadBanner(path)
		if err != nil {
			return err
		}
		banners[name] = m
	}
	return nil
}

func LoadBanner(filename string) (map[rune][]string, error){
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("Ops! Empty Files")
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) != 855 {
		return nil, errors.New("Ops! Incoplete Banner file")
	}

	result := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		start := i*9+1
		r := rune(i+32)
		result[r] = lines[start:start+8]
	}
	return result, nil
}