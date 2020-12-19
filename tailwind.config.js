module.exports = {
    purge: {
        content: [
            './*.go',
            './vendor/github.com/pyros2097/wapp/*.go',
        ],
        options: {
            safelist: ['flex', 'flex-row', 'justify-center', 'items-center'],
        }
    },
    theme: {
        extend: {
            fontSize: {
                '10xl': ['20rem', { lineHeight: '1' }],
            },
        },
    },
    variants: {
        extend: {
            opacity: ['disabled'],
        }
    },
    plugins: [],
}