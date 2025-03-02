import { createSignal } from "solid-js";
import "./index.scss";
import api from "../../api/api";
import { useNavigate } from "@solidjs/router";

export default function Login() {
  const [username, setUserName] = createSignal("");
  const [pwd, setPWD] = createSignal("");
  const navigate = useNavigate();
  function login() {
    if (username() === "" || pwd() === "") return;
    api.authLogin(username(), pwd()).then((res) => {
      if (res.code === 0 && res.data.jwt_token) {
        localStorage.setItem("jwt", res.data.jwt_token);
        navigate("/home");
      }
    });
  }

  return (
    <>
      <div class="login-box">
        <div>Login</div>
        <input
          value={username()}
          onChange={(e) => setUserName(e.target.value)}
          type="text"
        />
        <input
          value={pwd()}
          onChange={(e) => setPWD(e.target.value)}
          type="password"
        />
        <button onClick={login}>login</button>
      </div>
    </>
  );
}
