import { createSignal, onMount, Show } from "solid-js";
import api, { Archive } from "../../api/api";
import ArchiveList from "../components/ArchiveList";
import { useTitle } from "../../signal/title";

export default function Home() {
  const [archives, setArchives] = createSignal<Archive[]>([]);
  const [isLoading, setIsLoading] = createSignal(true);
  const [_, setTitle] = useTitle();

  onMount(() => {
    setTitle("首页");
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
      <Show when={!isLoading()}>
        <ArchiveList archives={archives()} />
      </Show>
    </>
  );
}
