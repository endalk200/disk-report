package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

func formatSize(size int64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func main() {
	flag.Parse()
	root := "."
	if flag.Arg(0) != "" {
		root = flag.Arg(0)
	}

	var totalSize int64

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		indent := strings.Repeat("  ", strings.Count(path, string(os.PathSeparator))-1)

		if !info.IsDir() {
			size := info.Size()
			totalSize += size
			sizeStr := formatSize(size)
			fmt.Printf("%s%s%*s\n", indent, path, width-len(indent)-len(path)-4, sizeStr)
		} else {
			fmt.Printf("%s%s\n", indent, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalSizeStr := formatSize(totalSize)
	fmt.Printf("\nTotal size:%*s\n", width-15, totalSizeStr)
}
