
export function formatLongDate(date: number): string {
  return new Date(date * 1000).toLocaleString();
}

export function formatShorterDate(date: number): string {
  return new Date(date * 1000).toLocaleDateString();
}



