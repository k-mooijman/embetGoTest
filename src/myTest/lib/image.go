package lib

import (
	"fmt"
	"log"
	"os/exec"
)

func GetImageData(file string) {

	//exiftool -json -DateTimeOriginal -d "%Y-%m-%d_%H-%M-%S %f"  -make -model -GPSDateTime -SubSecDateTimeOriginal

	filename := "/home/kasper/Downloads/20240611_225243.jpg"
	command := "/usr/bin/exiftool"
	//args := []string{"/usr/bin/exiftool", "-time:SubSecDateTimeOriginal", "-G1", "-a", "-s", filename}
	args := []string{"-json", "-DateTimeOriginal", "-d", "\"%Y-%m-%d_%H-%M-%S %f\"", "-make", "-model", "-GPSDateTime", "-SubSecDateTimeOriginal", filename}

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("variable out = %v is of type %T \n", out, out)
	myString := string(out)
	fmt.Printf("\n\n=========================\n\n")

	fmt.Printf("%v  \n\n", myString)

	fmt.Printf("\n\n=========================\n\n")

}
