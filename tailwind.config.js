module.exports = {
    purge: {
        content: [
            './*.go',
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