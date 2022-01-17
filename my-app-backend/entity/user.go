package entity

type UserInfo struct {
	Avatar    string `bson:"avatar" json:"avatar"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" bson:"phone"`
	WeChat    string `bson:"we_chat" json:"weChat"`
	QQ        string `bson:"qq" json:"qq"`
	GitHub    string `bson:"git_hub" json:"gitHub"`
	Location  string `bson:"location" json:"location"`
	Signature string `bson:"signature" json:"signature"`
}
