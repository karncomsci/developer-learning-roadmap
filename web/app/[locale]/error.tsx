"use client";

export default function RolesError({ reset }: { error: Error; reset: () => void }) {
  return (
    <main className="mx-auto w-full max-w-5xl flex-1 px-4 py-10">
      <p className="mb-4 text-red-700 dark:text-red-400">โหลดข้อมูลไม่สำเร็จ</p>
      <button
        type="button"
        onClick={reset}
        className="rounded-md border border-zinc-300 px-3 py-1.5 text-sm hover:bg-zinc-100 dark:border-zinc-700 dark:hover:bg-zinc-900"
      >
        ลองอีกครั้ง
      </button>
    </main>
  );
}
