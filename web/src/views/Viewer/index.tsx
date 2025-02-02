import { createSignal, For, onMount } from "solid-js";
import api, { ArchiveFile } from "../../api/api";
import { useParams } from "@solidjs/router";
import { LazyLoad } from "../../utils/lazyload";
import { Title } from "@solidjs/meta";
import { isImage, isVideo } from "../../utils/checkfilename";

export default function Viewer() {
  const [files, setFiles] = createSignal<ArchiveFile[]>([]);
  const params = useParams();
  onMount(() => {
    api.getArchiveByName(params.name).then((res) => {
      if (res.code === 0) {
        setFiles(res.data.files);
      }
    });
    document.body.style.backgroundColor = "black";
  });
  return (
    <>
      <Title>{decodeURIComponent(params.name)}</Title>
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
