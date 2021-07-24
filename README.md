# HTCD-JSON (HTTP-Response-Code-To-Json)
## About
`hosts.csv`に記述されているURLにGETリクエストを送信し、返ってきたHTTPレスポンスコードとURLのペアをjsonファイルに保存します。
```json
  {
    "disp": "DISPLAY NAME",
    "url": "http://example.com/",
    "code": 200,
    "time": "2000-01-01T00:00:00.0000000+09:00"
  }
```
上記のようなjson形式で保存されます。
```csv
DISPLAYNAME,http://example.com/
DISPLAYNAME2,http://fuga.example.com/
```
`hosts.csv`は上記のように記述してください。

csvファイルは必ず同じディレクトリに置いてください。
## License
The source code is licensed MIT. see LICENSE.