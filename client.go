package main

import (
    "context"
    "log"
    "time"
    "google.golang.org/grpc"
    pb "training-go-grpc-client-with-buf/gen/go" // プロジェクトのモジュールパスを基に更新
)

func main() {
    // gRPCサーバーのアドレス
    address := "localhost:50051"
    // セキュリティが不要な場合は、WithInsecure() を使用します
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewGreeterClient(conn)

    // コンテキストとタイムアウトの設定
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // SayHello RPCを呼び出し
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}

