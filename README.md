# Pingyingqi Online compiling and running
a practice of RPC micro service  
English | [简体中文](docs/README-CN.md)

## Usage Example
Here is a call example written in Python  
First, Copy `config/.env.template` and rename it to `.env.config`. And fill your config items.   
You need to use the idl `idl/CodePro.proto` to generate an interface class.  
Then, use it like this  

```python
from plugins.pypyq.idl import CodePro_pb2 as pb2
from plugins.pypyq.idl import CodePro_pb2_grpc as pb2_grpc

channel = grpc.insecure_channel("127.0.0.1:28966")
msgList.append(MessageSegment.text(f"lang={lang}\n"))
stub = pb2_grpc.CodeProProgramerStub(channel)
resp = stub.CodePro(pb2.CodeProRequest(code=code, lang=lang))
```
use the defined parameter `code` and `lang` previously, you can receive your result in `resp`.

## Extension
You can create a provider to adapt other major language models under `pingyingqi/service/CodeAi/provider/`, following the examples already written.  
Specifically, you need to implement the interface defined in `pingyingqi/models/AiProvider/interface` and define an init function. When the corresponding api_key configuration item in the configuration file is not empty, instantiate a provider and add it to the singleton `models/AiProvider/AiHelper.AiHelper`.  
The program will load your provider during the init phase and call the `prompt()` method of your provider during the polling calls in the `Mainprompt()` method.  