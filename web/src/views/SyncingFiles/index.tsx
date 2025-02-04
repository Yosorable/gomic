import { onMount } from "solid-js";
import { createSignal, For } from "solid-js";
import api from "../../api/api";
import { useTitle } from "../../signal/title";

export default function SyncingFiles() {
  const [status, setStatus] = createSignal(false);
  const [records, setRecords] = createSignal<string[]>([]);
  const [msg, setMsg] = createSignal("");
  const [_, setTitle] = useTitle();
  onMount(() => {
    setTitle("同步");
    refreshStatus();
  });

  function refreshStatus() {
    api.scannerStatus().then((res) => {
      if (res.code === 0) {
        setStatus(res.data.status);
        setRecords(res.data.records.reverse());
      }
    });
  }

  function startSyncing() {
    api.startScanner().then((res) => {
      if (res.code === 0) setStatus(true);
      setMsg(res.msg ?? "");
    });
  }

  return (
    <>
      <div
        style={{
          display: "flex",
          gap: "10px",
        }}
      >
        <button
          class="p-1 rounded bg-gray-500 active:bg-gray-400"
          onClick={startSyncing}
        >
          start
        </button>
        <button
          class="p-1 rounded bg-gray-500 active:bg-gray-400"
          onClick={refreshStatus}
        >
          status
        </button>
      </div>

      <div>msg: {msg()}</div>
      <div>status: {status() ? "syncing" : "idle"}</div>
      <div>
        <div>records:</div>
        <For each={records()}>{(item) => <div>{item}</div>}</For>
      </div>
    </>
  );
}
