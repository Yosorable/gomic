import { A } from "@solidjs/router";

export default function NavBar() {
  return (
    <div
      style={{
        padding: "10px",
      }}
    >
      <nav>
        <A class="nav-a" href="/">
          Home
        </A>
        <A class="nav-a" href="/authors">
          Authors
        </A>
        <A class="nav-a" href="/syncing">
          Syncing
        </A>
      </nav>
    </div>
  );
}
