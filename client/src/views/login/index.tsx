import logoTransparent from '@/assets/images/logo_transparent.png'
import { Helmet, HelmetProvider } from 'react-helmet-async'
import { PersonIcon, EyeOpenIcon, EyeClosedIcon, GearIcon, LockClosedIcon } from '@radix-ui/react-icons'
import { useEffect, useRef, useState } from 'react'
import { useRequest } from 'alova/client'
import apis from '@/apis/apis'
import { Bounce, toast } from 'react-toastify'
import { useForm, Controller } from 'react-hook-form'
import { Spinner } from '@radix-ui/themes'
import { useAtom } from 'jotai'
import { userAtom } from '@/stores/userAtom'
import { useNavigate } from 'react-router-dom'
const Login = () => {
  interface LoginFormProps {
    username: string
    password: string
    captcha: string
  }

  const { loading, send, onError, onSuccess } = useRequest(apis.getCaptcha, {
    initialData: {}, // 设置data状态的初始数据
    immediate: true, // 是否立即发送请求，默认为true
  })

  const navigate = useNavigate()
  const [captchaSrc, setCaptchaSrc] = useState<string | undefined>(undefined)
  const [, setUser] = useAtom(userAtom)

  onError((event) => {
    toast.error(event.error, {
      position: 'top-center',
    })
  })

  onSuccess(({ data }) => {
    if (data instanceof Blob) {
      const url = URL.createObjectURL(data)
      setCaptchaSrc(url)
      return () => URL.revokeObjectURL(url) // 清理URL
    }
  })

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormProps>({
    mode: 'onBlur',
    defaultValues: {
      username: '',
      password: '',
      captcha: '',
    },
  })

  const [openEye, setOpenEye] = useState(false)
  const passwordInputRef = useRef<HTMLInputElement | null>(null)
  const [isLogin, setIsLogin] = useState<boolean>(false)

  useEffect(() => {
    if (passwordInputRef.current) {
      passwordInputRef.current.type = openEye ? 'text' : 'password'
    }
  }, [openEye])

  const onSubmit = async (data: LoginFormProps) => {
    try {
      setIsLogin(true)
      const result = await apis.login(data)

      setUser({
        userInfo: result.data,
        token: result.token,
      })
      localStorage.setItem('token', result.token)
      localStorage.setItem('userInfo', JSON.stringify(result.data))
      toast.success('登录成功!', {
        position: 'top-center',
        autoClose: 2000,
        closeOnClick: true,
        pauseOnHover: true,
        transition: Bounce,
      })
      //跳转到主页
      navigate('/')
    } catch (e) {
      setIsLogin(false)
      if (e instanceof Error) {
        toast.error('登录失败：' + e.message, {
          position: 'top-center',
        })
      }
    }
  }

  const handleRefreshCaptcha = () => {
    send()
  }

  return (
    <>
      <HelmetProvider>
        <Helmet>
          <title>登录页 | Gin React Admin</title>
          <meta name="description" content="GRA登录页" />
        </Helmet>
      </HelmetProvider>
      <div className="bg-login-bg w-full h-full bg-cover flex items-center justify-center py-6">
        <div className="bg-[#18181C] px-0 py-10 rounded-lg shadow-xl flex flex-col justify-center md:flex-row items-center max-w-4xl w-full">
          <img src={logoTransparent} alt="Logo" className="w-72 h-72 object-cover mb-6 md:mb-0 md:mr-2" />

          <div className="text-[#6a6a6a] p-6 rounded-md w-full max-w-md flex flex-col gap-2">
            <h2 className="text-3xl mb-3 text-center">Gin React Admin</h2>
            <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2">
              <div className="mb-4">
                <div className="flex flex-col mb-2">{errors.username && <p className="text-red-500 text-xs mt-1">用户名不能为空</p>}</div>
                <div>
                  <div className="flex items-center relative">
                    <PersonIcon className="w-5 h-5 absolute left-3" />
                    <Controller
                      name="username"
                      control={control}
                      rules={{ required: true }}
                      render={({ field }) => (
                        <input
                          autoComplete="off"
                          {...field}
                          type="text"
                          placeholder="请输入用户名"
                          className="w-full pl-10 text-white py-3 bg-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        />
                      )}
                    />
                  </div>
                </div>
              </div>
              <div className="mb-6">
                <div className="flex flex-col mb-2">{errors.password && <p className="text-red-500 text-xs mt-1">密码不能为空</p>}</div>
                <div>
                  <div className="flex items-center relative">
                    <LockClosedIcon className="w-5 h-5 absolute left-3" />
                    <Controller
                      name="password"
                      control={control}
                      rules={{ required: true }}
                      render={({ field }) => (
                        <input
                          autoComplete="off"
                          {...field}
                          ref={passwordInputRef}
                          type="text"
                          placeholder="请输入密码"
                          className="w-full pl-10 py-3 bg-gray-700 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        />
                      )}
                    />
                    {openEye ? (
                      <EyeOpenIcon className="w-5 h-5 absolute right-4 cursor-pointer" onClick={() => setOpenEye(false)} />
                    ) : (
                      <EyeClosedIcon className="w-5 h-5 absolute right-4 cursor-pointer" onClick={() => setOpenEye(true)} />
                    )}
                  </div>
                </div>
              </div>
              <div className="mb-6">
                <div className="flex flex-col mb-2">{errors.captcha && <p className="text-red-500 text-xs mt-1">验证码不能为空</p>}</div>
                <div>
                  <div className="flex items-center relative">
                    <GearIcon className="w-5 h-5 absolute left-3" />
                    <div className="flex gap-5 justify-between w-full items-center">
                      <Controller
                        name="captcha"
                        control={control}
                        rules={{ required: true }}
                        render={({ field }) => (
                          <input
                            autoComplete="off"
                            {...field}
                            type="text"
                            placeholder="请输入验证码"
                            className="flex-1 pl-10 py-3 bg-gray-700 text-white rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                          />
                        )}
                      />
                      {loading ? (
                        <div className="w-full h-11  flex items-center justify-center rounded-md border-2 border-solid border-[#598b81] gap-3">
                          <span>正在加载</span>
                          <Spinner size={'2'} loading={true} />
                        </div>
                      ) : (
                        <img onClick={handleRefreshCaptcha} id="captcha" src={captchaSrc} className="w-full h-11 object-cover cursor-pointer rounded-md" alt="验证码" />
                      )}
                    </div>
                  </div>
                </div>
              </div>
              <div className="flex gap-6">
                <button
                  disabled={isLogin}
                  type="button"
                  className={`w-full bg-[#18181C] border-[#598B8E] border-2 border-solid text-[#598B8E] py-3 rounded-md  transition duration-300 ${
                    isLogin ? ' cursor-not-allowed' : 'hover:text-white'
                  }`}
                >
                  一键体验
                </button>
                {isLogin ? (
                  <button disabled className="w-full bg-[#598B8E] text-black py-3 rounded-md cursor-not-allowed transition duration-300">
                    <div className="flex gap-2 justify-center items-center">
                      正在登录
                      <Spinner size={'2'} loading={true} />
                    </div>
                  </button>
                ) : (
                  <button type="submit" className="w-full bg-[#598B8E] text-black py-3 rounded-md hover:bg-[#598B8E]/50 transition duration-300">
                    登录
                  </button>
                )}
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  )
}

export default Login
