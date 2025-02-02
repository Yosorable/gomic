import { For } from "solid-js";
import { Archive } from "../../api/api";
import { A } from "@solidjs/router";
import { LazyLoad } from "../../utils/lazyload";

export default function ArchiveList({ archives }: { archives: Archive[] }) {
  return (
    <div class="cover-list">
      <For each={archives}>
        {(item) => (
          <A href={"/viewer/" + item.name} target="_blank" class="item-box">
            <div>
              <img
                class="cover lazy-img"
                ref={(el) => LazyLoad(el)}
                data-src={item.cover_url}
              />
            </div>
            <div>{item.name}</div>
          </A>
        )}
      </For>
    </div>
  );
}
