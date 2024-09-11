import { useDark } from '@/hooks/useDark'
import { useRef, useState } from 'react'
import ReactQuill from 'react-quill'
import 'react-quill/dist/quill.snow.css'
import { toast } from 'react-toastify'

const Editor = () => {
  const { darkMode } = useDark()
  const quillRef = useRef<ReactQuill | null>(null)
  const bgColors = ['purple', '#785412', '#452632', '#856325', '#963254', '#254563', !darkMode ? '#fff' : '#000']
  const fontColors = ['purple', '#785412', '#452632', '#856325', '#963254', '#254563', darkMode ? '#fff' : '#000']
  const modules = {
    toolbar: [
      [{ header: [1, 2, 3, 4, 5, 6, false] }],
      [{ font: [] }], // 添加字体选择
      [{ size: [] }], // 添加字体大小选择
      ['bold', 'italic', 'underline', 'strike', 'blockquote'],
      [{ align: ['right', 'center', 'justify'] }],
      [{ list: 'ordered' }, { list: 'bullet' }, { indent: '-1' }, { indent: '+1' }],
      ['link', 'image'],
      [{ color: fontColors }],
      [{ background: bgColors }],
      ['clean'],
    ],
  }

  const formats = ['header', 'font', 'size', 'bold', 'italic', 'underline', 'strike', 'blockquote', 'list', 'bullet', 'indent', 'link', 'color', 'image', 'background', 'align', 'clean']

  const [code, setCode] = useState(() => {
    // 从localStorage读取内容
    const savedContent = localStorage.getItem('editorContent')
    console.log('savedContent:', savedContent)
    return savedContent || 'hello guys you can also add fonts and other features to this editor.'
  })

  const handleProcedureContentChange = (content: string) => {
    setCode(content)
    localStorage.setItem('editorContent', content) // 保存到localStorage
  }

  const handleSave = () => {
    //清除编辑器内容
    const editor = quillRef.current!.getEditor()
    editor.setText('')
    console.log(quillRef.current!.value)
    setCode('')
    //清除localStorage
    localStorage.removeItem('editorContent')
    toast.success('保存成功')
  }

  return (
    <>
      <div>
        {/* 强制组件重新渲染 */}
        <ReactQuill ref={quillRef} key={darkMode ? 'dark' : 'light'} theme="snow" modules={modules} formats={formats} value={code} onChange={handleProcedureContentChange} />

        {/* 添加保存按钮 */}
        <button onClick={handleSave} className="mt-3 px-4 py-2 bg-blue-500 text-white dark:text-black rounded-md hover:bg-blue-600">
          点击保存内容
        </button>
      </div>
    </>
  )
}

export default Editor
