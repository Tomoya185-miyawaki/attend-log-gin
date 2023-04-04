import { Employee } from "@/types/model"
import { StampDetail, StampList, StampsByEmployeeId } from "@/types/stamp"

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

export type GetEmployees = {
  statusCode: number,
  employees: Employee[]
}

export type GetEmployeesByIdRes = {
  employee: Employee
}

export type GetEmployeesByPaginateRes = {
  statusCode: number,
  currentPage: number
  employees: Employee[]
  lastPage: number
}

export type GetEmployeesRes = {
  employees: Employee[]
}

export type CreateEmployeeRes = {
  status: number
  message: string
}

export type UpdateEmployeeRes = {
  status: number
  message: string
}

export type DeleteEmployeeRes = {
  status: number
  message: string
}

export type GetStampsByEmployeeId = {
  status: number
  stamps: StampsByEmployeeId[]
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

export type CreateStampRes = {
  status: number
  message: string
}

export type UpdateStampRes = {
  status: number
  message: string
}

export type DeleteStampRes = {
  status: number
  message: string
}

export type AdminCreateStampRes = {
  status: number
  message: string
}