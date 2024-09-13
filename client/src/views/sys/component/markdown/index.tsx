import MarkdownEditor from '@/components/editor/md'
import MarkdownMode, { type modeType } from '@/components/editor/mdMode'
import { useState } from 'react'

const MarkdownEditorView = () => {
  const [mode, setMode] = useState<modeType>('editor')
  const handleSetMode = (mode: modeType) => {
    setMode(mode)
  }
  return (
    <>
      <MarkdownMode handleSetMode={handleSetMode} />
      <MarkdownEditor mode={mode} />
    </>
  )
}
export default MarkdownEditorView
