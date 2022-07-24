package file

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func PathExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CurrentUserDir() string {

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return u.HomeDir

	//fmt.Println("Home dir:", u.HomeDir)
}

func AppDataDir(name string) string {
	return CurrentUserDir() + string(os.PathSeparator) + name
}

func ReadToString(filePath string) string {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(fileContent)
	return text
}

func ReadLines(filePath string) []string {

	readFile, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func writeLines(lines []string, path string) (err error) {
	var (
		file *os.File
	)

	if file, err = os.Create(path); err != nil {
		return
	}
	defer file.Close()

	//writer := bufio.NewWriter(file)
	for _, item := range lines {
		//fmt.Println(item)
		_, err := file.WriteString(strings.TrimSpace(item) + "\n")
		//file.Write([]byte(item));
		if err != nil {
			//fmt.Println("debug")
			log.Fatal(err)
			break
		}
	}
	/*content := strings.Join(lines, "\n")
	  _, err = writer.WriteString(content)*/
	return
}

func WriteToString(filePath string, content string) bool {
	fileObj, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer fileObj.Close()
	//创建一个往文件里写的对象
	writer := bufio.NewWriter(fileObj)

	writer.WriteString(content) //写到了缓存

	writer.Flush() //将缓存中的内容写到文件里
	return true
}
