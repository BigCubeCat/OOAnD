import { useState } from "react";
import { LoginButton } from "@telegram-auth/react";
import LoginScreen from "./components/login/LoginScreen";

function App() {
  const [count, setCount] = useState(0);
  return (
    <div className="App">
      {count}
      <LoginScreen />
      <LoginButton
        botUsername={"ooad_project_bot"}
        onAuthCallback={(data) => {
          console.log(data);
          setCount(count + 1);
          // call your backend here to validate the data and sign in the user
        }}
      />
    </div>
  );
}

export default App;
