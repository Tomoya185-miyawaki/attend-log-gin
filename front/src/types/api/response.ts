import { Employee } from "@/types/model"
import { StampDetail, StampList } from "@/types/stamp"

export type LoginResponse = {
  status: number
  message: string
}

export type LogoutResponse = {
  status: number
  message: string
}

export type PasswordResetResponse = {
  status: number
  message: string
}

export type GetEmployeesByIdRes = {
  employee: Employee
}

export type GetEmployeesByPaginateRes = {
  currentPage: number
  employees: Employee[]
  lastPage: number
}

export type CreateEmployeeRes = {
  status: number
  message: string
}

export type UpdateEmployeeRes = {
  status: number
  message: string
}

export type GetStampsByPaginateRes = {
  currentPage: number
  stamps: StampList[]
  lastPage: number
  employeeIds: number[]
}

export type GetStampDetailRes = {
  employeeName: string
  stamps: StampDetail[]
}