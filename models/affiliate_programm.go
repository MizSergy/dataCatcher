package models

type AffiliateProgram struct {
	ID          int    `form:"id" json:"id" bson:"_id"`
	AffiliateId int `form:"affiliate_id" json:"affiliate_id,omitempty" bson:"affiliate_id" binding:"required"`
	Title       string `form:"title" json:"title,omitempty" bson:"title" binding:"required"`
	IsArchiving bool   `form:"is_archiving" json:"is_archiving,omitempty" bson:"is_archiving"`
	CurrencyID  int    `form:"currency_id" json:"currency_id,omitempty" bson:"currency_id"  binding:"required"`
}

func (m AffiliateProgram) GetID() int {
	return m.ID
}
