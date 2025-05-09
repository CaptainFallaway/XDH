<script lang="ts">
  import { flip } from 'svelte/animate';
  // import { createEventDispatcher } from 'svelte';
  import type { internal } from "../wailsjs/go/models";
  
  // Props
  let {
    scans,
    unit,
    enabledScans
  }: {
    scans: internal.Scan[],
    unit: string,
    enabledScans: Record<number, boolean>
  } = $props();
  
  // Format date from timestamp
  function formatDate(timestamp: number): string {
    return new Date(timestamp*1000).toLocaleDateString();
  }
  
  // Format time from timestamp
  function formatTime(timestamp: number): string {
    return new Date(timestamp*1000).toLocaleTimeString();
  }
  
  // Move scan up in the list
  function moveScanUp(index: number): void {
    if (index > 0) {
      const newScans = [...scans];
      [newScans[index - 1], newScans[index]] = [newScans[index], newScans[index - 1]];
    }
  }
  
  // Move scan down in the list
  function moveScanDown(index: number): void {
    if (index < scans.length - 1) {
      const newScans = [...scans];
      [newScans[index], newScans[index + 1]] = [newScans[index + 1], newScans[index]];
    }
  }
  
  // Toggle scan enabled state
  function toggleScan(index: number): void {
    // dispatch('toggleScan', { index });
    enabledScans[index] = !enabledScans[index];
  }
</script>

<div class="overflow-x-auto">
  <table class="table table-zebra w-full">
    <thead>
      <tr>
        <th>Actions</th>
        <th>Reading</th>
        <th>Duration</th>
        <th>Operator</th>
        <th>Date/Time</th>
        <th>Pb ({unit})</th>
        <th>Zn ({unit})</th>
        <th>Cu ({unit})</th>
        <th>Sn ({unit})</th>
        <th>Enabled</th>
      </tr>
    </thead>
    <tbody>
      {#each scans as scan, i (scan.reading)}
        <tr 
          animate:flip={{ duration: 300 }}
          class:opacity-50={!enabledScans[i]}
        >
          <td class="flex gap-1">
            <button 
              class="btn btn-xs" 
              onclick={() => moveScanUp(i)}
              disabled={i === 0}
            >
              ↑
            </button>
            <button 
              class="btn btn-xs" 
              onclick={() => moveScanDown(i)}
              disabled={i === scans.length - 1}
            >
              ↓
            </button>
          </td>
          <td>{scan.reading}</td>
          <td>{scan.duration.toFixed(2)}</td>
          <td>{scan.operator}</td>
          <td>
            {formatDate(scan.date)}
            <br />
            <span class="text-xs opacity-70">{formatTime(scan.date)}</span>
          </td>
          <td class={scan.violations['pb'] ? 'text-error font-bold' : ''}>{scan.pb.toFixed(2)}</td>
          <td class={scan.violations['zn'] ? 'text-error font-bold' : ''}>{scan.zn.toFixed(2)}</td>
          <td class={scan.violations['cu'] ? 'text-error font-bold' : ''}>{scan.cu.toFixed(2)}</td>
          <td class={scan.violations['sn'] ? 'text-error font-bold' : ''}>{scan.sn.toFixed(2)}</td>
          <td>
            <input 
              type="checkbox" 
              class="toggle toggle-primary toggle-sm" 
              checked={enabledScans[i]} 
              onchange={() => toggleScan(i)}
            />
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>