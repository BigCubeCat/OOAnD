import {Box} from "@mui/joy";
import {Route, Switch} from "wouter";

import Header from "./bars/Header";
import Footer from "./bars/Footer";
import LoginScreen from "./login/LoginScreen";
import Debts from "./statistic/debts/Debts";


export default function Screen() {
    return (
        <Box sx={{
            width: '100vw',
            height: '100vh',
            display: 'flex',
            justifyContent: 'space-between',
            flexDirection: 'column'
        }}>
            <Header/>
            <Switch>
                <Route path="/" component={Debts}/>
                <Route path="/login/" component={LoginScreen}/>
            </Switch>
            <Footer/>
        </Box>
    );
}