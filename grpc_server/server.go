package main

import (
	"log"
	"net"
	"net/rpc"
	"sync"

	"learn.com/hw/grpc"
)

//
// Server
//

type KV struct {
	mu   sync.Mutex
	data map[string]string
}

func server() {
	kv := new(KV)
	kv.data = map[string]string{}
	rpcs := rpc.NewServer()
	rpcs.Register(kv)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("server listening at %v", l.Addr())
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		} else {
			log.Fatalf("failed to serve req: %v", err)
			break
		}
	}
}

func (kv *KV) Get(args *grpc.GetArgs, reply *grpc.GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	val, ok := kv.data[args.Key]
	if ok {
		reply.Err = grpc.OK
		reply.Value = val
	} else {
		reply.Err = grpc.ErrNoKey
		reply.Value = ""
	}
	return nil
}

func (kv *KV) Put(args *grpc.PutArgs, reply *grpc.PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[args.Key] = args.Value
	reply.Err = grpc.OK
	return nil
}

func main() {
	server()
}
