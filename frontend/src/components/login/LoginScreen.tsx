import { Box } from "@mui/joy";
import EmailForm from "./email";

export default function LoginScreen() {
  return (
    <Box
      sx={{
        display: "flex",
        justifyContent: "center",
      }}
    >
      <Box
        sx={{
          width: "100vh",
          display: "flex",
          justifyContent: "center",
          flexDirection: "column",
        }}
      >
        <EmailForm />
      </Box>
    </Box>
  );
}
