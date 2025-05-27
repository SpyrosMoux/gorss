import { CssBaseline, ThemeProvider } from "@mui/material";
import { theme } from "./theme/theme";
import { Router } from "./router/Router";

function App() {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router />
    </ThemeProvider>
  );
}

export default App;
