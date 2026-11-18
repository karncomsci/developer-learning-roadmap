import "server-only";

import type { components } from "./schema";

export type ApiSchemas = components["schemas"];
export type ApiError = ApiSchemas["Error"];

// The browser never talks to the Go API directly (ADR 0003); this module is
// server-only so API_BASE_URL and, later, the JWT secret stay on the server.
function baseUrl(): string {
  const url = process.env.API_BASE_URL;
  if (!url) throw new Error("API_BASE_URL is not set");
  return url.replace(/\/$/, "");
}

export async function apiGet<T>(path: string): Promise<T> {
  const res = await fetch(`${baseUrl()}${path}`, { cache: "no-store" });
  if (!res.ok) {
    const body = (await res.json().catch(() => null)) as ApiError | null;
    throw new Error(`GET ${path} failed: ${res.status} ${body?.code ?? ""}`.trim());
  }
  return (await res.json()) as T;
}
