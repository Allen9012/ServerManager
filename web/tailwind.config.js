/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    important: true,
    theme: {
        extend: {
            backgroundColor: {
                "main": "#F5F5F5",
            }
        },
    },
    plugins: [],
    corePlugins: {
        preflight: false
    },
    configureWebPack: {
        devtool: 'source-map' // 输出 source-map 方便直接调试 ES6 源码
    }
}