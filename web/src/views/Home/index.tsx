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
    nextPage(0);
  });

  const [page, setPage] = createSignal(1);
  function nextPage(offset: number) {
    setIsLoading(true);
    const pg = page() + offset;
    setPage(pg);
    api
      .getAllArchives(pg)
      .then((res) => {
        if (res.code === 0) {
          setArchives(res.data);
        }
      })
      .finally(() => {
        setIsLoading(false);
      });
  }

  return (
    <>
      <Show when={!isLoading()}>
        <ArchiveList archives={archives()} />

        <div class="flex gap-1">
          <button onClick={() => nextPage(-1)}>prev</button>
          <span>{page()}</span>
          <button onClick={() => nextPage(1)}>next</button>
        </div>
      </Show>
    </>
  );
}
