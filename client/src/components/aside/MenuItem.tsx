import React from 'react'
import { Icon } from '@iconify/react'

interface MenuItemProps {
  item: {
    id: number
    name: string
    icon: string
  }
}

const MenuItem: React.FC<MenuItemProps> = ({ item }) => {
  return (
    <div className="flex items-center p-2 gap-2 rounded cursor-pointer text-[1em]">
      <Icon icon={'carbon:' + item.icon} width={20} height={20} />

      <span>{item.name}</span>
    </div>
  )
}

export default MenuItem
