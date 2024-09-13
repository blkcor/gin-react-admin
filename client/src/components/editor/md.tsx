import { Flex, TextArea } from '@radix-ui/themes'
import { ChangeEvent, FC, useEffect, useState } from 'react'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import remarkBreaks from 'remark-breaks'
import remarkToc from 'remark-toc'
import rehypeHighlight from 'rehype-highlight'
import rehypeRaw from 'rehype-raw'
import 'github-markdown-css'
import './index.css'
import { useClipboard } from '@/hooks/useClipBorad'
import ClipBorad from '@/assets/svg/ClipBoard.svg?raw'
import ClipedBorad from '@/assets/svg/ClipedBoard.svg?raw'
import { toast } from 'react-toastify'
type modeType = 'editor' | 'preview'

type MarkdownEditorProps = {
  mode: modeType
}

const MarkdownEditor: FC<MarkdownEditorProps> = ({ mode }) => {
  const [content, setContent] = useState('')
  const { copy } = useClipboard()
  const handleInputMarkdown = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setContent(e.target.value)
  }

  useEffect(() => {
    const addCopyButtons = () => {
      document.querySelectorAll('.preview-container pre code').forEach((block) => {
        // Create button
        const button = document.createElement('button')
        button.innerHTML = ClipBorad
        // Style button
        Object.assign(button.style, {
          position: 'absolute',
          top: '5px',
          right: '5px',
          cursor: 'pointer',
          padding: '2px 5px',
          fontSize: '12px',
        })

        // Add copy functionality
        button.addEventListener('click', () => {
          copy(block.textContent)
          toast.success('Copied to clipboard')
          button.innerHTML = ClipedBorad

          setTimeout(() => {
            button.innerHTML = ClipBorad
          }, 2000)
        })

        // Add button to pre element (parent of code)
        const pre = block.parentNode
        if (pre instanceof HTMLElement) {
          pre.style.position = 'relative'
          pre.insertBefore(button, pre.firstChild)
        }
      })
    }

    // Run the function after a short delay to ensure the Markdown has been rendered
    if (mode === 'preview') {
      setTimeout(addCopyButtons, 100)
    }
  }, [content, mode])

  return (
    <div className="markdown-body">
      <Flex direction={mode === 'editor' ? 'column' : 'row'} gap="3">
        {mode === 'editor' && (
          <Flex direction="column" flexGrow="1">
            <TextArea variant="surface" value={content} onChange={handleInputMarkdown} style={{ height: '70vh' }} />
          </Flex>
        )}
        {mode === 'preview' && (
          <Flex direction="column" flexGrow="1">
            <div className="preview-container">
              <Markdown remarkPlugins={[remarkGfm, remarkBreaks, [remarkToc, { heading: 'Contents' }]]} rehypePlugins={[rehypeHighlight, rehypeRaw]}>
                {content}
              </Markdown>
            </div>
          </Flex>
        )}
      </Flex>
    </div>
  )
}

export default MarkdownEditor
