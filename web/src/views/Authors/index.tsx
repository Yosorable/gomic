import { createSignal, For, onMount } from "solid-js";
import api, { Author } from "../../api/api";
import { useTitle } from "../../signal/title";

export default function Authors() {
  const [author, setAuthor] = createSignal<Author[]>([]);
  const [_, setTitle] = useTitle();
  onMount(() => {
    setTitle("作者");
    api.getAuthors().then((res) => {
      if (res.code === 0) {
        setAuthor(res.data);
      }
    });
  });
  return (
    <>
      <div class="main-box">
        <div class="tag-box">
          <For each={author()}>
            {(item) => (
              <a class="tag" href={"/author/" + item.name}>
                {item.name}
              </a>
            )}
          </For>
        </div>
      </div>
    </>
  );
}
