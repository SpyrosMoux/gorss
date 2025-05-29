/* eslint-disable */
import { PaletteOptions, Palette } from "@mui/material/styles";
import { AppBarPropsColorOverrides } from "@mui/material/AppBar";
import { IconButtonPropsColorOverrides } from "@mui/material/IconButton";

declare module "@mui/material/styles" {
  interface Palette {
    white: Palette["primary"];
  }
  interface PaletteOptions {
    white?: PaletteOptions["primary"];
  }
}

declare module "@mui/material/AppBar" {
  interface AppBarPropsColorOverrides {
    white: true;
  }
}

declare module "@mui/material/IconButton" {
  interface IconButtonPropsColorOverrides {
    white: true;
  }
}
