/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "#0f172a",
        primary: "#3b82f6",
        secondary: "#64748b",
        accent: "#f43f5e",
      }
    },
  },
  plugins: [],
}
