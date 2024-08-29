import { useState } from "react"

export const useFullScreen = () => {
  const [isFullScreen, setIsFullScreen] = useState(false)

  function toggleFullScreen() {
    if (!document.fullscreenElement) {
      document.documentElement.requestFullscreen()
      setIsFullScreen(true)
    } else if (document.exitFullscreen) {
      document.exitFullscreen()
      setIsFullScreen(false)
    }
  }

  return { isFullScreen, toggleFullScreen }
}
