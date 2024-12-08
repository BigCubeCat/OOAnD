import Main from "./statistic/Main.tsx";

import { CssVarsProvider } from '@mui/joy/styles';
import Header from "./bars/Header.tsx";
import Footer from "./bars/Footer.tsx";
import appTheme from "./themeController.ts";
import {Box} from "@mui/joy";

import {Route, Switch} from "wouter";
import LoginScreen from "./login/LoginScreen.tsx";

export default function Sceen() {
    // TODO: понять, авторизиованы ли мы
    return (
        <CssVarsProvider theme={appTheme}>
            <Box sx={{
                width: '100vw',
                height: '100vh',
                display:'flex',
                justifyContent: 'space-between',
                flexDirection: 'column'
            }}>
                <Header />
                <Switch>
                    <Route path="/" component={Main} />
                    <Route path="/login/" component={LoginScreen} />
                    {/* Default route in a switch */}
                    <Route>404: No such page!</Route>
                </Switch>
                <Footer />
            </Box>
        </CssVarsProvider>
    );
}