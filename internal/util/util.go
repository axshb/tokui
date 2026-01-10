package util

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/pkoukk/tiktoken-go"
)

// tokeniser wraps tiktoken
type Tokenizer struct {
	encoding *tiktoken.Tiktoken
}

func NewTokenizer() (*Tokenizer, error) {
	tke, err := tiktoken.GetEncoding("cl100k_base")
	if err != nil {
		return nil, err
	}
	return &Tokenizer{encoding: tke}, nil
}

func (t *Tokenizer) Count(text string) int {
	return len(t.encoding.Encode(text, nil, nil))
}

// implements list.Item
type FileItem struct {
	TitleStr, PathStr string
}

func (i FileItem) Title() string       { return i.TitleStr }
func (i FileItem) Description() string { return i.PathStr }
func (i FileItem) FilterValue() string { return i.TitleStr }

func GetFiles(dir string) []list.Item {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var items []list.Item
	for _, entry := range entries {
		if !entry.IsDir() && entry.Name()[0] != '.' {
			items = append(items, FileItem{
				TitleStr: entry.Name(),
				PathStr:  filepath.Join(dir, entry.Name()),
			})
		}
	}
	return items
}
