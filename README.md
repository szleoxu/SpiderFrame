## SpiderFrame简介
基于GO开发的后台服务框架。

其实这个不算真正的爬虫框架，因为里面没有集成任何第三方的爬虫依赖库，也没有重写一套全新的爬虫库。使用者可以根据自己喜好添加最新且更成熟的第三方爬虫库。

这个更像一个**后台数据处理框架**，毕竟爬虫爬取的数据都是要存入数据库的。所以框架集成了处理爬虫数据需要的公共方法和数据库相关的mysql,redis等模块，并附有示例。


**推荐几个优秀的第三方爬虫库**

* [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
* [gocolly/colly](https://github.com/gocolly/colly)
* [chromedp/chromedp](https://github.com/chromedp/chromedp)
