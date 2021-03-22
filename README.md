# SpiderFrame简介
基于GO开发的框架。其实这个不算真正的爬虫框架，里面没有集成任何第三方的爬虫依赖库，也没有自己重写一套爬虫。大家可以根据自己喜好添加被市场验证过的成熟而更专业的第三方爬虫库。
这个更像一个**后台服务框架**，毕竟爬虫爬取的数据都是要存入数据库的。所以里面集成了mysql,redis等模块，并附有示例。

推荐几个优秀的第三方爬虫库

* [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
* [gocolly/colly](https://github.com/gocolly/colly)
* [chromedp/chromedp](https://github.com/chromedp/chromedp)
