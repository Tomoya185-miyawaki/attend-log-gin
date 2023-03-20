/**
 * フォーマットされた今日の日付を返す
 * @param {string} formatFirst 1番目のフォーマット文字列
 * @param {string} formatSecond 2番目のフォーマット文字列
 * @param {string|null} formatThird 3番目のフォーマット文字列
 * @return {String} フォーマットされた今日の日付
 */
export function getFormatToday(
  formatFirst: string,
  formatSecond: string,
  formatThird: string = ''
) {
  const date = new Date()
  const year = date.getFullYear()
  const month = ('0' + (date.getMonth() + 1)).slice(-2)
  const day = ('0' + (date.getDate())).slice(-2)
  return formatThird === '' ? `${year}${formatFirst}${month}${formatSecond}${day}` : `${year}${formatFirst}${month}${formatSecond}${day}${formatThird}`
}