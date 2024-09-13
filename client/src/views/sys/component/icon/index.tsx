import { useEffect } from 'react'
import { Icon, loadIcons } from '@iconify/react'
import icons from '@/assets/icons/mdi.json'

const IconView = () => {
  const collection = 'mdi'

  useEffect(() => {
    loadIcons(icons)
  }, [])

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6 text-center">Icon Gallery</h1>
      <div className="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-6 lg:grid-cols-8 gap-4">
        {icons.map((icon, index) => (
          <div key={index} className="flex flex-col items-center justify-center p-2 border rounded hover:shadow-md transition-shadow">
            <Icon icon={`${collection}:${icon}`} className="text-3xl mb-2" />
            <span className="text-xs text-center break-all">{icon}</span>
          </div>
        ))}
      </div>
    </div>
  )
}

export default IconView
