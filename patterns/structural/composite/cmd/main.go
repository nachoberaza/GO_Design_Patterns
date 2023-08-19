package main

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/composite/internal/domain"

func main() {
	file1 := &domain.File{Name: "File1"}
	file2 := &domain.File{Name: "File2"}
	file3 := &domain.File{Name: "File3"}

	folder1 := &domain.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &domain.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
