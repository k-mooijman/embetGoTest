package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

func CountFiles() int {
	count := 0
	imagesFiles := 0
	videoFiles := 0

	nonImagesFiles := 0
	extension := make(map[string]int)
	otherExt := []string{".bin"}

	//rootPath := "/home/kasper/Downloads/documents/"
	//rootPath := "/home/kasper/Downloads/"
	//rootPath := "/home/kasper/"
	//rootPath := "/media/kasper/S990P2TB/"
	rootPath := "/media/kasper/S990P2TB/fotos"

	filepath.WalkDir(rootPath, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !file.IsDir() {
			file.Type()
			//if !file.IsDir() {

			//fmt.Println(path)
			//info := file.Type()
			//if err != nil {
			//	return nil
			//}

			//fmt.Printf("variable info.Name() = %v is of type %T \n", info.Type(), info.Type())

			images := []string{".jpg", ".jpeg", ".bmp", ".gif", ".png", ".ttf", ".png", ".nef", ".svg", ".jpg_large", ".dng", ".thm"}
			videos := []string{".3gp", ".mov", ".mjpeg", ".mp4"}

			ext := strings.ToLower(filepath.Ext(path))

			if slices.Contains(images, ext) {
				//fmt.Println(path)
				//GetImageData(path)
				imagesFiles++
			} else if slices.Contains(videos, ext) {
				videoFiles++
			} else {
				nonImagesFiles++

				val, ok := extension[ext]
				// If the key exists
				if ok {
					extension[ext] = val + 1
				} else {
					extension[ext] = 1
				}

				if !slices.Contains(otherExt, strings.ToLower(filepath.Ext(path))) && len(filepath.Ext(path)) < 5 {
					otherExt = append(otherExt, strings.ToLower(filepath.Ext(path)))
				}

				fmt.Println(path)
				fmt.Println(strings.ToLower(filepath.Ext(path)))

			}
			count++
		}

		return nil
	})

	fmt.Printf("variable otherExt = %v is of type %T \n \n \n", otherExt, otherExt)
	fmt.Printf("imagesFiles = %v is of type %T \n", imagesFiles, imagesFiles)
	fmt.Printf("videoFiles = %v is of type %T \n", videoFiles, videoFiles)
	fmt.Printf("nonimages = %v is of type %T \n", nonImagesFiles, nonImagesFiles)
	fmt.Printf("extension = %v is of type %T \n", extension, extension)

	return count
}
