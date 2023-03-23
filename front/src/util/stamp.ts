import { StampsByEmployeeId } from '@/types/stamp'
import { StampStatus, StampAttendRestStatus } from '@/enums/stamp'

/**
 * 出勤状況を表す情報を返す
 * @return {array} 出勤状況を情報テキスト（1番目は表示するテキスト、2番目はボタンのテキスト、3番目は登録するstatus）
 */
export function getAttendRestStatus(stamps: StampsByEmployeeId[]) {
  let text = '出勤前'
  let buttonText = '出勤'
  let status = StampStatus.Attend
  if (stamps.length > 0) {
    stamps.sort((a, b) => a.status - b.status)
    for (let i = 0; i < stamps.length; i++) {
      if (stamps[i].status === StampAttendRestStatus.AttendLeaving && stamps[i].stamp_end_date !== '') {
        text = '退勤済み'
        buttonText = ''
        return [text, buttonText, status]
      }
      // 出勤済みの場合（後続処理を優先するためreturnしない）
      if (stamps[i].status === StampAttendRestStatus.AttendLeaving && stamps[i].stamp_start_date !== '') {
        text = '出勤済み'
        buttonText = '休憩開始'
        status = StampStatus.RestStart
      }
      // 休憩終了している場合
      if (stamps[i].status === StampAttendRestStatus.Rest && stamps[i].stamp_end_date !== '') {
        text = '休憩終了済み'
        buttonText = '退勤'
        status = StampStatus.Leaving
        return [text, buttonText, status]
      }
      // 休憩開始している場合
      if (stamps[i].status === StampAttendRestStatus.Rest && stamps[i].stamp_start_date !== '') {
        text = '休憩開始済み'
        buttonText = '休憩終了'
        status = StampStatus.RestEnd
        return [text, buttonText, status]
      }
    }
  }
  return [text, buttonText, status]
}