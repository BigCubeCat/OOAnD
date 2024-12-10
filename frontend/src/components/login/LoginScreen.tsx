import {LoginButton} from "@telegram-auth/react";
import {envBotName} from "../../utils/env.ts";
import {Box} from "@mui/joy";
import {apiLogin} from "../../api/auth.ts";

export default function LoginScreen() {
    console.log(envBotName());
    return (
        <Box sx={{display: 'flex', alignContent: 'center'}}>
            <LoginButton
                botUsername={envBotName()}
                onAuthCallback={async (data: any) => {
                    console.log(typeof data);
                    console.log(data);
                    await apiLogin(data);
                    // call your backend here to validate the data and sign in the user
                }}
            />
        </Box>
    );
}
