# Mock

这是一个简单的mock工具

# 开始使用
[必须] 在程序目录下创建apis.json文件，这个文件用来定义请求地址和返回结果文件。
如：
```json
[
  {
    "path": "/api/xx",
    "resp_file": "mock1.json"
  },
  {
    "path": "/api/yy",
    "resp_file": "mock2.pdf"
  }
]
```

- path表示请求地址

- resp_file表示返回数据，如果是文件，可以直接填写路径名称，如果是json，直接填写json文件的路径即可
