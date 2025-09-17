package dto

type ResolveJacketsResponse struct {
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

type CreateJacketsRequest struct {
	Name             string `json:"name"`
	PhotoSizeChart   string `json:"photoSizeChart"`
	PhotoFrontJacket string `json:"photoFrontJacket"`
	PhotoBackJacket  string `json:"photoBackJacket"`
	BasePrice        int    `json:"basePrice"`
	ExtraPrice       int    `json:"extraPrice"`
}

type CreateJacketsResponse struct {
	Id               int64  `db:"id"  json:"id"`
	Name             string `json:"name"`
	PhotoSizeChart   string `json:"photoSizeChart"`
	PhotoFrontJacket string `json:"photoFrontJacket"`
	PhotoBackJacket  string `json:"photoBackJacket"`
	BasePrice        int    `json:"basePrice"`
	ExtraPrice       int    `json:"extraPrice"`
}

type UpdateJacketsRequest struct {
	Name             *string `json:"name,omitempty" db:"name"`
	PhotoSizeChart   *string `json:"photoSizeChart,omitempty" db:"photo_size_chart"`
	PhotoFrontJacket *string `json:"photoFrontJacket,omitempty" db:"photo_front_jacket"`
	PhotoBackJacket  *string `json:"photoBackJacket,omitempty" db:"photo_back_jacket"`
	BasePrice        *int    `json:"basePrice,omitempty" db:"base_price"`
	ExtraPrice       *int    `json:"extraPrice,omitempty" db:"extra_price"`
}

type DeleteJacketsResponse struct {
	Deleted bool `json:"deleted"`
}
