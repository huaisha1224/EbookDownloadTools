/*
 *Copyright 2014 DownBookZi5
 *author      	=	"Sam Huang"
 *name    		=	"DownBookZi5"
 *version 		=   "0.0.2"
 *url 			=	"http://www.hiadmin.org"
 *author_email	=	"sam.hxq@gmail.com"
 */
package main

import (
	"fmt"
	"github.com/opesun/goquery"
	"github.com/widuu/goini"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetZi5PageUrl(category string) {
	/*
	 *通过传入的分类标签；得到分类地址
	 *用For循环传递分页页码；然后用goquery得到单个书籍的URL地址；
	 *如果分页里面得不到值就认为分页结束
	 *
	 *
	 */
	url := "http://book.zi5.me/archives/book-gentre/"
	//var category string
	for i := 1; i < 100; i++ {
		url := url + category + "/page/" + strconv.Itoa(i) //分页地址
		fmt.Println("\n")
		fmt.Println("分页地址", i, url)
		r, err := goquery.ParseUrl(url)
		if err != nil {
			panic(err)
		} else {
			//查找<class="thumb-holder" > 和<a
			text := r.Find(".thumb-holder a")
			//fmt.Println(text)
			if text.Length() > 0 {
				//取到text里面的所有"href"属性的数据
				for x := 0; x < text.Length(); x++ {
					bookUrl := text.Eq(x).Attr("href")
					DownloadLinks(bookUrl)
				}
			} else { //如果text的长度小于0表示没有找到
				fmt.Print("本分类下已经没有书籍\n")
				break
			}
		}
	}
}

func DownloadLinks(url string) {
	/*
	 *通过传入的书籍URL地址；提取到下载地址
	 *提取书籍的名称和url地址
	 *
	 */
	fmt.Println("\n")
	fmt.Println("书籍地址", url)
	r, err := goquery.ParseUrl(url)
	if err != nil {
		panic(err)
	} else {
		text := r.Find(".download-link")
		bookName := r.Find("h1").Text()
		fmt.Println("书籍名称:", bookName)
		//fmt.Println(text)
		for i := 0; i < text.Length(); i++ {
			downloadlink := text.Eq(i).Attr("href")
			fmt.Println("下载地址", downloadlink)
			DownloadBook(downloadlink, bookName)
		}

	}

}

func DownloadBook(bookUrl, bookName string) {
	conf := goini.SetConfig("./config.ini")
	savePath := conf.GetValue("info", "SavePath") + "/"
	var name string
	if strings.Contains(bookUrl, ".pdf") == true {
		name = ".pdf"
	} else if strings.Contains(bookUrl, ".mobi") == true {
		name = ".mobi"
	} else {
		name = ".epub"
	}
	res, _ := http.Get(bookUrl)
	file, _ := os.Create(savePath + bookName + name)
	if res.Body != nil {
		defer res.Body.Close()
	}
	io.Copy(file, res.Body)
	fmt.Println("下载完成")

}
func main() {
	conf := goini.SetConfig("./config.ini")
	category := conf.GetValue("info", "BookCategory")
	GetZi5PageUrl(category)
	//bookUrl := "http://book.zi5.me/download/50000860/d430.pdf"
	//bookName := "黑色大丽花"
	//downloadbook(bookUrl, bookName)

}
