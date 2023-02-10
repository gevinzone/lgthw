package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gevinzone/lgthw/lgthw/c07Micro/grpcjson/keyvalue/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
)

type Client struct {
	gen.KeyValueClient
}

func NewClient(cc grpc.ClientConnInterface) *Client {
	return &Client{
		KeyValueClient: gen.NewKeyValueClient(cc),
	}
}

//func (c *Client) set(ctx context.Context, r *gen.SetKeyValueRequest) (*gen.KeyValueResponse, error) {
//	return c.Set(ctx, r)
//}
//
//func (c *Client) get(ctx context.Context, r *gen.GetKeyValueRequest) (*gen.KeyValueResponse, error) {
//	return c.Get(ctx, r)
//}

func (c *Client) GetHandler(w http.ResponseWriter, r *http.Request) {
	getReq := gen.GetKeyValueRequest{Key: r.URL.Query().Get("key")}
	resp, err := c.Get(r.Context(), &getReq)
	if err != nil {
		if errors.Is(status.Errorf(codes.NotFound, "key not set"), err) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("404 Not Found"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func (c *Client) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req gen.SetKeyValueRequest
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := c.Set(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ = json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

//type KeyValueRequest struct {
//	Key   string `json:"key,omitempty"`
//	Value string `json:"value,omitempty"`
//}

//func (k *KeyValueRequest) ToGrpcRequest() *gen.SetKeyValueRequest {
//	return &gen.SetKeyValueRequest{
//		Key:   k.Key,
//		Value: k.Value,
//	}
//}

func main() {
	//cc, err := grpc.Dial(":8888", grpc.WithInsecure())
	cc, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer cc.Close()
	if err != nil {
		panic(err)
	}
	//client := gen.NewKeyValueClient(cc)
	//resp, err := client.Get(context.Background(), &gen.GetKeyValueRequest{Key: "key"})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp)
	client := NewClient(cc)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/get", client.GetHandler)
	http.HandleFunc("/set", client.SetHandler)
	fmt.Printf("start serve on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	//r := &gen.SetKeyValueRequest{Key: "a", Value: "b"}
	//resp, err := client.Set(context.Background(), r)
	////resp, err := client.Get(context.Background(), &gen.GetKeyValueRequest{Key: "key"})
	//fmt.Println(resp)
}
