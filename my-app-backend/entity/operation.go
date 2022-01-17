package entity

import "time"

type OperationRecord struct {
	Id            string    `bson:"_id" json:"-"`
	OperationName string    `bson:"operation_name" json:"operationName"`
	OperationDesc string    `bson:"operation_desc" json:"operationDesc"`
	Operator      string    `bson:"operator" json:"operator"`
	Target        string    `bson:"target" json:"target"`
	Ip            string    `bson:"ip" json:"ip"`
	ClientInfo    string    `bson:"client_info" json:"clientInfo"`
	Time          time.Time `bson:"time" json:"time"`
}
