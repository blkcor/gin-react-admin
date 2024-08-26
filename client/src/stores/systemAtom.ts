import { atom } from 'jotai'
import { SystemConfig } from './types'

export const systemAtom = atom<SystemConfig>({
  collapsed: false
})
