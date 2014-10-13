EbookDownloadTools
==================

下载子乌书简上的电子书

为了减轻对子乌书简服务器的压力；EbookDownloadTools一次只能下载一个分类下的书籍

###为什么写这个工具

  去年用Python写了一个脚本将子乌书简上的电子书内容全部下载到本地；
  
  后来子乌书简经历了关闭改版重新开张；所以以前的Python脚本已经不可用；本来想着修改一下之前的Python脚本即可；
  
  不过最近在学习Go；所以就索性用Go重新写了一个程序
  
###用到的一些第三方模块

    1. goquery  ；用于提取网页里面的内容；如标题/url下载地址等等
    2. goini    ；用于从配置文件里面读取内容

###配置文件格式

    [info]
    BookCategory = tljs
    SavePath = F:/book

###配置文件说明

    [info]
    BookCategory = tljs   
    ；要下载的分类书籍|请确保填写的分类地址和http://book.zi5.me/gentre上的分类缩写是一样的；
    ；http://book.zi5.me/archives/book-gentre/tljs 这个是推理惊悚的分类；填写tljs即可
    SavePath = F:/book
    ；本地存放书籍的目录，请填写一个本地已存在的目录
    
