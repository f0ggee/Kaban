package Dto

type GrpcPacket struct {
	AesKeyData []byte `json:"aes_key_data"`
	CipherData []byte `json:"cipher_data"`
}
