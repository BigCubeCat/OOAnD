import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

import { LoginButton } from "@telegram-auth/react";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <LoginButton
        botUsername="ooad_project_bot"
        authCallbackUrl="/path/to/callback/url"
        buttonSize="large" // "large" | "medium" | "small"
        cornerRadius={5} // 0 - 20
        showAvatar={true} // true | false
        lang="en"
      />
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
