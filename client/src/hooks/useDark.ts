import { useEffect, useState } from "react"


export const useDark = () => {
  const [darkMode, setDarkMode] = useState(() => {
    const isDarkMode = localStorage.getItem('darkMode') === 'true' || (!('darkMode' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
    return isDarkMode
  })

  useEffect(() => {
    // Apply dark mode class to body
    if (darkMode) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
    // Save preference tolocalStorage
    localStorage.setItem('darkMode', darkMode.toString())
  }, [darkMode])


  const toggleDarkMode = () => {
    setDarkMode(!darkMode)
  }


  return { darkMode, toggleDarkMode }
}
