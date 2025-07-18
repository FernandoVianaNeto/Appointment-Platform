export const defaultTheme = {
    colors: {
      primary: '#1500ff',
      background: '#f2f2f2',
      highlighBackground: '#f7f7f7',
      highlightButton: '#b40005',
      text: '#333',
      button: '#000',
      white: '#fff',
      grey: '#808080'
    },
    spacing: {
      sm: '8px',
      md: '16px',
      lg: '24px',
    },
    borderRadius: '8px',
    font: {
      family: 'Mulish, sans-serif',
      size: {
        sm: '12px',
        md: '16px',
        lg: '20px',
      },
      weight: {
        regular: 400,
        medium: 500,
        bold: 700,
      },
    }
  };
  
export type ThemeType = typeof defaultTheme;