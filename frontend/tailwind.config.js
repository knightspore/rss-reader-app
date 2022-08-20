module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Barlow", "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
