import { CssVarsProvider } from '@mui/joy/styles';
import Header from "./bars/Header";
import Footer from "./bars/Footer";
import appTheme from "./themeController";
import {Box, Typography} from "@mui/joy";

import {Route, Switch} from "wouter";
import LoginScreen from "./login/LoginScreen";
import Debts from "./statistic/debts/Debts";
import {apiAddress} from "../utils/env";

export default function Screen() {
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
                    <Route path="/" component={Debts} />
                    <Route path="/login/" component={LoginScreen} />
                </Switch>
                <Typography variant="body2" color="textSecondary">
                   api= {apiAddress()}
                </Typography>
                <Footer />
            </Box>
        </CssVarsProvider>
    );
}