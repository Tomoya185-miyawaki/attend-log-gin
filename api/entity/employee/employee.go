/*
エンティティ（従業員）
*/
package employee

type ID uint
type Name string
type HourlyWage uint

type Employee struct {
	ID         ID         `json:"id"`
	Name       Name       `json:"name"`
	HourlyWage HourlyWage `json:"hourly_wage"`
}

type Employees []Employee
