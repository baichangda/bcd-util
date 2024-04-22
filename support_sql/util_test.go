package support_sql

import (
	"testing"
	"time"
)

type User struct {
	id         int
	username   string
	createTime time.Time
}

func TestGenerateCode_InsertBatch(t *testing.T) {
	GenerateCode_InsertBatch("t_test", User{}, "id")
}

func TestGenerateCode_UpdateBatch(t *testing.T) {
	GenerateCode_UpdateBatch("t_test", User{}, "id")
}

func TestGenerateCode_Query(t *testing.T) {
	GenerateCode_Query("t_test", User{}, "id")
}
