import { extendTheme } from '@mui/joy/styles';

const palette = {
    primary: {
        700: '#FFCB3B',
        400: '#FFCB3B',
    },
    secondary: {
        700: '#D9D9D9',
    },
    neutral: {
        solidBg: '#6c757d',
        solidBorder: '#6c757d',
        solidHoverBg: '#5c636a',
        solidHoverBorder: '#565e64',
        solidActiveBg: '#565e64',
        solidActiveBorder: '#51585e',
        solidDisabledBg: '#6c757d',
        solidDisabledBorder: '#6c757d',
        // btn-light
        softColor: '#000',
        softBg: '#f8f9fa',
        softBorder: '#f8f9fa',
        softHoverBg: '#f9fafb',
        softHoverBorder: '#f9fafb',
        softActiveBg: '#f9fafb',
        softActiveBorder: '#f9fafb',
        softDisabledBg: '#f8f9fa',
        softDisabledBorder: '#f8f9fa',
    },
    success: {
        100: '#337A45',
    },
    danger: {
        100: '#FF5454'
    },
    warning: {
        solidColor: '#000',
        solidBg: '#ffc107',
        solidBorder: '#ffc107',
        solidHoverBg: '#ffca2c',
        solidHoverBorder: '#ffc720',
        solidActiveBg: '#ffcd39',
        solidActiveBorder: '#ffc720',
        solidDisabledBg: '#ffc107',
        solidDisabledBorder: '#ffc107',
    },
    info: {
        0: '#000',
        solidBg: '#0dcaf0',
        solidBorder: '#0dcaf0',
        solidHoverBg: '#31d2f2',
        solidHoverBorder: '#25cff2',
        solidActiveBg: '#3dd5f3',
        solidActiveBorder: '#25cff2',
        solidDisabledBg: '#0dcaf0',
        solidDisabledBorder: '#0dcaf0',
    },
};

const appTheme = extendTheme({
    cssVarPrefix: 'bs',
    colorSchemes: {
        light: { palette },
        dark: { palette },
    },
    components: {
        JoyButton: {
            styleOverrides: {
                root: ({ ownerState, theme }) => ({
                    letterSpacing: 'normal',
                    fontWeight: theme.vars.fontWeight.md,
                    fontFamily: theme.vars.fontFamily.fallback,
                    outlineWidth: 0,
                    borderRadius: '0.375rem',
                    transition:
                        'color .15s ease-in-out,background-color .15s ease-in-out,border-color .15s ease-in-out,box-shadow .15s ease-in-out',
                    ...(ownerState.size === 'md' && {
                        paddingInline: '0.75rem',
                        minHeight: 38,
                    }),
                }),
            },
        },
    },
});

export default appTheme;