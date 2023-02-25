import Cookies from 'js-cookie';

/**
 * ログイン判定
 * @return {boolean} ログインしているかどうかの結果
 */
export function isAdminLoggedIn() {
  const cookie = Cookies.get('attend-log-gin-session')
  return typeof cookie !== 'undefined' ? true : false
}