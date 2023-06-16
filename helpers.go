package main

import (
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

func getSize(path string) string {
	info, err := os.Stat(path)

	if err != nil {
		panic(err)
	}

	if !info.IsDir() {
		return formatSize(info.Size())
	}

	totalDirSize := int64(0)

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		totalDirSize += info.Size()

		return nil
	})

	if err != nil {
		panic(err)
	}

	return formatSize(totalDirSize)
}

func getSizeVerbose(path string) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	totalSize := int64(0)

	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
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
		panic(err)
	}

	fmt.Printf("\nTotal size:%*s\n", width-15, formatSize(totalSize))
}
