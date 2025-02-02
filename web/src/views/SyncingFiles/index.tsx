import { Title } from "@solidjs/meta";
import NavBar from "../components/NavBar";
import { onMount } from "solid-js";
import { createSignal, For } from "solid-js";
import api from "../../api/api";

export default function SyncingFiles() {
  const [status, setStatus] = createSignal(false);
  const [records, setRecords] = createSignal<string[]>([]);
  const [msg, setMsg] = createSignal("");
  onMount(() => {
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
      <Title>同步文件</Title>
      <div class="main-box">
        <NavBar />
        <div
          style={{
            display: "flex",
            gap: "10px",
          }}
        >
          <button onClick={startSyncing}>start</button>
          <button onClick={refreshStatus}>status</button>
        </div>

        <div>msg: {msg()}</div>
        <div>status: {status() ? "syncing" : "idle"}</div>
        <div>
          <div>records:</div>
          <For each={records()}>{(item) => <div>{item}</div>}</For>
        </div>
      </div>
    </>
  );
}
