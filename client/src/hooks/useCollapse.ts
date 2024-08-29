import { systemAtom } from "@/stores/systemAtom"
import { useAtom } from "jotai"


export const useCollapse = () => {
  const [systemConfig, setSystemConfig] = useAtom(systemAtom)

  const toggleCollapse = () => {
    setSystemConfig((prev) => ({ ...prev, collapsed: !prev.collapsed }))
    //更新localStorage
    localStorage.setItem('systemConfig', JSON.stringify({ ...systemConfig, collapsed: !systemConfig.collapsed }))
  }
  return { collapsed: systemConfig.collapsed, toggleCollapse }
}
