package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func WriteToFile(lineToWrite string) {
	// Open the file for appending
	myfile, err := os.OpenFile("myLog.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer myfile.Close()

	// Write the string to the file  date -Iseconds
	time, err := exec.Command("date", "-Iseconds").Output()

	localDate := string(time)
	localDate = strings.TrimSuffix(localDate, "\n")
	lineToWrite = localDate + "  --  " + lineToWrite + "\n"

	_, err = myfile.WriteString(lineToWrite)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The string was appended to the file successfully.")
}
