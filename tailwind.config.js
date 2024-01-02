/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.{html,js,templ,go}"],
    theme: {
        extend: {
            animation: {
                "ping-slow": "ping 3s cubic-bezier(0, 0, 0.2, 1) infinite"
            }
        }
    },
    plugins: []
}
