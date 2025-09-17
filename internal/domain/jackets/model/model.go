package model

type Jacket struct {
	TotalData        int    `db:"total_data" json:"-"`
	Id               int    `db:"id"  json:"id"`
	Name             string `db:"name" json:"name"`
	PhotoSizeChart   string `db:"photo_size_chart" json:"photoSizeChart"`
	PhotoFrontJacket string `db:"photo_front_jacket" json:"photoFrontJacket"`
	PhotoBackJacket  string `db:"photo_back_jacket" json:"photoBackJacket"`
	BasePrice        int    `db:"base_price" json:"basePrice"`
	ExtraPrice       int    `db:"extra_price" json:"extraPrice"`
	CreatedAt        string `db:"created_at" json:"createdAt"`
	UpdatedAt        string `db:"updated_at" json:"updatedAt"`
}
