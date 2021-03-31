package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var resourceDir = "resources/images/"

func RemoveFile(folder string, name string) {
	e := os.Remove(name)
	if e != nil {
		fmt.Printf("File not deleted")
	}
}

func GetFullName(name string, folder string) string {

	return resourceDir + folder + "/" + name

}

func UploadFile(r *http.Request, paramName string, folder string) (error, string) {
	file, handler, err := r.FormFile(paramName)
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return err, ""
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(resourceDir+folder, folder+"-*.png")
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return err, ""
	}
	tempFile.Write(fileBytes)

	return nil, tempFile.Name()[strings.LastIndex(tempFile.Name(), "/")+1:] // return file name from full dir
}
