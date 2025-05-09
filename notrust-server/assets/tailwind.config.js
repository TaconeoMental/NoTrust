/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './assets/js/**/*.js',
    './assets/css/**/*.css',
    './src/html/**/*.html',
    './**/*.html'
  ],
  theme: {
    extend: {
      colors: {
        primary: '#e11d48' // Rosado oscuro (tailwind red-600 aprox)
      },
      animation: {
        'fade-in-out': 'fadeInOut 3s ease-in-out forwards'
      },
      keyframes: {
        fadeInOut: {
          '0%': { opacity: '0' },
          '10%': { opacity: '1' },
          '90%': { opacity: '1' },
          '100%': { opacity: '0' }
        }
      }
    }
  },
  plugins: []
}

