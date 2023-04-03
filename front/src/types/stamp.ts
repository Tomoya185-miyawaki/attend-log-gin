import { StampAttendRestStatus, StampStatus } from '@/enums/stamp'

export type StampList = {
  attend_date: string
  leaving_date: string
  rest_date: string
  working_date: string
}

export type StampsByEmployeeId = {
  id: number
  employee_id: number
  status: StampAttendRestStatus
  stamp_start_date: string
  stamp_start: string
  stamp_end_date?: string
  stamp_end?: string
}

export type StampDetail = {
  id: number
  status: StampStatus
  stamp_start_date: string
  stamp_end_date?: string
}

export type CreateStampData = {
  employeeId: number
  status: StampStatus
  today: string
}

export type UpdateStampData = {
  id: number
  status: StampAttendRestStatus
  stamp_start_date: Date|null
  stamp_end_date: Date|null
}

export type AdminCreateStampData = {
  employeeId: number
  status: StampAttendRestStatus
  stamp_start_date: Date
  stamp_end_date: Date
}