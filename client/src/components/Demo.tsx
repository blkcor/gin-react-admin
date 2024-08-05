import { fetchUrlAtom } from '@/stores/countAtom'
import { useAtom } from 'jotai'

const Demo = () => {
  const [data] = useAtom(fetchUrlAtom)

  return <pre>{JSON.stringify(data, null, 2)}</pre>
}
export default Demo
