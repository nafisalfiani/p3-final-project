package entity

type Wishlist struct {
	Id              string   `json:"id"`
	CategoryName    string   `json:"category_name"`
	RegionName      string   `json:"region_name"`
	SubscribedUsers []string `json:"subscribed_users"`
}

type WishlistGetRequest struct {
	Id string `uri:"id"`
}

type WishlistSubscribeRequest struct {
	CategoryName string `form:"category_name"`
	RegionName   string `form:"region_name"`
}
