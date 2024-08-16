import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      backgroundImage: {
        'login-bg': "url('/src/assets/images/login-bg.jpg')",
      },
      colors: {
        disabledBackground: '#e0e0e0',
        disabledText: '#a0a0a0',
        disabledBorder: '#d0d0d0',
      },
      cursor: {
        'not-allowed': 'not-allowed',
      },
    },
  },
  plugins: [],
}

export default config
