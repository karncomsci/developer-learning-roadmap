import { Suspense } from "react";
import { notFound } from "next/navigation";
import { connection } from "next/server";

import { RoleCard } from "@/components/role-card";
import { apiGet, type ApiSchemas } from "@/lib/api/client";
import { isLocale, locales } from "@/lib/locales";

export function generateStaticParams() {
  return locales.map((locale) => ({ locale }));
}

async function RoleGrid() {
  await connection();
  const roles = await apiGet<ApiSchemas["RoleSummary"][]>("/v1/roles");

  if (roles.length === 0) {
    return <p className="text-zinc-500">ยังไม่มี Role</p>;
  }
  return (
    <ul className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
      {roles.map((role) => (
        <li key={role.id}>
          <RoleCard role={role} />
        </li>
      ))}
    </ul>
  );
}

function RoleGridSkeleton() {
  return (
    <ul className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3" aria-busy="true">
      {[0, 1, 2].map((i) => (
        <li key={i} className="h-40 animate-pulse rounded-xl bg-zinc-100 dark:bg-zinc-900" />
      ))}
    </ul>
  );
}

export default async function RolesPage({ params }: PageProps<"/[locale]">) {
  const { locale } = await params;
  if (!isLocale(locale)) notFound();

  return (
    <main className="mx-auto w-full max-w-5xl flex-1 px-4 py-10">
      <h1 className="mb-8 text-3xl font-semibold tracking-tight">Developer Roadmap</h1>
      <Suspense fallback={<RoleGridSkeleton />}>
        <RoleGrid />
      </Suspense>
    </main>
  );
}
