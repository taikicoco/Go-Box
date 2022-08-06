## run 
```
$ go run main.go
```
HTTP/0.9を疑似的に試す
```
$ curl --http1.0 http://localhost:18888/greeting
```
HTTP/1.0
```
$ curl -v --http1.0 http://localhost:18888/greeting
```
ヘッダーの送信
```
$ curl --http1.0 -H "X-Test:Hello" http://localhost:18888/greeting
```
