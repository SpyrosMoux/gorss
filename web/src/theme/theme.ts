import { createTheme } from "@mui/material";
import { palette } from "./palette";

export const theme = createTheme({
  typography: {
    fontFamily: "Space Mono",
  },

  palette: palette,
  components: {
    MuiAppBar: {
      defaultProps: {
        color: "white",
        sx: {
          boxShadow: "none",
          borderBottom: "1px dashed #b2b2b2",
        },
      },
    },
    MuiButtonBase: {
      defaultProps: {
        disableRipple: true,
      },
    },
    MuiButton: {
      styleOverrides: {
        text: {
          border: "1px dashed transparent",
          borderRadius: 0,

          ":hover": {
            border: "1px dashed #000",
            borderRadius: 0,
            backgroundColor: "transparent",
          },
        },
      },
    },
    MuiIconButton: {
      styleOverrides: {
        root: {
          borderRadius: 0,
          border: "1px dashed transparent",
          "&:hover": {
            backgroundColor: "transparent",
            border: "1px dashed #000",
          },
        },
      },
      defaultProps: {
        sx: {
          borderRadius: 0,
        },
      },
    },
  },
});
