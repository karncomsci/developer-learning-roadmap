import type { ApiSchemas } from "@/lib/api/client";
import { progressPercent } from "@/lib/progress";

type Role = ApiSchemas["RoleSummary"];

const fallbackColor = "#71717a"; // zinc-500

export function RoleCard({ role }: { role: Role }) {
  const { done, total } = role.progress;
  const percent = progressPercent(done, total);
  const color = role.color ?? fallbackColor;

  return (
    <article className="flex h-full flex-col gap-3 rounded-xl border border-zinc-200 bg-white p-5 shadow-sm dark:border-zinc-800 dark:bg-zinc-950">
      <header className="flex items-center gap-3">
        {role.icon ? (
          <span className="text-2xl" aria-hidden="true">
            {role.icon}
          </span>
        ) : (
          <span className="size-3 shrink-0 rounded-full" style={{ backgroundColor: color }} aria-hidden="true" />
        )}
        <h2 className="text-lg font-semibold">{role.name}</h2>
      </header>

      {role.description && (
        <p className="line-clamp-3 text-sm text-zinc-600 dark:text-zinc-400">{role.description}</p>
      )}

      <div className="mt-auto flex flex-col gap-1.5">
        <div
          className="h-2 w-full overflow-hidden rounded-full bg-zinc-100 dark:bg-zinc-800"
          role="progressbar"
          aria-valuemin={0}
          aria-valuemax={100}
          aria-valuenow={percent}
          aria-label={`${role.name} progress`}
        >
          <div className="h-full rounded-full" style={{ width: `${percent}%`, backgroundColor: color }} />
        </div>
        <p className="text-sm tabular-nums text-zinc-600 dark:text-zinc-400">
          {done}/{total} · {percent}%
        </p>
      </div>
    </article>
  );
}
