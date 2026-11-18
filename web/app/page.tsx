import { Suspense } from "react";
import { connection } from "next/server";

import { apiGet, type ApiSchemas } from "@/lib/api/client";

async function checkHealth(): Promise<{ ok: true; status: string } | { ok: false; error: string }> {
  try {
    const health = await apiGet<ApiSchemas["Health"]>("/healthz");
    return { ok: true, status: health.status };
  } catch (err) {
    return { ok: false, error: err instanceof Error ? err.message : String(err) };
  }
}

// M0 placeholder: proves the web -> Go API path works end to end.
async function ApiStatus() {
  await connection();
  const result = await checkHealth();
  if (!result.ok) {
    return <p className="text-red-700 dark:text-red-400">API unreachable: {result.error}</p>;
  }
  return <p className="text-green-700 dark:text-green-400">API: {result.status}</p>;
}

export default function Home() {
  return (
    <main className="mx-auto flex w-full max-w-3xl flex-1 flex-col gap-4 px-4 py-16">
      <h1 className="text-3xl font-semibold tracking-tight">Developer Roadmap</h1>
      <Suspense fallback={<p className="text-zinc-500">Checking API…</p>}>
        <ApiStatus />
      </Suspense>
    </main>
  );
}
