/*
 *Copyright 2014 EbookDownloadTools
 *author      	=	"Sam Huang"
 *name    		=	"EbookDownloadTools"
 *version 		=   "0.6"
 *url 			=	"http://www.hiadmin.org"
 *author_email	=	"sam.hxq@gmail.com"
 *
 *子乌书简是我看电子书内容的主要来源；书简里面的书由子午手动校对所以质量都比较可靠；
 *去年子乌书简关闭之前我写过一个Python脚本吧所有的书籍下载回来；今年子乌书简重开之后Python不能用
 *所以就用GOlang重新写了一个程序来下载子乌书简上的书籍；这次为了减轻对服务器的压力；
 *不能一次性下载全部书籍；一次性只能下载一个分类；我想很多书友也不需要下载全部的书籍。
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
					GetDownloadLinks(bookUrl)
				}
			} else { //如果text的长度小于0表示没有找到
				fmt.Print("本分类下已经没有书籍\n")
				break
			}
		}
	}
}

func GetDownloadLinks(url string) {
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
	/*
	 *通过传入书籍的下载URL和书籍名称来下载书籍并命名
	 *获取配置文件里面的存放路径
	 *通过判断传入的下载URL地址结尾来判断文件名称是pdf/mobi/epub格式
	 *然后用http.Get访问书籍地址
	 *最后用io.Copy拷贝文件到本地
	 */
	conf := goini.SetConfig("./config.ini")
	category := conf.GetValue("info", "BookCategory")
	tempSavePath := conf.GetValue("info", "SavePath")
	bookType := conf.GetValue("info", "BookType")
	savePath := tempSavePath + "/" + category

	if _, err := os.Stat(tempSavePath); err != nil {
		os.Mkdir(tempSavePath, 0777)
		os.Mkdir(savePath, 0777)
	} else {
		os.Mkdir(savePath, 0777)
	}

	//判断传入的下载地址结尾是pdf/mobi/epub、用来区分文件类型
	var name string
	if strings.Contains(bookUrl, ".mobi") == true {
		name = ".mobi"
	} else if strings.Contains(bookUrl, ".epub") == true {
		name = ".epub"
	} else if strings.Contains(bookUrl, ".azw") == true {
		name = ".azw"
	} else {
		name = ".pdf"
	}

	//判断用户希望下载的书籍类型并下载保存到本地
	if name == bookType || name == ".pdf" {
		res, _ := http.Get(bookUrl)
		file, _ := os.Create(savePath + "/" + bookName + name)
		if res.Body != nil {
			defer res.Body.Close()
		}
		io.Copy(file, res.Body)
		fmt.Println("下载完成")

	}

}
func main() {
	conf := goini.SetConfig("./config.ini")
	category := conf.GetValue("info", "BookCategory")
	GetZi5PageUrl(category)

}
