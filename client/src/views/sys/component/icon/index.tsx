import { useEffect } from 'react'
import { Icon, loadIcons, loadIcon, IconifyIcon } from '@iconify/react'
import icons from '@/assets/icons/mdi.json'
import { useClipboard } from '@/hooks/useClipBorad'
import { toast } from 'react-toastify'

const IconView = () => {
  const collection = 'mdi'
  const { copy } = useClipboard()
  useEffect(() => {
    loadIcons(icons)
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
        copy(svgObjToString(svg))
        toast.success('已复制svg', {
          autoClose: 800,
        })
      })
      .catch((err) => {
        toast.error('复制svg错误', err)
      })
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6 text-center">Icon Gallery</h1>
      <div className="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-6 lg:grid-cols-8 gap-4">
        {icons.map((icon, index) => (
          <div
            onClick={() => handleCopyIcon(icon)}
            key={index}
            className="flex flex-col items-center justify-center p-2 border rounded hover:shadow-md transition-shadow cursor-pointer hover:bg-gray-400/30"
          >
            <Icon icon={`${collection}:${icon}`} className="text-3xl mb-2" />
            <span className="text-xs text-center break-all">{icon}</span>
          </div>
        ))}
      </div>
    </div>
  )
}

export default IconView
