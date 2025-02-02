import { createSignal, onMount, Show } from "solid-js";
import api, { Archive } from "../../api/api";
import NavBar from "../components/NavBar";
import ArchiveList from "../components/ArchiveList";
import { Title } from "@solidjs/meta";

export default function Home() {
  const [archives, setArchives] = createSignal<Archive[]>([]);
  const [isLoading, setIsLoading] = createSignal(true);
  onMount(() => {
    api
      .getAllArchives()
      .then((res) => {
        if (res.code === 0) {
          setArchives(res.data);
        }
      })
      .finally(() => {
        setIsLoading(false);
      });
  });
  return (
    <>
      <Title>首页</Title>
      <div class="main-box">
        <NavBar />
        <Show when={!isLoading()}>
          <ArchiveList archives={archives()} />
        </Show>
      </div>
    </>
  );
}
