package fetcher

import (
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) (string, error){
	resp, err := http.Get(url)
	if err != nil{
		log.Printf("http get url: %s err = %v ",url,err)
		return "",err
	}

	//log.Printf("%+v",resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("read body err")
	}
	html := ConvertToString(string(body), "gbk", "utf-8")
	return html,nil
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}