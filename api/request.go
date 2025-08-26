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
}

type ProjectChangeStatusRequest struct {
	ProjectOnchainID int64  `json:"projectOnchainID" binding:"required"`
	Owner            string `json:"owner" binding:"required"`
}

type ProjectInfoRequest struct {
	// 通过get传参这里应当是form
	ProjectID int64 `form:"projectID" binding:"required"`
}

/*
Participant
*/

type ParticipateRequest struct {
	Name             string `json:"name" binding:"required"`
	Address          string `json:"address" binding:"required"`
	LuckyNum         int64  `json:"luckyNum" binding:"required"`
	ProjectID        int64  `json:"projectID" binding:"required"`
	ProjectOnchainID int64  `json:"projectOnchainID" binding:"required"`
}

type ParticipatedProjectsRequest struct {
	Address string `form:"address" binding:"required"`
}

type ParticipateInfoRequest struct {
	ProjectOnchainID int64  `form:"projectOnchainID" binding:"required"`
	Address          string `form:"address" binding:"required"`
}

type VerifyMerkleProofRequest struct {
	ProjectOnchainID int64    `json:"projectOnchainID" binding:"required"`
	Address          string   `json:"address" binding:"required"`
	Proof            []string `json:"proof" binding:"required"`
}
