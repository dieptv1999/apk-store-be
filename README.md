gentool -dsn "review:Techlens@321@tcp(localhost:3306)/apk_store?charset=utf8mb4&parseTime=True&loc=Local" -tables "apk,category,comment,review"

gentool -c "./gen.tool"