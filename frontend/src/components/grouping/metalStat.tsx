function MetalStat({ value }: { value: number }) {
  return (
    <div class="stat">
      <div class="stat-title">Lead (Pb)</div>
      <div class={`${value > 0 ? 'text-error' : ''} stat-value text-2xl`}>{value > 0 ? value : 'OK'}</div>
      <div class="stat-desc">Violations</div>
    </div>
  );
}

export default function MetalSummary({ violations }: { violations: Record<string, number> }) {
  return (
    <div class="stats card bg-base-100 shadow w-full mb-4">
      <MetalStat value={violations.pb} />
      <MetalStat value={violations.zn} />
      <MetalStat value={violations.cu} />
      <MetalStat value={violations.sn} />
    </div>
  );
}
