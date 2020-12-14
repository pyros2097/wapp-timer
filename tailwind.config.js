module.exports = {
    purge: [
        './*.go',
        './vendor/github.com/pyros2097/wapp/*.go',
    ],
    theme: {
        extend: {
            fontSize: {
                '10xl': ['20rem', { lineHeight: '1' }],
            },
        },
    },
    variants: {},
    plugins: [],
}