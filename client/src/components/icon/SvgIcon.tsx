import React, { useMemo } from 'react'

type SvgIconProps = {
  prefix?: string
  name: string
  color?: string
  size?: string
}

const SvgIcon: React.FC<SvgIconProps> = ({ prefix = 'icon', name, color = '#333', size = '1em' }) => {
  const symbolId = useMemo(() => `${prefix}-${name}`, [prefix, name])
  return (
    <>
      <svg aria-hidden="true" width={size} height={size}>
        <use href={symbolId} fill={color} />
      </svg>
    </>
  )
}
export default SvgIcon
