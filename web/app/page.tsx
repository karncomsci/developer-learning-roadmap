import { redirect } from "next/navigation";

import { defaultLocale } from "@/lib/locales";

// Remembering the last locale in a cookie arrives with i18n in M6.
export default function Root() {
  redirect(`/${defaultLocale}`);
}
