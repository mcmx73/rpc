package rpc

import "fmt"

func NewUrl(host, port, uri string) *RpcURL {
	return &RpcURL{
		scheme: "http",
		host:   host,
		port:   port,
		uri:    uri,
	}
}

func NewSslUrl(host, port, uri string) *RpcURL {
	return &RpcURL{
		scheme: "https",
		host:   host,
		port:   port,
		uri:    uri,
	}
}

type RpcURL struct {
	scheme string
	host   string
	uri    string
	port   string
}

func (u *RpcURL) String() string {
	return fmt.Sprintf("%s:%s/%s", u.scheme, u.host, u.uri)
}
