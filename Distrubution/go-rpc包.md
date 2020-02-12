# Go net/rpc 包
## 服务端
- `rpc.ServeConn`
## 客户端
- `rpc.Dial()`
- `rpc.DialHTTP()`
- `rpc.Call()`同步处理
- `rpc.Go()`异步处理

## 范例
### 服务端
```go
package server

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}
```
显然，这上面是一个服务端程序。
```go
arith := new(Arith)
rpc.Register(arith)
rpc.HandleHTTP()
l, e := net.Listen("tcp", ":1234")
if e != nil {
    log.Fatal("listen error:", e)
}
go http.Serve(l, nil)
```
这时，服务端注册了一个`Arith`类型的对象和它的公开方法。
### 客户端
```go
client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
if err != nil {
    log.Fatal("dialing:", err)
}
```
客户端与服务端建立连接。
```go
args := &server.Args{7,8}
var reply int
err = client.Call("Arith.Multiply", args, &reply)
if err != nil {
    log.Fatal("arith error:", err)
}
fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
```
同步调用服务端方法。
```go
quotient := new(Quotient)
divCall := client.Go("Arith.Divide", args, &quotient, nil)
replyCall := <-divCall.Done
```
异步调用服务端方法。