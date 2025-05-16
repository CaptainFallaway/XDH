<script lang="ts">
  // Props
  let { violationCounts }: {
    violationCounts: Record<string, number>
  } = $props();


  $effect(() => {
    console.log(violationCounts);
    console.log('MetalSummary');
  });

  $inspect(violationCounts);
</script>

<div class="stats card bg-base-100 shadow w-full mb-4">
  <div class="stat">
    <div class="stat-title">Lead (Pb)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.Pb > 0}>
      {violationCounts.Pb > 0 ? violationCounts.Pb : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
  
  <div class="stat">
    <div class="stat-title">Zinc (Zn)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.Zn > 0}>
      {violationCounts.Zn > 0 ? violationCounts.Zn : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
  
  <div class="stat">
    <div class="stat-title">Copper (Cu)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.Cu > 0}>
      {violationCounts.Cu > 0 ? violationCounts.Cu : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
  
  <div class="stat">
    <div class="stat-title">Tin (Sn)</div>
    <div class="stat-value text-2xl" class:text-error={violationCounts.Sn > 0}>
      {violationCounts.Sn > 0 ? violationCounts.Sn : 'OK'}
    </div>
    <div class="stat-desc">Violations</div>
  </div>
</div>