import { Outlet } from 'react-router-dom'

export default function Root() {
  return (
    <>
      <header>this is header!</header>
      <Outlet />
      <footer>this is footer!</footer>
    </>
  )
}
