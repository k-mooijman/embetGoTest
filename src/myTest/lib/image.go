package lib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

type ImageData struct {
	SourceFile             string `json:"SourceFile" sql:"SourceFile"`
	DateTimeOriginal       string `json:"DateTimeOriginal" sql:"DateTimeOriginal"`
	Make                   string `json:"Make" sql:"Make"`
	Model                  string `json:"Model" sql:"Model"`
	GPSDateTime            string `json:"GPSDateTime" sql:"GPSDateTime"`
	SubSecDateTimeOriginal string `json:"SubSecDateTimeOriginal" sql:"SubSecDateTimeOriginal"`
	Hash                   string `json:"Hash" sql:"Hash"`
	Extension              string `json:"Extension" sql:"Extension"`
	OSDateTime             string `json:"OSDateTime" sql:"OSDateTime"`
	Size                   int64  `json:"Size" sql:"Size"`
}

func (image *ImageData) AddHash() {
	image.Hash = calcSha256(image.SourceFile)
}

func (image *ImageData) AddOsData() {

	fileInfo, err := os.Stat(image.SourceFile)

	if err != nil {
		log.Fatal(err)
	}
	image.OSDateTime = fileInfo.ModTime().Format("2006-01-02 15:04:05")
	image.Size = fileInfo.Size()
	image.Extension = path.Ext(fileInfo.Name())
}

func GetImageData(filename string) (ImageData, error) {

	command := "/usr/bin/exiftool"
	args := []string{"-json", "-DateTimeOriginal", "-d", "\"%Y-%m-%d_%H-%M-%S %f\"", "-make", "-model", "-GPSDateTime", "-SubSecDateTimeOriginal", filename}

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	myString := string(out)
	//
	//fmt.Printf("\n\n=========================\n\n")
	//
	//fmt.Printf("%v  \n\n", myString)
	//
	//fmt.Printf("\n\n=========================\n\n")

	var imageBlock []ImageData

	errJ := json.Unmarshal([]byte(myString), &imageBlock)
	if errJ != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}
	if len(imageBlock) == 0 {
		var imageData ImageData
		return imageData, errors.New("could not collect Exif data")
	}
	var image = imageBlock[0]
	return image, nil
}

func calcSha256(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha
}
