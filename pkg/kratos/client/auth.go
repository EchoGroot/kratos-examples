package client

import "context"

func NewKeyAuthPerRPCCredentials(apikey string) *keyAuthPerRPCCredentials {
	return &keyAuthPerRPCCredentials{
		apikey: apikey,
	}
}

type keyAuthPerRPCCredentials struct {
	apikey string
}

func (k *keyAuthPerRPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"apikey": k.apikey,
	}, nil
}

// 是否使用 TLS 安全加密
func (k *keyAuthPerRPCCredentials) RequireTransportSecurity() bool {
	return false
}
