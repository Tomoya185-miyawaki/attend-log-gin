/*
エンティティ（出退勤）
*/
package stamp

type AttendDate string
type LeavingDate string
type RestDate string
type WorkingDate string

type StampDate struct {
	AttendDate  AttendDate  `json:"attend_date"`
	LeavingDate LeavingDate `json:"leaving_date"`
	RestDate    RestDate    `json:"rest_date"`
	WorkingDate WorkingDate `json:"working_date"`
}

type StampsDate []StampDate

func NewStampDate() *StampDate {
	return &StampDate{}
}

func (stampDate *StampDate) GetAttendDate() AttendDate {
	return stampDate.AttendDate
}

func (stampDate *StampDate) GetLeavingDate() LeavingDate {
	return stampDate.LeavingDate
}

func (stampDate *StampDate) GetRestDate() RestDate {
	return stampDate.RestDate
}

func (stampDate *StampDate) GetWorkingDate() WorkingDate {
	return stampDate.WorkingDate
}

func (stampDate *StampDate) SetAttendDate(attendDate AttendDate) {
	stampDate.AttendDate = attendDate
}

func (stampDate *StampDate) SetLeavingDate(leavingDate LeavingDate) {
	stampDate.LeavingDate = leavingDate
}

func (stampDate *StampDate) SetRestDate(restDate RestDate) {
	stampDate.RestDate = restDate
}

func (stampDate *StampDate) SetWorkingDate(workingDate WorkingDate) {
	stampDate.WorkingDate = workingDate
}
