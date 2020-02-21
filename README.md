# crawl some user's information by single process

Follow the steps below

Step 1:
    创建数据库spider_profile
    导入本地数据库表结构，文件名：profile_info.sql后，生成profile_info数据表
	
Step 2:
    go build main.go 
	
Step 3:
    chmod 777 程序名称
	
Step 4:
    nohup ./程序名 &
	
Step 5:
    查看数据表数据，目前只设置了爬一个城市的用户，如果要修改爬取城市的长度，在config.ini中修改Common下Length的大小