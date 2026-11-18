// Role Progress percentage is computed on the display side (P3); 0/0 shows 0% (P2).
export function progressPercent(done: number, total: number): number {
  return total === 0 ? 0 : Math.round((done / total) * 100);
}
