import { useParams } from "@solidjs/router";
import { createSignal, onMount, Show } from "solid-js";
import api, { Archive } from "../../api/api";
import NavBar from "../components/NavBar";
import ArchiveList from "../components/ArchiveList";
import { Title } from "@solidjs/meta";

export default function Author() {
  const params = useParams();
  const [archives, setArchives] = createSignal<Archive[]>([]);
  const [isLoading, setIsLoading] = createSignal(true);
  onMount(() => {
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
      <Title>{decodeURIComponent(params.name)}</Title>
      <div class="main-box">
        <NavBar />
        <Show when={!isLoading()}>
          <ArchiveList archives={archives()} />
        </Show>
      </div>
    </>
  );
}
