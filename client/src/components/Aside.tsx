import logo from '@/assets/images/logo_transparent.svg'

const Aside = () => {
  return (
    <aside className="w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700">
      <div className="flex items-center justify-center h-16 border-b border-gray-200 dark:border-gray-700">
        {/* logo区域 */}
        <div className="flex items-center p-2 gap-2">
          <img src={logo} alt="logo" className="w-10 h-10" />
          <span className="text-2xs font-bold text-gray-800 dark:text-white">Gin React Admin</span>
        </div>
      </div>
    </aside>
  )
}

export default Aside
