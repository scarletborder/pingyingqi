# 凭依绮在线Code
RPC微服务
## 示例
以下内容都是从作者的一个nonebot2插件扣下来的部分
### GoPro() and PyPro()
```py
channel = grpc.insecure_channel("127.0.0.1:28966")
if lang == "go":
    msgList.append(MessageSegment.text("lang=golang\n"))
    stub = pb2_grpc.ExlangProgramerStub(channel)
    resp = stub.GoPro(pb2.ExlangRequest(code=code))

elif lang == "py":
    msgList.append(MessageSegment.text("lang=python\n"))
    stub = pb2_grpc.ExlangProgramerStub(channel)
    resp = stub.PyPro(pb2.ExlangRequest(code=code))
    pass

else:
    msgList.append(MessageSegment.text("lang=nil\n"))
    resp = pb2.ExlangResp(data="default fail", code=1)

# according to the code, output different info
if resp.code == 3:
    msgList.append(
        MessageSegment.text(
            "WARNING:Your code compiled successfully but run too much time\n"
        )
    )
    msgList.append(MessageSegment.text(resp.data))
elif resp.code == 2:
    msgList.append(
        MessageSegment.text(
            "ERROR:Your code contains some disabled package\n" + resp.data
        )
    )
elif resp.code == 1:
    msgList.append(MessageSegment.text("ERROR:Program failed\n" + resp.data))

elif resp.code == 0:
    msgList.append(MessageSegment.text(resp.data))
else:
    msgList.append(MessageSegment.text("ERROR:Inner wrong\n" + resp.data))

await pypyq.finish(Message(msgList))
```

### Dislike()
```py
if lang := args.extract_plain_text():
    myList = lang.split(" ")
    if len(myList) < 2:
        await pyqDis.finish("too few args")
    data = ""
    if myList[0] == "enable":
        data += "0;"
    else:
        data += "1;"

    channel = grpc.insecure_channel("127.0.0.1:28966")
    stub = pb2_grpc.ExlangProgramerStub(channel)
    resp = stub.Dislike(pb2.DislikedPackage(pack=data + myList[1]))
    pass
```