import { userAtom } from '@/stores/userAtom'
import { useAtom } from 'jotai'
import { useEffect } from 'react'

const Home = () => {
  const [user, setUser] = useAtom(userAtom)
  useEffect(() => {
    //页面刷新的时候从localStorage进行初始化userAtom
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
    <div>
      <pre>{JSON.stringify(user)}</pre>
    </div>
  )
}
export default Home
