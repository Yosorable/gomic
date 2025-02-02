import { createSignal, For, onMount } from "solid-js";
import api from "../../api/api";
import NavBar from "../components/NavBar";
import { Title } from "@solidjs/meta";

export default function Authors() {
  const [author, setAuthor] = createSignal<string[]>([]);
  onMount(() => {
    api.getAuthors().then((res) => {
      if (res.code === 0) {
        setAuthor(res.data);
      }
    });
  });
  return (
    <>
      <Title>作者</Title>
      <div class="main-box">
        <NavBar />
        <div class="tag-box">
          <For each={author()}>
            {(item) => (
              <a class="tag" href={"/author/" + item}>
                {item}
              </a>
            )}
          </For>
        </div>
      </div>
    </>
  );
}
