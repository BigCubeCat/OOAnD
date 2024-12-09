import * as React from "react";
import { Box, Button, IconButton, Input } from "@mui/joy";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import SendIcon from "@mui/icons-material/Send";
import useEmailFormHook from "../../hooks/login/email_form_hook";

type TEmailButtonForm = "button" | "email" | "token";

export default function EmailForm() {
  const [state, setState] = React.useState<TEmailButtonForm>("button");
  const { validEmail, email, setEmail, token, setToken } = useEmailFormHook();
  let view: any;
  if (state == "button") {
    view = (
      <Button sx={{ width: 400 }} onClick={() => setState("email")}>
        Auth with email
      </Button>
    );
  } else if (state == "email") {
    view = (
      <Input
        type="email"
        sx={{ width: 400 }}
        placeholder="Enter email address"
        color={validEmail ? "neutral" : "danger"}
        value={email}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
          setEmail(e.target.value)
        }
        startDecorator={
          <IconButton
            variant="solid"
            color="neutral"
            onClick={() => setState("button")}
          >
            <ArrowBackIcon />
          </IconButton>
        }
        endDecorator={
          <IconButton
            variant="solid"
            color="primary"
            onClick={() => (validEmail ? setState("token") : setState("email"))}
          >
            <SendIcon />
          </IconButton>
        }
      />
    );
  } else {
    view = (
      <Input
        value={token}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
          setToken(e.target.value)
        }
        sx={{ width: 400 }}
        placeholder="Enter token from your email"
        startDecorator={
          <IconButton onClick={() => setState("email")}>
            <ArrowBackIcon />
          </IconButton>
        }
        endDecorator={
          <IconButton onClick={() => setState("button")}>
            <SendIcon />
          </IconButton>
        }
      />
    );
  }
  return <Box>{view}</Box>;
}
