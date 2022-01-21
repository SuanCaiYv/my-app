package entity

type UserInfo struct {
	Avatar    string `bson:"avatar" json:"avatar"`
	Nickname  string `bson:"nickname" json:"nickname"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" bson:"phone"`
	WeChat    string `bson:"we_chat" json:"we_chat"`
	QQ        string `bson:"qq" json:"qq"`
	Github    string `bson:"github" json:"github"`
	Location  string `bson:"location" json:"location"`
	Signature string `bson:"signature" json:"signature"`
}
