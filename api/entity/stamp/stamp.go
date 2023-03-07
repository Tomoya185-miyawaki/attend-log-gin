/*
エンティティ（出退勤）
*/
package stamp

type ID uint
type EmployeeID uint
type Status uint
type StampStartDate string
type StampStart string
type StampEndDate string
type StampEnd string

type Stamp struct {
	ID             ID             `json:"id"`
	EmployeeID     EmployeeID     `json:"employee_id"`
	Status         Status         `json:"status"`
	StampStartDate StampStartDate `json:"stamp_start_date"`
	StampStart     StampStart     `json:"stamp_start"`
	StampEndDate   StampEndDate   `json:"stamp_end_date"`
	StampEnd       StampEnd       `json:"stamp_end"`
}

type Stamps []Stamp
