package models

type Filters struct {
	Type        int      `json:"type"`
	IsArchiving bool     `json:"is_archiving"`
	Action      int      `json:"action"`
	Values      []string `json:"values"`
}

type Stream struct {
	ID          int       `form:"id" json:"id" bson:"_id"`
	CompanyID   int       `json:"company_id" bson:"company_id"`
	SourceID    int       `json:"source_id" bson:"source_id"`
	Type        int       `json:"type" bson:"type"`
	Weight      int       `json:"weight,omitempty" bson:"weight,omitempty"`
	IsArchiving bool      `json:"is_archiving,omitempty" bson:"is_archiving"`
	Filters     []Filters `json:"filters,omitempty" bson:"filters,omitempty"`

	// TDS 1
	PrelandID  int  `json:"preland_id,omitempty" bson:"preland_id,omitempty"`
	IsRedirect bool `json:"is_redirect,omitempty" bson:"is_redirect,omitempty"`

	// TDS 2
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Url         string `json:"url,omitempty" bson:"url,omitempty"`
	CurrencyID  int    `json:"currency_id,omitempty" bson:"currency_id,omitempty"`
	AffiliateID int    `json:"affiliate_id,omitempty" bson:"affiliate_id,omitempty"`
}
