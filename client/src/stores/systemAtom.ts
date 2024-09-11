import { atom } from 'jotai'
import { SystemConfig } from './types'

export const systemAtom = atom<SystemConfig>({
  collapsed: false
})

// 初始化暗色模式状态
const initialDarkMode =
  localStorage.getItem('darkMode') === 'true' ||
  (!('darkMode' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches);


export const darkModeAtom = atom(initialDarkMode);

export const toggleDarkModeAtom = atom(
  (get) => get(darkModeAtom),
  (get, set) => {
    const newDarkMode = !get(darkModeAtom);
    set(darkModeAtom, newDarkMode);
    localStorage.setItem('darkMode', JSON.stringify(newDarkMode));
    if (newDarkMode) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }
);

