<script lang="ts">
  // Props
  let { violationCounts }: {
    violationCounts: Record<string, number>
  } = $props();

  console.log(violationCounts);
  $inspect(violationCounts);
</script>

<div class="stats card bg-base-100 shadow w-full mb-4">
  <div class="stat">
    <div class="stat-title">Lead (Pb)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.pb > 0}>
      {violationCounts.pb > 0 ? violationCounts.pb : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
  
  <div class="stat">
    <div class="stat-title">Zinc (Zn)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.zn > 0}>
      {violationCounts.zn > 0 ? violationCounts.zn : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
  
  <div class="stat">
    <div class="stat-title">Copper (Cu)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.cu > 0}>
      {violationCounts.cu > 0 ? violationCounts.cu : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
  
  <div class="stat">
    <div class="stat-title">Tin (Sn)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.sn > 0}>
      {violationCounts.sn > 0 ? violationCounts.sn : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
</div>