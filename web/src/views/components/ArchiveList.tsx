import { For } from "solid-js";
import { Archive } from "../../api/api";
import Book from "./Book";

export default function ArchiveList({ archives }: { archives: Archive[] }) {
  return (
    <div class="cover-list">
      <For each={archives}>
        {(item) => (
          <Book
            cover={item.cover_url}
            target={"/viewer/" + item.name}
            title={item.name}
          />
        )}
      </For>
    </div>
  );
}
