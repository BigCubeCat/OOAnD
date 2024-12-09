import Screen from "./components/Screen";
import appTheme from "./components/themeController.ts";
import {CssVarsProvider} from "@mui/joy/styles";


function App() {
    return (
        <CssVarsProvider theme={appTheme}>
            <Screen/>
        </CssVarsProvider>
    );
}

export default App;
