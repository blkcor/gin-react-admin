import { FC, useState } from 'react'
import { InputIcon, TrackPreviousIcon } from '@radix-ui/react-icons'
import './index.css'

export type modeType = 'editor' | 'preview'

type MarkdownModeProps = {
  handleSetMode: (mode: modeType) => void
}
const MarkdownMode: FC<MarkdownModeProps> = ({ handleSetMode }) => {
  const [mode, setMode] = useState<modeType>('editor')
  const handleChangeMode = (mode: modeType) => {
    setMode(mode)
    handleSetMode(mode)
  }
  return (
    <div className="flex flex-col">
      <div className="inline-flex">
        <button className={`p-1 box-border rounded-md  ${mode === 'editor' ? 'bg-gray-400/30' : ''}`} onClick={() => handleChangeMode('editor')}>
          <InputIcon />
        </button>
        <button className={`p-1 box-border rounded-md  ${mode === 'preview' ? 'bg-gray-400/30' : ''}`} onClick={() => handleChangeMode('preview')}>
          <TrackPreviousIcon />
        </button>
      </div>
    </div>
  )
}
export default MarkdownMode
