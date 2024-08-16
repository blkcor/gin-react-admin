import { atom } from 'jotai'
import { UserInfo } from '@/apis/types'

/**用户信息存储 */
export const userAtom = atom({
  userInfo: {} as UserInfo,
  token: '',
})
