module.exports = {
    purge: {
        content: [
            './*.go',
        ],
        options: {
            safelist: ['flex', 'flex-col', 'flex-row', 'justify-center', 'items-center'],
        }
    },
    theme: {
        extend: {
            fontSize: {
                '8xl': ['6rem', { lineHeight: '1' }],
                '10xl': ['20rem', { lineHeight: '1' }],
            },
            fontFamily: {
                "body": ["Helvetica"],
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