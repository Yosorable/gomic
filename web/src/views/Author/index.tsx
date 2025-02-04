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
    api
      .getArchiveByAuthorName(params.name)
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
