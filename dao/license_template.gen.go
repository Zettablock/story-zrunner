// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameLicenseTemplate = "license_template"

// LicenseTemplate mapped from table <license_template>
type LicenseTemplate struct {
	BlockNumber int64  `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   int64  `gorm:"column:block_time;not null" json:"block_time"`
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	MetadataURI string `gorm:"column:metadata_uri" json:"metadata_uri"`
}

// TableName LicenseTemplate's table name
func (*LicenseTemplate) TableName() string {
	return TableNameLicenseTemplate
}
