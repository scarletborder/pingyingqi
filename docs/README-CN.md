# 凭依绮在线编译运行服务
a practice of RPC micro service    
[English](../README.md) | 简体中文  

## 调用示例
这里是用python示范的调用案例。
首先你需要复制`config/.env.template`并将其重命名为`.env.config`并且填写你的配置项。    
接着你需要生成接口语言`idl/CodePro.proto`。  
然后你就可以愉快的使用了。    
```python
from plugins.pypyq.idl import CodePro_pb2 as pb2
from plugins.pypyq.idl import CodePro_pb2_grpc as pb2_grpc

channel = grpc.insecure_channel("127.0.0.1:28966")
msgList.append(MessageSegment.text(f"lang={lang}\n"))
stub = pb2_grpc.CodeProProgramerStub(channel)
resp = stub.CodePro(pb2.CodeProRequest(code=code, lang=lang))
```
传入你定义的`code`和`lang`，在`resp`中接受结果。  

## 拓展
你可以在`pingyingqi/service/CodeAi/provider/`下仿照已经写好的案例创建适配其他大语言模型的provider。    
具体来说，你需要实现`pingyingqi/models/AiProvider/interface`定义的接口，以及定义一个init函数，在配置文件相应的api_key配置项不为空时new一个provider，加入单例models/AiProvider/AiHelper.AiHelper中。    
程序会在init时加载你的provider并在Mainprompt()方法的轮询调用中调用你的provider的prompt()成员方法。  