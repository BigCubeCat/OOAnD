import { CssVarsProvider } from '@mui/joy/styles';
import Header from "./bars/Header";
import Footer from "./bars/Footer";
import appTheme from "./themeController";
import {Box} from "@mui/joy";

import {Route, Switch} from "wouter";
import LoginScreen from "./login/LoginScreen";
import Debts from "./statistic/debts/Debts";
import useAuthToken from "../hooks/useAuthToken.ts";


export default function Screen() {
    const {token, authorized} = useAuthToken();
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
                    <Route path="/" component={Debts} />
                    <Route path="/login/" component={LoginScreen} />
                </Switch>
                <Footer />
            </Box>
        </CssVarsProvider>
    );
}