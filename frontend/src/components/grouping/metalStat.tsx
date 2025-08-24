function MetalStat({ value }: { value: number }) {
  return (
    <div class="stat">
      <div class="stat-title">Lead (Pb)</div>
      <div class={`${value > 0 ? 'text-error' : ''} stat-value text-2xl`}>{value > 0 ? value : 'OK'}</div>
      <div class="stat-desc">Violations</div>
    </div>
  );
}

export default function MetalSummary({
  violations,
  westCoastFlag,
}: {
  violations: Record<string, number>;
  westCoastFlag: boolean;
}) {
  return (
    <div class="stats card bg-base-100 shadow w-full mb-4">
      {westCoastFlag ? (
        <>
          <MetalStat value={violations.Pb} />
          <MetalStat value={violations.Sn} />
        </>
      ) : (
        <>
          <MetalStat value={violations.Cu} />
          <MetalStat value={violations.Zn} />
        </>
      )}
    </div>
  );
}
