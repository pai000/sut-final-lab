package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string //valid:"require~Name cant Blank"`
	Email      string
	CustomerID string `valid:"matches(^[L]\\d{7}$)~Please Enter L or M or H and Follow by number 7 letter, matches(^[M]\\d{7}$)~Please Enter L or M or H and Follow by number 7 letter, matches(^[H]\\d{7}$)~Please Enter L or M or H and Follow by number 7 letter"` // รหัสลูกค้าขึนต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}
