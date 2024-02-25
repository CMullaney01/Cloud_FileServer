/** @type {import('tailwindcss').Config}*/
module.exports = {
  content: ["./src/**/**/*.{js,ts,jsx,tsx}"],
  daisyui: {
    themes: ["light", "dark", "dim"],
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("daisyui")
  ],
};