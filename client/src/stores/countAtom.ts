import { atom } from 'jotai'

export const countAtom = atom(0)

//create a read-only atom
export const doubleCountAtom = atom((get) => get(countAtom) * 2)

//create an atom from multiple atoms
const count1 = atom(1)
const count2 = atom(2)
const count3 = atom(3)
export const sumAtom1 = atom((get) => get(count1) + get(count2) + get(count3))
//or you can use fp pattern
const atoms = [count1, count2, count3]
export const sumAtom2 = atom((get) => atoms.map((atom) => get(atom)).reduce((a, b) => a + b, 0))

//create an async atom
// 创建一个异步 atom
const urlAtom = atom('http://suggest.taobao.com/sug?code=utf-8&q=商品关键字&callback=cb')
export const fetchUrlAtom = atom(async (get) => {
  const url = get(urlAtom)
  const response = await fetch(url)
  if (!response.ok) {
    throw new Error('Network response was not ok')
  }
  return await response.json()
})
