import { StampAttendRestStatus, StampStatus } from '@/enums/stamp'
import { StampDetail } from '@/types/stamp'

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

export function setEvent(stamp: StampDetail) {
  return {
    id: stamp.id,
    title: getTitle(stamp.status),
    start: stamp.stamp_start_date,
    end: stamp.stamp_end_date,
    backgroundColor: getColor(stamp.status),
    borderColor: getColor(stamp.status),
    editable: true
  }
}

export function getStatusByTitle(title: string) {
  if (title === '勤務') {
    return StampAttendRestStatus.AttendLeaving
  }
  if (title === '休憩') {
    return StampAttendRestStatus.Rest
  }
  return '';
}