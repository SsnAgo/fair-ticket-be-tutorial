package api

/*
Project
*/
type ProjectCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"imageUrl" binding:"required"`
	TotalSupply int64  `json:"totalSupply" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
	// 这里对应地新增两个参数 signature和message
	Signature string `json:"signature" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

type ProjectChangeStatusRequest struct {
	ProjectOnchainID int64  `json:"projectOnchainId" binding:"required"`
	Owner            string `json:"owner" binding:"required"`
}

type ProjectInfoRequest struct {
	// 通过get传参这里应当是form
	ProjectID int64 `form:"projectId" binding:"required"`
}

/*
Participant
*/
type ParticipateRequest struct {
	Name             string `json:"name" binding:"required"`
	Address          string `json:"address" binding:"required"`
	LuckyNum         int64  `json:"luckyNum" binding:"required"`
	ProjectID        int64  `json:"projectId" binding:"required"`
	ProjectOnchainID int64  `json:"projectOnchainId" binding:"required"`
}

type ParticipatedProjectsRequest struct {
	Address string `form:"address" binding:"required"`
}

type ParticipateInfoRequest struct {
	ProjectOnchainID int64  `form:"projectOnchainId" binding:"required"`
	Address          string `form:"address" binding:"required"`
}

type VerifyMerkleProofRequest struct {
	ProjectOnchainID int64    `json:"projectOnchainId" binding:"required"`
	Address          string   `json:"address" binding:"required"`
	Proof            []string `json:"proof" binding:"required"`
}
