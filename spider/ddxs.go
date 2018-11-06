package spider

import (
	"fmt"
	"fiction/common"
	"fiction/models"
	"github.com/PuerkitoBio/goquery"
	"database/sql"
	"fiction/config"
	"time"
	"os"
)


func DDXS(){
	start := time.Now()
	fmt.Println("spider is running")
	//get book list
	urlDDXS:=config.GetValue("spiderUrl","ddxs")
	dom:=common.UrlResponse(urlDDXS)
	if dom!=nil{
		dbBook:=models.InitBook()
		dbChapter:=models.InitChapter()
		dom.Find("#hotcontent > div.l > div").Each(func(i int, content *goquery.Selection) {
			name := common.GbkToUtf8(content.Find("dl > dt > a").Text())
			fmt.Println("Find book:"+name)
			isExist:=dbBook.IsExistBook(name)
			if isExist{
				fmt.Println("Exist book:"+name)
			}else{
				author := common.GbkToUtf8(content.Find("dl > dt > span").Text())
				image,_ := content.Find("div.image > a > img").Attr("src")
				tbBook:=models.BookTB{}
				var bookName=sql.NullString{name,true}
				tbBook.Name=bookName
				tbBook.Author=author
				tbBook.Image=image
				tbBook.SiteUrl=urlDDXS
				insertLastID:=dbBook.InsertData(tbBook)
				if insertLastID>0{
					//get chapter list
					chapterListUrl,_:= content.Find("div.image > a").Attr("href")
					dom=common.UrlResponse(chapterListUrl)
					var chapterNum=0
					dom.Find("#list > dl > dt:nth-child(10)").NextAll().EachWithBreak(func(i int, content *goquery.Selection) bool{
						//get chapter
						chapterUrl,_:= content.Find("a").Attr("href")
						chapterUrl=chapterListUrl+chapterUrl
						dom=common.UrlResponse(chapterUrl)
						if dom!=nil{
							chapterNum++
							title:=common.GbkToUtf8(dom.Find("#wrapper > div.content_read > div > div.bookname > h1").Text())
							isExist=dbChapter.IsExistChapter(insertLastID,title)
							if isExist{
								fmt.Println("Exist chapter:"+title)
							}else{
								chapterContent:=dom.Find("#content").Text()
								chapterContent=common.GbkToUtf8(chapterContent)
								pre,_:=dom.Find("#wrapper > div.content_read > div > div.bottem2 > a:nth-child(2)").Attr("href")
								pre=chapterListUrl+pre
								next,_:=dom.Find("#wrapper > div.content_read > div > div.bottem2 > a:nth-child(4)").Attr("href")
								next=chapterListUrl+next
								tbChapter:=models.ChapterTB{}
								tbChapter.Bid=insertLastID
								var chapterTitle=sql.NullString{title,true}
								tbChapter.Title=chapterTitle
								tbChapter.Content=chapterContent
								tbChapter.Sort=chapterNum
								tbChapter.Pre=pre
								tbChapter.Next=next
								isInsert:=dbChapter.InsertData(tbChapter)
								if isInsert==true{
									fmt.Println("Insert chapter success")
								}else{
									fmt.Println("Insert chapter fail")
								}
							}
						}
						return true
					})
				}
			}
		})
		dbBook.Db.Close()
		dbChapter.Db.Close()
		fmt.Println("complete")
	}
	cost := time.Since(start)
	fmt.Printf("time=[%s]",cost)
	defer os.Exit(0)
}

func AfreshBook(bookID int64){

}

func AfreshChapter(bookID int64,chapterID int64){

}