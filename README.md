# dousheng-simple-demo-hertz-idl-proto

## 抖音项目服务端简单示例

借鉴自[第五届青训营抖声官方给出的示例项目](https://github.com/RaymondCode/simple-demo)

使用hertz的hz命令配合当时给出的protobuf的idl，生成的hertz代码框架。但是，真的不好用。。。。。

测试中有一些是运行不了的，一部分是没有实现单独存在的chatservice导致的，一部分是跟原项目一样本来就是不能通过测试的。

## hz和hertz存在的问题

现在的hz使用idl生成代码框架不支持绑定文件，因此publish/action中post的内容需要去掉data部分然后自己将这部分代码添加进去。这个问题的具体讨论在[issue](https://github.com/cloudwego/hertz/issues/601)和[文档中](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/#%E7%BB%91%E5%AE%9A%E6%96%87%E4%BB%B6)，本代码中通过修改自动生成的的代码biz/model/core/core.pb.go进行参数接收。

另一个问题是data的数据过大时会出问题，这里是因为hertz默认的文件大小限制参数导致的（查hertz的手册可以知道hertz不设置最大请求体大小的话默认最大只能传4MB的内容）需要手动配置参数。使用`WithMaxRequstBodySize`方法，本代码在main.go中将请求体的最大大小设为了20MB。
