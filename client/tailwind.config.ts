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
      borderColor: {
        middle: 'rbga(239,239,245,1)',
      },
    },
  },
  plugins: [],
}

export default config
