import React from 'react'
import { Icon } from '@iconify/react'
import cls from 'classnames'

type SvgIconProps = {
  collection?: string
  name: string
  color?: string
  size?: string
}

const GraIcon: React.FC<SvgIconProps> = ({ collection = 'mdi', name, color = '#333', size = '3xl' }) => {
  return <Icon icon={`${collection}:${name}`} className={cls(`text-${color}-400 text-${size}`)} />
}
export default GraIcon
