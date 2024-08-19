import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      backgroundImage: {
        'login-bg': "url('/src/assets/images/login-bg.jpg')",
      },
      cursor: {
        'not-allowed': 'not-allowed',
      },
    },
  },
  plugins: [],
}

export default config
