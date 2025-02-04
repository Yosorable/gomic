import { createSignal, For, onMount } from "solid-js";
import api, { ArchiveFile } from "../../api/api";
import { useParams } from "@solidjs/router";
import { LazyLoad } from "../../utils/lazyload";
import { isImage, isVideo } from "../../utils/checkfilename";
import { useTitle } from "../../signal/title";

export default function Viewer() {
  const [files, setFiles] = createSignal<ArchiveFile[]>([]);
  const params = useParams();
  const [_, setTitle] = useTitle();
  onMount(() => {
    setTitle(decodeURIComponent(params.name));
    api.getArchiveByName(params.name).then((res) => {
      if (res.code === 0) {
        setFiles(res.data.files);
      }
    });
  });
  return (
    <>
      <div class="viewer-box">
        <For each={files()}>
          {(item) => {
            const fileName = item.name;
            if (isImage(fileName))
              return (
                <img
                  ref={(el) =>
                    LazyLoad(el, (e) => {
                      const el = e.target as HTMLElement;
                      el.style.height = "auto";
                      el.style.opacity = "1";
                    })
                  }
                  class="lazy-img viewer-img"
                  data-src={item.url}
                />
              );
            else if (isVideo(fileName))
              return (
                <video width="100%" controls>
                  <source src={item.url} />
                </video>
              );
            return <></>;
          }}
        </For>
      </div>
    </>
  );
}
