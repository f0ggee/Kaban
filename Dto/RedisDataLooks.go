package Dto

type RedisDataLooksLike struct {
	AesKey    []byte `redis:"aes"`
	PlainText []byte `redis:"plaintext"`
	Signature []byte `redis:"signature"`
}
