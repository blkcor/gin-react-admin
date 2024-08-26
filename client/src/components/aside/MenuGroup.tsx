import React from 'react'
import * as Collapsible from '@radix-ui/react-collapsible'
import MenuItem from './MenuItem'
import { ChevronRightIcon } from '@radix-ui/react-icons'
import { Icon } from '@iconify/react'
import * as Popover from '@radix-ui/react-popover'
import { AnimatePresence, motion } from 'framer-motion'

interface MenuGroupProps {
  parentMenu: {
    id: number
    name: string
    icon: string
  }
  childMenus: Array<{
    id: number
    name: string
    icon: string
  }>
  isOpen: boolean
  onToggle: () => void
  collapsed: boolean
}

const MenuGroup: React.FC<MenuGroupProps> = ({ parentMenu, childMenus, isOpen, onToggle, collapsed }) => {
  if (collapsed) {
    return (
      <Popover.Root>
        <Popover.Trigger asChild>
          <button className="flex items-center justify-center w-16 h-16 hover:bg-gray-200 dark:hover:bg-gray-700">
            <Icon icon={'carbon:' + parentMenu.icon} className="text-xl" />
          </button>
        </Popover.Trigger>
        <Popover.Portal>
          <Popover.Content className="bg-white dark:bg-gray-800 shadow-lg rounded-md p-2 min-w-[200px]" sideOffset={5} side="right">
            {childMenus.map((item) => (
              <MenuItem key={item.id} item={item} collapse />
            ))}
            <Popover.Arrow className="fill-white dark:fill-[#111827]" />
          </Popover.Content>
        </Popover.Portal>
      </Popover.Root>
    )
  }
  return (
    <Collapsible.Root className="mb-4" open={isOpen} onOpenChange={onToggle}>
      <Collapsible.Trigger className="flex items-center justify-between w-full p-4 gap-3 rounded hover:bg-[#ECF5FF] dark:hover:bg-[#343435]">
        <motion.div layout className="flex items-center gap-2">
          <Icon icon={'carbon:' + parentMenu.icon} className="text-xl" />
          <AnimatePresence>
            {!collapsed && (
              <motion.span
                style={{ minWidth: '100px' }}
                initial={{ scale: 0.8, opacity: 0 }}
                animate={{ scale: 1, opacity: 1 }}
                exit={{ scale: 0.8, opacity: 0 }}
                transition={{ duration: 0.3 }}
                className="text-[14px]"
              >
                {!collapsed && parentMenu.name}
              </motion.span>
            )}
          </AnimatePresence>
        </motion.div>
        <ChevronRightIcon className={`transform transition-transform duration-200 ${isOpen ? 'rotate-90' : ''}`} />
      </Collapsible.Trigger>
      <Collapsible.Content className="CollapsibleContent">
        <motion.div className="mt-2">
          {childMenus.map((item) => (
            <MenuItem key={item.id} item={item} />
          ))}
        </motion.div>
      </Collapsible.Content>
    </Collapsible.Root>
  )
}

export default MenuGroup
