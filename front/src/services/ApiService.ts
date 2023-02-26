import http from '@/util/http'
import { LoginFormData, PasswordResetFormData, EmployeeFormData } from '@/types/auth'
import {
  LoginResponse,
  LogoutResponse,
  PasswordResetResponse,
  GetEmployeesByIdRes,
  GetEmployeesByPaginateRes,
  CreateEmployeeRes,
  GetStampsByPaginateRes,
  GetStampDetailRes
} from '@/types/api/response'

class ApiService {
  login(formData: LoginFormData): Promise<LoginResponse> {
    return http.post('/api/admin/login', formData)
  }

  logout(): Promise<LogoutResponse> {
    return http.post('/api/admin/logout')
  }

  passwordReset(formData: PasswordResetFormData): Promise<PasswordResetResponse> {
    return http.post('/api/admin/password-reset', formData)
  }

  async getEmployeesById(employeeId: string): Promise<GetEmployeesByIdRes> {
    const response = await http.get(`/api/employee/${employeeId}`)
    return response.data
  }

  async getEmployeesByPaginate(page: number): Promise<GetEmployeesByPaginateRes> {
    const response = await http.get(`/api/employee?page=${page}`)
    return response.data
  }

  createEmployee(formData: EmployeeFormData): Promise<CreateEmployeeRes> {
    return http.post('/api/employee/create', formData)
  }

  updateEmployee(formData: EmployeeFormData, id: string): Promise<void> {
    return http.patch(`/api/employee/${id}`, formData)
  }

  deleteEmployee(id: string): Promise<void> {
    return http.delete(`/api/employee/${id}`)
  }

  async getStampsByPaginate(today: string, page: number): Promise<GetStampsByPaginateRes> {
    const response = await http.get(`/api/stamp?today=${today}&page=${page}`)
    return response.data
  }

  async getStampDetail(employeeId: string): Promise<GetStampDetailRes> {
    const response = await http.get(`/api/stamp/${employeeId}`)
    return response.data
  }
}

export default new ApiService()