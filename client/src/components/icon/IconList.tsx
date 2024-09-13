import { Icon, loadIcons } from '@iconify/react'
import icons from '@/assets/icons/mdi.json'
import { useEffect } from 'react'

const IconList = () => {
  const collection = 'mdi'
  useEffect(() => {
    loadIcons(icons)
  }, [])
  return (
    <div className="grid grid-cols-8">
      {icons.map((icon) => {
        return <Icon icon={`${collection}:${icon}`} style={{ fontSize: '24px' }} />
      })}
    </div>
  )
}

export default IconList
