import { useParams } from "@solidjs/router";
import { createSignal, onMount, Show } from "solid-js";
import api, { Archive } from "../../api/api";
import ArchiveList from "../components/ArchiveList";
import { useTitle } from "../../signal/title";

export default function Author() {
  const params = useParams();
  const [archives, setArchives] = createSignal<Archive[]>([]);
  const [isLoading, setIsLoading] = createSignal(true);
  const [_, setTitle] = useTitle();
  onMount(() => {
    setTitle("作者: " + decodeURIComponent(params.name));
    nextPage(0)
  });

  const [page, setPage] = createSignal(1)
  function nextPage(offset: number) {
    setIsLoading(true)
    const pg = page() + offset
    setPage(pg)
    api
    .getArchiveByAuthorName(params.name, pg)
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
          <button onClick={()=>nextPage(-1)}>prev</button>
          <span>{page()}</span>
          <button onClick={()=>nextPage(1)}>next</button>
        </div>
      </Show>
    </>
  );
}
