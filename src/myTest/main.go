package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"mooijman.info/myTest/src/myTest/lib"

	_ "mooijman.info/myTest/src/myTest/lib"
)

func main() {

	//lpass show --password 5997716042849405729

	token, err := exec.Command("lpass", "show", "--password", "5997716042849405729").Output()
	if err != nil {
		log.Fatal(err)
	}

	lib.TelegramBod(string(token))
	userID := int64(6065413981) //	kasper 0648947942
	//userID := int64(7344285136) //  Kasper 44221313
	lib.SendMessageToUser(userID, "Bosque just woke up!")

	//misteryBox := interface{}(10)
	//
	//test, ok := misteryBox.(int)
	//
	//if ok {
	//	fmt.Printf("images = %v is of type %T \n", test, test)
	//}

	//lib.FileWatcher()

	//#############################   Prompt   ################################

	//lib.BasicPrompt()
	//lib.SelectCustomPrompt()
	//lib.CustomPrompt()

	//lib.SelectAddPrompt()
	//#############################   prompt   ################################
	//#############################   Recursive   ################################
	//
	//var rec1a lib.Rec
	//rec1a.Val = 3
	//
	//var rec2a lib.Rec
	//rec2a.Val = 3
	//
	//var rec3a lib.Rec
	//rec3a.RecVal1 = &rec1a
	//rec3a.RecVal2 = &rec2a
	//
	//var rec1 lib.Rec
	//rec1.Val = 3
	//
	//var rec2 lib.Rec
	//rec2.Val = 3
	//
	//var rec3 lib.Rec
	//rec3.RecVal1 = &rec1
	//rec3.RecVal2 = &rec2
	//
	//var rec0 lib.Rec
	//rec0.RecVal1 = &rec3a
	//rec0.RecVal2 = &rec3
	//
	//rec0.Calc()
	//fmt.Println(rec0.Val)

	//#############################   Recursive   ################################

	//#############################   JSON   ################################
	//jsonStr := `{"name":"John Doe", "age":30, "isEmployed":true , "info":{"one":1, "two":2}}`
	//var result map[string]interface{}
	//
	//err := json.Unmarshal([]byte(jsonStr), &result)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(result)
	//#############################   JSON   ################################
	//#############################   imageWalker   ################################
	//
	//image, err := lib.GetImageData("/home/kasper/Downloads/20240611_225243.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////image.AddHash()
	////image.AddOsData()
	//
	//fmt.Printf("images = %v is of type %T \n", image, image)
	//fmt.Printf("hash = %v is of type %T \n", image.Hash, image.Hash)
	//fmt.Printf("Ext = %v is of type %T \n", image.Extension, image.Extension)
	//fmt.Printf("time = %v is of type %T \n", image.OSDateTime, image.OSDateTime)
	//fmt.Printf("size = %v is of type %T \n", image.Size, image.Size)
	//#############################   imageWalker   ################################

	//lib.Stat()

	//log.Printf("\n\nFiles in folder %v  \n", lib.CountFiles())

	//go lib.StartApi()
	//todo := lib.Todo{"2", "duStuff", true}
	//lib.AddTodo(todo)

	//lib.CallbackTest()

	//lib.Test()
	//lib.TestRead()

	//go embed.Start()

	//var testNr = 13
	//log.Printf("\n\nThis is test %v of me trying Go\n", testNr)

	//lib.PrintArguments()

	//startTime := time.Now()
	// Perform some operations

	//files, err := listFiles("/home/kasper/Downloads")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(files)

	// Calculate elapsed time
	//elapsed := time.Since(startTime)
	//fmt.Println("Elapsed time:", elapsed)

	time.Sleep(5 * time.Second)
	waitForQ()

}

func listFiles(dir string) ([]string, error) {
	listing := []string{}
	file, err := os.Open(dir)
	if err != nil {
		return listing, err
	}
	files, err := file.Readdir(0)
	if err != nil {
		return listing, err
	}
	for _, file := range files {
		if !file.IsDir() {
			fName := file.Name()

			sha := calcSha256(dir + "/" + file.Name())
			time := file.ModTime()
			extension := path.Ext(dir + "/" + file.Name()) //obtain the extension of file

			listing = append(listing, "\n", sha, "  -  ", extension, "  -  ", time.Format("2006-01-02 15:04:05"), "  -  ", strconv.FormatInt(file.Size(), 10), "  -  ", fName)
		}
	}
	return listing, nil
}

func waitForQ() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter command (q/quit/exit): ")
		scanner.Scan()
		fmt.Println(scanner.Text())

		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		if "q" == scanner.Text() || "quit" == scanner.Text() || "exit" == scanner.Text() {
			fmt.Println("Exiting gracefully ")
			os.Exit(0)
		}

	}
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
	//fmt.Printf("variable sha=%v is of type %T \n", sha, sha)

	return sha
}

func getImageData(file string) {

	filename := "/home/kasper/Downloads/20240611_225243.jpg"
	command := "/usr/bin/exiftool"
	args := []string{"/usr/bin/exiftool", "-time:SubSecDateTimeOriginal", "-G1", "-a", "-s", filename}

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("variable out = %v is of type %T \n", out, out)
	myString := string(out)
	fmt.Printf("variable myString = %v is of type %T \n", myString, myString)

	before, after, found := strings.Cut(myString, ":")
	fmt.Printf("strings.Cut():\nbefore: %s\nafter: %s\nseparator found: %t\n", before, after, found)

	//fmt.Printf(" \n\n exiftool \n\n")
	//fmt.Printf("The date from image is %s\n", strings.SplitN(myString, ":", 2))
	//fmt.Printf(" \n\n exiftool \n\n")
}
