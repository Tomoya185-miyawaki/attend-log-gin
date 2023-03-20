import { StampStatus } from '@/enums/stamp'

export function getTitle(status: StampStatus) {
  if (status === StampStatus.RestStart || status === StampStatus.RestEnd) {
    return '休憩'
  }
  return '勤務'
}

export function getColor(status: StampStatus) {
  if (status === StampStatus.RestStart || status === StampStatus.RestEnd) {
    return '#a5f1c0'
  }
  return '#3789d8'
}