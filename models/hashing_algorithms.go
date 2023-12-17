package models

type HashingAlgorithms struct {
	ID            int64  `json:"hashing_algorithms_id" gorm:"column:hashing_algorithms_id"`
	AlgorithmName string `json:"algorithm_name"`
}

// TableName gives table name of model
func (u HashingAlgorithms) TableName() string {
	return "hashing_algorithms"
}
