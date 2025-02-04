import { createSignal } from "solid-js";

const [title, _setTitle] = createSignal(document.title);
const titleElement = document.querySelector("title");
if (titleElement) {
  const observer = new MutationObserver(function (mutations) {
    mutations.forEach(function (mutation) {
      if (mutation.type === "childList" && mutation.addedNodes.length) {
        _setTitle(document.title);
      }
    });
  });
  const config = { childList: true, subtree: true };
  observer.observe(titleElement, config);
}

function setTitle(t: string) {
  document.title = t;
}

export function useTitle(): [typeof title, typeof setTitle] {
  return [title, setTitle];
}
