package user

type User struct {
	PrivyID       string `bson:"_id" json:"privy_id"`
	Username      string `bson:"username" json:"username"`
	WalletAddress string `bson:"wallet_address" json:"wallet_address"`
}
