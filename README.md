# SpiderFrame简介
基于GO开发的框架，其实这个不算真正的爬虫框架，里面没有集成任何第三方的爬虫依赖包，大家可以根据自己喜好添加第三方爬虫包。
这个更像一个后台服务框架，毕竟爬虫爬取的数据都是要存入数据库的。所以里面集成了mysql,redis等模块，并附有示例。

推荐集成爬虫包：github.com/gocolly/colly/v2 v2.1.0