package model

// Campaign basic campaign info
type Campaign struct {
	CampaignUin          string `json:"campaignUin" xorm:"'CAMPAIGNUIN'" example:"14579"`
	CampaignName         string `json:"campaignName" xorm:"'CAMPAIGNNAME'" example:"Survery Weekly"`
	Creator              string `json:"creator" xorm:"'CREATOR'" example:"shiming.xue@hgc.com.hk"`
	Status               int    `json:"status" xorm:"'STATUS'" example:"10"`
	IntimateCampaignType int    `json:"intimateCampaignType" xorm:"'INTIMATECAMPAIGNTYPE'" example:"1"`
	CreateTime           string `json:"createTime" xorm:"'CREATETIME'" example:"2021/11/24"`
	LastModifyTime       string `json:"lastModifyTime" xorm:"'LASTMODIFYTIME'" example:"2021/11/24"`
	ActualStartTime      string `json:"actualStartTime" xorm:"'ACTUALSTARTTIME'" exampe:""`
	ActualEndTime        string `json:"actualEndTime" xorm:"'ACTUALENDTIME'" example:""`
}
