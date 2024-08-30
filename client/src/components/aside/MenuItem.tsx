import React from 'react'
import { Icon } from '@iconify/react'
import { AnimatePresence, motion } from 'framer-motion'
import { Link } from 'react-router-dom'

interface MenuItemProps {
  item: {
    id: number
    name: string
    icon: string
    path: string
  }
  collapse?: boolean
}

const MenuItem: React.FC<MenuItemProps> = ({ item, collapse }) => {
  return (
    <Link to={item.path} className={` ${!collapse ? 'pl-10' : 'px-4'} flex items-center py-4  gap-3 rounded cursor-pointer text-black dark:text-white hover:bg-[#ECF5FF] dark:hover:bg-[#343435]`}>
      <Icon icon={'carbon:' + item.icon} className="text-xl" />
      <AnimatePresence>
        <motion.span
          style={{ minWidth: '100px' }}
          initial={{ scale: 0.8, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          exit={{ scale: 0.8, opacity: 0 }}
          transition={{ duration: 0.6 }}
          className="text-[14px]"
        >
          {item.name}
        </motion.span>
      </AnimatePresence>
    </Link>
  )
}

export default MenuItem
