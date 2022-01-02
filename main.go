package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	path := pwd()

	myApp := app.New()
	myWindow := myApp.NewWindow("Subtitle Renamifier")

	content := widget.NewButton("Rename Subtitles", func() {
		Renamify(path)
	})

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func Renamify(path string) {
	sr, _ := walkMatch(path, "*.srt")
	srts := getFileName(sr)

	vd, _ := walkMatch(path, "*.mkv")
	mkvs := getFileName(vd)

	re := regexp.MustCompile(`S\d+E\d+`)
	for _, v := range mkvs {
		match := re.FindStringSubmatch(v)
		for _, s := range srts {
			c := strings.Contains(s, match[0])
			if c {
				err := os.Rename(path+"/"+s, path+"/"+strings.Replace(v, ".mkv", ".srt", 1))
				if err != nil {
					return
				}
			}
		}
	}
}

func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

func getFileName(fps []string) (fns []string) {
	for _, v := range fps {
		fns = append(fns, filepath.Base(v))
	}

	return fns
}

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
