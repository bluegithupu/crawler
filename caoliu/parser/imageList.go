package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
)

func imgeListParse(html string) []Request {

	//<img data-src='https://www.privacypic.com/images/2019/12/10/04122d406fbeb193a3a.jpg'>&nbsp;
	//<img[^']+data-src='[^']+'>&nbsp;
	//title
	re := regexp.MustCompile(`<title>(.+)</title>`)
	match := re.FindAllStringSubmatch(html,-1)
	log.Println(match[0][1])
	title := match[0][1]


	re = regexp.MustCompile(`([a-zA-z]+://[^\s]*)'>&nbsp;`)
	match = re.FindAllStringSubmatch(html,-1)
	for _, m := range match {
		log.Println(m[1])
		saveImage(m[1],title)
	}
	return nil
}

func saveImage(url string,dirName string){
	imgPath := `H:\image\` + dirName[0:20]
	imgUrl := url

	fileName := path.Base(imgUrl)
	makeDir(imgPath)
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)



	file, err := os.Create( imgPath + `\` + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	io.Copy(writer, reader)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func makeDir(imgPath string){

	exist, err := PathExists(imgPath)
	if err != nil{
		panic(err)
	}

	if exist {
		fmt.Printf("has dir![%v]\n", imgPath)
	} else {
		fmt.Printf("no dir![%v]\n", imgPath)
		// 创建文件夹
		err := os.Mkdir(imgPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}
