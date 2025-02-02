const observer = new IntersectionObserver((entries, observer) => {
  entries.forEach((entry) => {
    if (
      !entry.isIntersecting ||
      !(entry.target instanceof HTMLImageElement) ||
      !entry.target.hasAttribute("data-src")
    ) {
      return;
    }
    const ds = entry.target.getAttribute("data-src");
    if (ds) {
      entry.target.src = ds;
    }
    observer.unobserve(entry.target);
  });
});

export function LazyLoad(
  dom: HTMLElement,
  onload?: (this: GlobalEventHandlers, ev: Event) => any
) {
  if (onload) {
    dom.onload = onload;
  } else {
    dom.onload = (e) => {
      e && e.target && ((e.target as HTMLElement).style.opacity = "1");
    };
  }
  observer.observe(dom);
}

export function Unobserve(dom: HTMLElement) {
  observer.unobserve(dom);
}
