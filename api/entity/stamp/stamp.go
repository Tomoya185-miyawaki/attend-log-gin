/*
エンティティ（出退勤）
*/
package stamp

type ID uint
type EmployeeID uint
type Status uint8
type StampStartDate string
type StampEndDate string

type Stamp struct {
	ID             ID             `json:"id"`
	EmployeeID     EmployeeID     `json:"employee_id"`
	Status         Status         `json:"status"`
	StampStartDate StampStartDate `json:"stamp_start_date"`
	StampEndDate   StampEndDate   `json:"stamp_end_date"`
}

type Stamps []Stamp
