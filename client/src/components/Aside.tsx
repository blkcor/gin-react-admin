import Menu from './Menu'
import logo from '@/assets/images/logo_with_no_word.svg'

const Aside = () => {
  return (
    <div className="h-screen flex flex-col border-r-2 border-middle w-[15vw]">
      <div className="flex items-center p-2 border-b-2 border-middle bg-white">
        <img className="h-8" src={logo} alt="logo" />
        <span className="ml-2 text-xs font-bold flex-shrink-0">Gin React Admin</span>
      </div>
      <Menu />
    </div>
  )
}

export default Aside
