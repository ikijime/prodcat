/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{templ,html,js}"],
  theme: {
    extend: {
      colors: {
        'ink': '#181a26',
        'gray-violet0': '#12131c',
        'gray-violet': '#2C2F40',
        'gray-violet2': '#1E202F',
        'dead-skin': '#A69D9C',
        'near-black': '#08080D',
        'vein-blood': '#401021',
        'blonde': '#D1CEBF',
        'firebrick': '#b22222',
        'bloodred': '880808'
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
  ],
}

