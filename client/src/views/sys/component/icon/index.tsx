import { useEffect, useState } from 'react'
import { loadIcon, IconifyIcon } from '@iconify/react'
import icons from '@/assets/icons/mdi.json'
import { useClipboard } from '@/hooks/useClipBorad'
import { toast } from 'react-toastify'
import { loadIconAsync } from '@/utils/icon'
import GraIcon from '@/components/icon/GraIcon'
import * as Switch from '@radix-ui/react-switch'
import './icon.css'

const IconView = () => {
  const collection = 'mdi'
  const { copy } = useClipboard()
  const [displayedIcons, setDisplayedIcons] = useState<string[]>([]) // 当前显示的图标
  const [remainingIcons, setRemainingIcons] = useState<string[]>([]) // 剩余图标
  const [allIconsLoaded, setAllIconsLoaded] = useState(false) // 是否加载了全部图标
  const INITIAL_LOAD_COUNT = 24 // 初始加载的图标数量
  const [componentOrSvg, setComponentOrSvg] = useState<boolean>(true)

  useEffect(() => {
    // 初始化时只加载前 N 个图标
    setDisplayedIcons(icons.slice(0, INITIAL_LOAD_COUNT))
    setRemainingIcons(icons.slice(INITIAL_LOAD_COUNT)) // 存储剩余图标
  }, [])

  useEffect(() => {
    async function loadIcons() {
      await loadIconAsync(icons.slice(INITIAL_LOAD_COUNT))
    }
    loadIcons()
  }, [])

  // 将IconifyIcon对象转换为svg字符串
  function svgObjToString(svg: Required<IconifyIcon>): string {
    const { body, hFlip, height, left, rotate, top, vFlip, width } = svg
    return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="${left} ${top} ${width} ${height}" width="${width}" height="${height}" transform="rotate(${rotate} ${width / 2} ${height / 2}) ${
      hFlip ? `scale(-1, 1) translate(-${width} 0)` : ''
    } ${vFlip ? `scale(1, -1) translate(0 -${height})` : ''}">${body}</svg>`
  }

  const handleCopyIcon = (icon: string) => {
    loadIcon(collection + ':' + icon)
      .then((svg) => {
        let content = svgObjToString(svg)
        if (componentOrSvg) {
          content = `<GraIcon name="${icon}" />`
        }
        copy(content)

        toast.success(componentOrSvg ? '已复制组件图标' : '已复制svg图标', {
          autoClose: 800,
        })
      })
      .catch((err) => {
        toast.error('复制svg错误', err)
      })
  }

  const loadAllIcons = () => {
    // 加载剩余图标并将其显示
    setDisplayedIcons((prevIcons) => [...prevIcons, ...remainingIcons])
    setRemainingIcons([]) // 清空剩余图标
    setAllIconsLoaded(true) // 设置标志为已加载全部
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6 text-center">Icon展馆</h1>
      <div className="flex gap-2 mb-4">
        <span>复制SVG</span>
        <Switch.Root className="SwitchRoot" checked={componentOrSvg} onCheckedChange={() => setComponentOrSvg((prev) => !prev)}>
          <Switch.Thumb className="SwitchThumb" />
        </Switch.Root>
        <span>复制组件</span>
      </div>
      <div className="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-6 lg:grid-cols-8 gap-4">
        {displayedIcons.map((icon, index) => (
          <div
            onClick={() => handleCopyIcon(icon)}
            key={index}
            className="flex flex-col items-center justify-center p-2 border rounded hover:shadow-md transition-shadow cursor-pointer hover:bg-gray-400/30"
          >
            <GraIcon name={icon} />
            <span className="text-xs text-center break-all">{icon}</span>
          </div>
        ))}
      </div>
      {/* 如果还没有加载全部图标，显示加载按钮 */}
      {!allIconsLoaded && (
        <div className="text-center mt-6">
          <button onClick={loadAllIcons} className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-700">
            加载全部图标
          </button>
        </div>
      )}
    </div>
  )
}

export default IconView
