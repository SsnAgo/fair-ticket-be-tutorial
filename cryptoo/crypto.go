package cryptoo

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
这里需要解释一下为什么要对message进行转换
我们传入的参加签名的原始消息是 "create project"
但是我们在前端使用钱包进行evm签名时 实际上构造的消息是以如下格式
即 evm消息签名前缀"\x19Ethereum Signed Message:\n" + message的长度 + message
这一段对原始消息的拼接  才是真正需要参与签名的消息
因此我们从签名中恢复地址时，参与运算的消息也应该是这串拼接后的消息
*/
func buildMessageBytes(message string) []byte {
	prefix := "\x19Ethereum Signed Message:\n"
	length := strconv.Itoa(len(message))
	messageBytes := []byte(prefix + length + message)
	return messageBytes
}

func VerifySignature(message string, signature string, address string) bool {
	// 构造参与签名的消息
	messageBytes := buildMessageBytes(message)
	// 计算消息的hash
	messageHash := crypto.Keccak256Hash(messageBytes)
	// 将签名转换为十六进制数组
	signatureBytes := common.FromHex(signature)
	// 如果签名是使用非EIP-155标准，需要将恢复ID减去27 使最后一个字节保持0或者1
	if signatureBytes[64] >= 27 {
		signatureBytes[64] -= 27
	}
	// 从签名和消息hash中恢复出公钥
	sigPublicKey, err := crypto.SigToPub(messageHash.Bytes(), signatureBytes)
	if err != nil {
		return false
	}
	// 从公钥恢复地址
	recoveredAddress := crypto.PubkeyToAddress(*sigPublicKey)
	// 比较恢复的地址与提供的地址
	return common.HexToAddress(address) == recoveredAddress
}
