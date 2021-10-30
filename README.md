# CLI Gqlgen Example
這是一個用 Graphql 建立 Golang CLI Tool 的範例
目的是希望不用 Web Server 的同時可以使用 Graphql

# Get Start 🚀
可以用以下的指令試試會輸出什麼結果
```
go run main.go | jq .
```

# 作法
Golang net/http 開啟 Server 的方法是將 Handler 丟入 GoRoutine 不斷請求
替代的解決方案就是利用 httptest 模擬 Request 並執行 Graphql 的 Handler

# 參考
[http request Mocking 的講解](https://budougumi0617.github.io/2020/05/29/go-testing-httptest/)

