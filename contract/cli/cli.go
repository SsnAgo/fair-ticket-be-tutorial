package cli

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 简单清晰起见，直接在代码里面编写相关配置
const (
	// 注意 这里的协议要使用ws而不是http，建立长连接，确保能够接收到合约的触发事件。
	RPC_URL            = "ws://127.0.0.1:8545"
	ADMIN_PRIVATEK_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	GasLimit           = 30_000_000
)

type EthClient struct {
	*ethclient.Client
}

var once sync.Once
var clientInstance *EthClient

func GetClient() *EthClient {
	once.Do(func() {
		client, err := ethclient.Dial(RPC_URL)
		if err != nil {
			log.Fatal(err)
		}
		clientInstance = &EthClient{
			Client: client,
		}
	})
	return clientInstance
}

// 声明一个通用的交易设置，在后面调用合约，写数据时候会使用到
func (c *EthClient) TransactorOpt() *bind.TransactOpts {
	// common.FromHex 会帮助我们把0x开头的字符串转换为字节数组
	keyBytes := common.FromHex(ADMIN_PRIVATEK_KEY)
	// 注意这里不能使用string(keyBytes)
	// 因为keyBytes每个字符是16进制表示 每个字符占4位。
	// 所以需要使用Bytes2Hex 把字节数组转换为16进制字符串 如果直接用 string(),会按照一个字节8位来算一个字符。
	// 在后续进行 [32]byte 转换成字符串类型时 也同样需要注意。
	key := common.Bytes2Hex(keyBytes)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	fromAddr := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))
	nonce, _ := c.Client.PendingNonceAt(context.Background(), fromAddr)
	gasPrice, _ := c.Client.SuggestGasPrice(context.Background())
	chainId, err := c.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	opt, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	opt.Nonce = big.NewInt(int64(nonce))
	opt.GasLimit = GasLimit
	opt.GasPrice = gasPrice
	return opt
}

// call 指的是调用合约的只读函数 不会发起交易 不需要消耗gas 不需要签名
func (c *EthClient) CallOpt() *bind.CallOpts {
	return &bind.CallOpts{}
}
