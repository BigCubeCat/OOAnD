import { useState } from "react";

const validateEmail = (email: string) => {
  return String(email)
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
    );
};

export default function useEmailFormHook() {
  const [validEmail, setValidEmail] = useState(false);
  const [email, setEmail] = useState<string>("");
  const [token, setToken] = useState<string>("");

  const SetEmail = (mail: string) => {
    if (validateEmail(mail)) {
      setValidEmail(true);
      setEmail(mail);
    }
  };

  return {
    validEmail,
    email,
    setEmail: SetEmail,
    token,
    setToken,
  };
}
