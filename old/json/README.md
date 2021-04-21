# json

该包实现了一个github.com/mrchar/hraf中定义的Encoder  

Encoder将传入的内容一restful的json形式写入http response中  
将http request中的内容解码到传入的struct中  

## 解码参数

根据结构体定义中的tag进行解码  

```
type Params struct {
    Limit   int     `hraf:"limit,url"`
    Offset  int     `hraf:"offset,url"`
    Token   string  `hraf:"WWW-Authenticate,header"`
    ID      string  `hraf:"id,content"`
}
```

hraf标签描述这个参数应该如何获取，第一个参数为参数的key，第二个参数为获取的来源。  
比如Limit的key为limit，从url参数中获取，则Encoder会在url参数中查找名称为limit、类型为int的参数。  
url表示从url参数获取，header表示从http header中获取，content表示从正文中获取。  
如果没有表明来源，首先从正文中获取、然后从url参数中获取、最后从header中获取。