package dao

import (
	"github.com/lib/pq"
)

const TableNameIPAsset = "ip_asset"

// IPAsset mapped from table <ip_asset>
type IPAsset struct {
	BlockNumber   int64          `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime     int64          `gorm:"column:block_time;not null" json:"block_time"`
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	IPID          string         `gorm:"column:ip_id" json:"ip_id"`
	ChainID       int64          `gorm:"column:chain_id" json:"chain_id"`
	TokenContract string         `gorm:"column:token_contract" json:"token_contract"`
	TokenID       int64          `gorm:"column:token_id" json:"token_id"`
	Metadata      []byte         `gorm:"type:jsonb;default:'{}';column:metadata" json:"metadata"`
	ChildIPIDs    pq.StringArray `gorm:"type:text[];column:child_ip_ids" json:"child_ip_ids"`
	ParentIPIDs   pq.StringArray `gorm:"type:text[];column:parent_ip_ids" json:"parent_ip_ids"`
	RootIPIDs     pq.StringArray `gorm:"type:text[];column:root_ip_ids" json:"root_ip_ids"`
	NftName       string         `gorm:"column:nft_name" json:"nft_name"`
	NftTokenURI   string         `gorm:"column:nft_token_uri" json:"nft_token_uri"`
	NftImageURL   string         `gorm:"column:nft_image_url" json:"nft_image_url"`
}

// TableName IPAsset's table name
func (*IPAsset) TableName() string {
	return TableNameIPAsset
}
