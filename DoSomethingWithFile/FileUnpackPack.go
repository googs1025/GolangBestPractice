package DoSomethingWithFile

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)
// archive包中提供了tar、zip两种打包/解包方法

func unpackZip(zipfile string)  {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
	}
}

func writerZip()  {
	// Create archive
	zipPath := "out.zip"
	zipFile, err := os.Create(zipPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new zip archive.
	w := zip.NewWriter(zipFile)
	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"asong.txt", "This archive contains some text files."},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

