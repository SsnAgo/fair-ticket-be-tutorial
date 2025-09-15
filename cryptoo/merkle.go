package cryptoo

import (
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Leave = common.Hash

func hash(data []byte) Leave {
	h := crypto.Keccak256(data)
	return Leave(h[:32])
}

// commutativeHash 对两个Leave进行可交换哈希计算
// 确保无论输入顺序如何，相同的两个值总是产生相同的哈希
func commutativeHash(a, b Leave) Leave {
	if a.Big().Cmp(b.Big()) > 0 {
		return hash(append(b.Bytes(), a.Bytes()...))
	}
	return hash(append(a.Bytes(), b.Bytes()...))
}

// BuildMerkleTree 构建Merkle树并返回根哈希和每个地址的证明路径
func BuildMerkleTree(leaves []common.Address) (root []byte, proofs map[common.Address][][]byte) {
	n := len(leaves)
	if n == 0 {
		return nil, nil
	}

	// 建立叶子结点的hash与原值的映射
	hash2addr := map[Leave]common.Address{}
	hashLeaves := make([]Leave, len(leaves))
	for i, leaf := range leaves {
		hash2addr[hash(leaf.Bytes())] = leaf
		hashLeaves[i] = hash(leaf.Bytes())
	}
	// 对hashLeaves进行排序，确保树的构建是确定性的
	sort.Slice(hashLeaves, func(i, j int) bool {
		return hashLeaves[i].Big().Cmp(hashLeaves[j].Big()) < 0
	})

	// levels 存储Merkle树的所有层级节点
	levels := [][]Leave{}
	// currentLayer 存储当前遍历的层级
	currentLayer := []Leave{}
	for _, leaf := range hashLeaves {
		currentLayer = append(currentLayer, leaf)
	}

	levels = append(levels, currentLayer)

	// 自底向上构建Merkle树
	for len(currentLayer) > 1 {
		var next []Leave
		for i := 0; i < len(currentLayer); i += 2 {
			if i+1 == len(currentLayer) {
				// 奇数个，最后一个节点复制自身
				next = append(next, commutativeHash(currentLayer[i], currentLayer[i]))
			} else {
				// 正常配对相邻的两个节点
				next = append(next, commutativeHash(currentLayer[i], currentLayer[i+1]))
			}
		}
		currentLayer = next
		levels = append(levels, currentLayer)
	}

	// 根哈希是最顶层的唯一节点
	root = levels[len(levels)-1][0].Bytes()

	// 为每个哈希叶子节点生成Merkle证明路径
	hashProofs := map[Leave][][]byte{}
	for i := range hashLeaves {
		idx := i
		var path [][]byte
		for l := 0; l < len(levels)-1; l++ {
			layer := levels[l]
			sibling := idx ^ 1 // 与当前 index 异或得出兄弟节点
			if sibling < len(layer) {
				path = append(path, layer[sibling].Bytes())
			} else {
				path = append(path, layer[idx].Bytes())
			}
			// 移动到父节点
			idx = idx / 2
		}
		// 存储叶子节点hash对应的proof路径
		hashProofs[hashLeaves[i]] = path
	}
	// 将叶子节点hash对应的proof路径转换为地址对应的proof路径
	proofs = map[common.Address][][]byte{}
	for h, p := range hashProofs {
		proofs[hash2addr[h]] = p
	}
	// 返回Merkle tree的root以及每个地址对应的proof
	return root, proofs
}
