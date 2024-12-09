import {LoginButton} from "@telegram-auth/react";
import {envBotName} from "../../utils/env.ts";

export default function LoginScreen() {
    console.log(envBotName());
  return (
      <LoginButton
          botUsername={envBotName()}
          onAuthCallback={(data: any) => {
              console.log(typeof data);
              console.log(data);
              // call your backend here to validate the data and sign in the user
          }}
      />
  );
}
