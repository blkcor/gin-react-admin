import { userAtom } from '@/stores/userAtom'
import { Avatar } from '@radix-ui/themes'
import { useAtom } from 'jotai'
import { useEffect } from 'react'

const Home = () => {
  const [user, setUser] = useAtom(userAtom)

  useEffect(() => {
    // 页面刷新的时候从localStorage进行初始化userAtom
    if (!user.token || !user.userInfo.userId) {
      const token = localStorage.getItem('token')!
      const userInfo = JSON.parse(localStorage.getItem('userInfo')!)
      setUser({
        token,
        userInfo,
      })
    }
  }, [])

  return (
    <div className="w-full flex flex-col ">
      <div className="p-2 grid grid-cols-3 gap-4 ">
        <div className="col-span-1 flex flex-col bg-white p-4 shadow-md">
          <div className="flex gap-3">
            <Avatar className="w-16 h-16 rounded-full" src={user.userInfo.avatar} fallback="A" />
            <div className="flex flex-col justify-center">
              <div className="text-lg opacity-80">Hello, {user.userInfo.roleCode}</div>
              <div className="text-sm text-gray-500 opacity-50">当前角色: {user.userInfo.userRole}</div>
            </div>
          </div>
          <p className="mt-2 text-sm opacity-60">Be yourself; everyone else is already taken.</p>
          <p className="mt-2 text-right text-xs opacity-40">--Oscar Wilde</p>
        </div>

        <div className="col-span-2 flex flex-col bg-white p-4 gap-4 shadow-md">
          <p className="text-lg font-semibold">✨ 欢迎使用 Gin React Admin 1.0</p>
          <p className="text-sm opacity-60">GRA(Gin-React-Admin)是一个基于Gin和React的管理模板项目。该项目旨在提供一个通用的后端管理系统基础设施， 帮助开发人员快速构建具有响应式界面的管理面板。</p>
          <div className="flex gap-4 mt-4 justify-end">
            <button className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">开发文档</button>
            <button className="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600">代码仓库</button>
          </div>
        </div>
      </div>

      <div className="p-2 grid grid-cols-4 gap-4 ">
        <div className="col-span-2 flex flex-col bg-white py-2 shadow-md">
          <div className="flex font-bold text-lg p-4 border-b-[1px] border-middle">💯 特性</div>
          <div className="flex flex-col py-4 px-6 opacity-90 gap-4 text-sm">
            <p>🚀 快速：在五分钟内构建一个生产管理面板应用程序。</p>
            <p>🛡️ 安全性：基于 Casbin 和 JWT 身份验证的开箱即用 RBAC 身份验证系统。</p>
            <p>🎨 主题：优雅、美丽的设计。</p>
            <p>📦 组件：组件丰富实用。</p>
            <p>📱 响应式：支持移动设备。</p>
            <p>🌐 多语言：支持多种语言。</p>
            <p>🎯 标准：遵循RESTful API设计规范</p>
            <p>📚 中间件：基于GIN WEB框架提供丰富的中间件支持</p>
            <p>📝 Doc：支持 swagger 文档。</p>
          </div>
        </div>

        <div className="col-span-2 flex flex-col bg-white py-2 shadow-md">
          <div className="flex font-bold text-lg p-4 border-b-[1px] border-middle">🛠️ 技术栈</div>
          <div className="flex flex-col py-4 px-6 opacity-90 gap-4 text-sm">
            <p>🚀 快速：在五分钟内构建一个生产管理面板应用程序。</p>
            <p>🛡️ 安全性：基于 Casbin 和 JWT 身份验证的开箱即用 RBAC 身份验证系统。</p>
            <p>🎨 主题：优雅、美丽的设计。</p>
            <p>📦 组件：组件丰富实用。</p>
            <p>📱 响应式：支持移动设备。</p>
            <p>🌐 多语言：支持多种语言。</p>
            <p>🎯 标准：遵循RESTful API设计规范</p>
            <p>📚 中间件：基于GIN WEB框架提供丰富的中间件支持</p>
            <p>📝 Doc：支持 swagger 文档。</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home
