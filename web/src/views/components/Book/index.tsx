import { A } from "@solidjs/router";
import { LazyLoad } from "../../../utils/lazyload";
import "./index.scss";

export default function Book({
  cover,
  target,
  title,
}: {
  cover: string;
  target: string;
  title: string;
}) {
  return (
    <A class="book-item" href={target} target="_blank">
      <div>
        <img
          class="cover lazy-img"
          ref={(el) => LazyLoad(el)}
          data-src={cover}
        />
      </div>
      <div class="line-clamp-2 overflow-ellipsis p-1">{title}</div>
    </A>
  );
}
