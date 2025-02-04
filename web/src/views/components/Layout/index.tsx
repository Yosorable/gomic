import { createSignal, JSX } from "solid-js";
import { Menu, X, User, Home, LibraryBig, RefreshCcw } from "lucide-solid";
import "./index.scss";
import { A } from "@solidjs/router";
import { useTitle } from "../../../signal/title";

export default function Layout({ children }: { children: JSX.Element }) {
  const [isOpen, setIsOpen] = createSignal(false);
  const [title, _] = useTitle();

  return (
    <div class="flex">
      <div
        class={`fixed inset-y-0 left-0 z-50 w-64 bg-gray-800 text-white p-5 
      transform ${isOpen() ? "translate-x-0" : "-translate-x-full"} 
      transition-transform duration-300 ease-in-out lg:translate-x-0`}
      >
        <button
          class="absolute top-4 right-4 text-gray-300 lg:hidden"
          onClick={() => setIsOpen(false)}
        >
          <X size={24} />
        </button>

        <h2 class="text-2xl font-bold mb-6">Gomic</h2>
        <nav>
          <ul class="space-y-4">
            <li>
              <A
                href="/home"
                class="block w-54 p-2 hover:bg-gray-700 [&.active]:bg-gray-700 rounded"
              >
                <Home class="float-left mr-2" /> 首页
              </A>
            </li>
            <li>
              <A
                href="/authors"
                class="block p-2 hover:bg-gray-700 [&.active]:bg-gray-700 rounded"
              >
                <LibraryBig class="float-left mr-2" /> 作者
              </A>
            </li>
            <li>
              <A
                href="/syncing"
                class="block p-2 hover:bg-gray-700 [&.active]:bg-gray-700 rounded"
              >
                <RefreshCcw class="float-left mr-2" /> 同步
              </A>
            </li>
          </ul>
        </nav>
      </div>

      <div class="flex-1 flex flex-col lg:ml-64">
        <header class="fixed top-0 left-0 w-full bg-gray-900 text-white p-4 flex items-center justify-between z-40">
          <button
            onClick={() => setIsOpen(true)}
            class="text-gray-300 lg:hidden"
          >
            <Menu size={24} />
          </button>

          <h1 class="text-lg font-semibold lg:pl-64 line-clamp-1 overflow-ellipsis">
            {title()}
          </h1>

          <div class="w-10 h-10 bg-gray-700 rounded-full flex items-center justify-center">
            <User size={24} class="text-white" />
          </div>
        </header>
        <main class="p-6 mt-20 main-content">{children}</main>
      </div>
    </div>
  );
}
