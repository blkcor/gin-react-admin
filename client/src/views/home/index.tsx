import { useAtom } from 'jotai'
import { doubleCountAtom, countAtom, sumAtom1, sumAtom2 } from '@/stores/countAtom'
// import Demo from '@/components/Demo'
// import { Suspense } from 'react'
import Button from '@mui/material/Button'
const Home = () => {
  const [count, setCount] = useAtom(countAtom)
  const [doubleCount] = useAtom(doubleCountAtom)
  const [sum1] = useAtom(sumAtom1)
  const [sum2] = useAtom(sumAtom2)

  return (
    <div>
      <h1>count is : {count}</h1>
      <h1>doubleCount is : {doubleCount}</h1>
      <h1>sumCount is : {sum1}</h1>
      <h1>sumCount is : {sum2}</h1>
      {/* <Suspense fallback={<div>loading...</div>}>
        <Demo />
      </Suspense> */}
      <p>HOME PAGE</p>
      <div className="flex gap-10 justify-center">
        <Button variant="contained" onClick={() => setCount((c) => c + 1)}>
          点我+1
        </Button>
        <Button variant="contained" onClick={() => setCount((c) => c - 1)}>
          点我-1
        </Button>
      </div>
    </div>
  )
}
export default Home
