package test

import (
	"fair-ticket-be-tutorial/service"
	"strings"
	"testing"
	"time"

	// 这里引入两个依赖 方便检验测试结果
	// go get github.com/stretchr/testify/assert
	// go get github.com/stretchr/testify/require
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFairTicketIntegration(t *testing.T) {
	// 初始化服务
	projectService := service.NewProjectService()
	participantService := service.NewParticipantService()
	// 启动监听服务 这里使用go协程 启动多个监听服务 确保能够接收到合约的触发事件
	{
		eventService := service.NewEventService()
		go eventService.HandleProjectCreated()
		go eventService.HandleProjectStarted()
		go eventService.HandleProjectFinished()
		go eventService.HandleMagicNumberPublished()
	}
	// 由于上面是启动多个协程 所以这里sleep确保 监听启动完成
	time.Sleep(2 * time.Second)

	// 测试数据准备
	// owner这里设置为anvil的默认Accounts(0)
	projectOwner := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	projectName := "测试公平抽票项目"
	projectDescription := "这是一个测试项目，用于验证公平抽票系统的完整流程"
	projectImageURL := "https://picsum.photos/seed/1/400/300"
	totalSupply := int64(2) // 设置为2，这样我们可以测试部分中奖的情况

	// 参与者测试数据
	participants := []struct {
		name     string
		address  string
		luckyNum int64
	}{
		// 参与者地址设置为anvil的默认Accounts(1-4)
		{"参与者1", "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", 100},
		{"参与者2", "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", 200},
		{"参与者3", "0x90F79bf6EB2c4f870365E785982E1f101E93b906", 300},
		{"参与者4", "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65", 400},
	}

	t.Run("完整的公平抽票流程测试", func(t *testing.T) {
		// Step 1: 创建项目
		t.Log("📝 步骤1: 创建项目")
		createTx, err := projectService.Create(
			projectName,
			projectDescription,
			projectImageURL,
			totalSupply,
			projectOwner,
		)
		require.NoError(t, err, "创建项目应该成功")
		require.NotNil(t, createTx, "创建项目的交易不应为空")
		t.Logf("✅ 项目创建成功，交易哈希: %s", createTx.Hash().Hex())

		// 模拟延迟  等待数据上链（在实际环境中，这里需要等待区块链确认和事件监听）
		// 在上一步创建项目成功后，我们的事件监听HandleProjectCreated会监听到合约中的ProjectCreated事件，
		// 会将db里面对应的记录进行更新 project_onchain_id
		time.Sleep(1 * time.Second)

		// 模拟从事件中获取项目ID（在实际环境中，这会由EventService处理）
		// 我们这里是部署一个新的合约，初始化project_id是从 1 开始，因此我们在这里直接写 1
		// 可以根据你合约中实际上的project_id来设置
		// 假设我们刚部署合约 如果上一步上链成功 那么获取到的projectOnchainID就是1
		projectOnchainID := int64(1)

		// Step 2: 开始项目
		t.Log("🚀 步骤2: 开始项目")
		startTx, err := projectService.Start(projectOnchainID, projectOwner)
		require.NoError(t, err, "开始项目应该成功")
		require.NotNil(t, startTx, "开始项目的交易不应为空")
		t.Logf("✅ 项目开始成功，交易哈希: %s", startTx.Hash().Hex())

		// 模拟延迟  等待数据上链 以及监听处理ProjectStarted事件
		time.Sleep(1 * time.Second)

		// 验证项目状态已更新
		project, err := projectService.GetOneByOnchainID(projectOnchainID)
		require.NoError(t, err, "获取项目信息应该成功")
		// 开始项目后，事件监听HandleProjectStarted会监听到合约中的ProjectStarted事件，
		// 会将db里面对应的记录进行更新对应项目的status字段
		assert.Equal(t, int32(service.PROJECT_STATUS_INPROGRESS), project.Status, "项目状态应该是进行中")

		// Step 3: 参与者参与项目
		t.Log("👥 步骤3: 参与者参与项目")
		for i, participant := range participants {
			participateTx, err := participantService.Participate(
				participant.name,
				participant.address,
				participant.luckyNum,
				int64(project.ID),
				projectOnchainID,
			)
			require.NoError(t, err, "参与者%d参与项目应该成功", i+1)
			require.NotNil(t, participateTx, "参与项目的交易不应为空")
			t.Logf("✅ %s 参与成功，地址: %s, 幸运数字: %d, 交易哈希: %s",
				participant.name, participant.address, participant.luckyNum, participateTx.Hash().Hex())
		}

		// Step 4: 结束项目
		t.Log("🏁 步骤4: 结束项目")
		finishTx, err := projectService.Finish(projectOnchainID, projectOwner)
		require.NoError(t, err, "结束项目应该成功")
		require.NotNil(t, finishTx, "结束项目的交易不应为空")
		t.Logf("✅ 项目结束成功，交易哈希: %s", finishTx.Hash().Hex())

		// 模拟延迟  等待数据上链 以及监听处理ProjectFinished事件
		time.Sleep(1 * time.Second)

		// 验证项目状态已更新
		project, err = projectService.GetOneByOnchainID(projectOnchainID)
		require.NoError(t, err, "获取项目信息应该成功")
		// 结束项目后，事件监听HandleProjectFinished会监听到合约中的ProjectFinished事件，
		// 会将db里面对应的记录进行更新对应项目的status字段
		assert.Equal(t, int32(service.PROJECT_STATUS_FINISHED), project.Status, "项目状态应该是已完成")

		// Step 5: 开启抽票
		t.Log("🎲 步骤5: 开启抽票")
		lotteryTx, err := projectService.Lottery(projectOnchainID, projectOwner)
		require.NoError(t, err, "开启抽票应该成功")
		require.NotNil(t, lotteryTx, "抽票交易不应为空")
		t.Logf("✅ 抽票开启成功，交易哈希: %s", lotteryTx.Hash().Hex())

		// 模拟幸运数字发布事件处理
		time.Sleep(1 * time.Second)

		// Step 6: 中奖验证
		// 在Lottery之后，会触发MagicNumberPublished事件，
		// 会进行一系列操作 详情回顾 HandleMagicNumberPublished()
		// 会进行参与者一致性校验，MerkleTree构建，项目和中奖者相关信息更新
		t.Log("🏆 步骤6: 中奖验证")

		// 获取更新后的项目信息（应该包含抽奖结果）
		project, err = projectService.GetOneByOnchainID(projectOnchainID)
		require.NoError(t, err, "获取项目信息应该成功")

		// 检查是否有抽奖数字和Merkle根
		if project.LotteryNum > 0 {
			t.Logf("🎯 抽奖数字: %d", project.LotteryNum)
			t.Logf("🌳 Merkle根: %s", project.MerkleRoot)
		}

		// 检查中奖参与者
		allParticipants, err := participantService.GetAll(projectOnchainID)
		require.NoError(t, err, "获取所有参与者应该成功")

		var winners []string
		var losers []string

		for _, participant := range allParticipants {
			if participant.Win {
				winners = append(winners, participant.Name)
				t.Logf("🎉 中奖者: %s (地址: %s)", participant.Name, participant.Address)
				// 验证中奖者有Merkle证明
				assert.NotEmpty(t, participant.MerkleProof, "中奖者应该有Merkle证明")
				// 调用智能合约去验证
				proof := strings.Split(participant.MerkleProof, ",")
				pass, err := participantService.VerifyMerkleProof(projectOnchainID, participant.Address, proof)
				require.NoError(t, err, "验证Merkle证明应该成功")
				assert.True(t, pass, "Merkle证明应该通过")
			} else {
				losers = append(losers, participant.Name)
				t.Logf("😔 未中奖: %s (地址: %s)", participant.Name, participant.Address)
			}
		}
		// 等待5秒再退出
		time.Sleep(5 * time.Second)
	})
}
