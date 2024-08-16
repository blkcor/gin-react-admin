import { useEffect } from 'react'
import { Outlet, useNavigate } from 'react-router-dom'

export default function Root() {
  const navigate = useNavigate()

  useEffect(() => {
    const handleUnauthorized = (event: CustomEventInit) => {
      navigate(event.detail)
    }

    window.addEventListener('unauthorized', handleUnauthorized)

    return () => {
      window.removeEventListener('unauthorized', handleUnauthorized)
    }
  }, [navigate])
  return (
    <>
      <header>this is header!</header>
      <Outlet />
      <footer>this is footer!</footer>
    </>
  )
}
