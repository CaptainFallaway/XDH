<script lang="ts">
  import type { internal } from "../wailsjs/go/models";

  // Props
  let { scans, unit }: { scans: internal.Scan[], unit: string } = $props();
  
  // Format date from timestamp
  function formatDate(timestamp: number): string {
    return new Date(timestamp).toLocaleDateString();
  }
  
  // Format time from timestamp
  function formatTime(timestamp: number): string {
    return new Date(timestamp).toLocaleTimeString();
  }
</script>

<div class="overflow-x-auto">
  <table class="table table-zebra w-full">
    <thead>
      <tr>
        <th>Reading</th>
        <th>Duration</th>
        <th>Operator</th>
        <th>Date/Time</th>
        <th>Pb ({unit})</th>
        <th>Zn ({unit})</th>
        <th>Cu ({unit})</th>
        <th>Sn ({unit})</th>
      </tr>
    </thead>
    <tbody>
      {#each scans as scan}
        <tr class="opacity-60">
          <td>{scan.reading}</td>
          <td class="text-error">{scan.duration.toFixed(2)}</td>
          <td>{scan.operator}</td>
          <td>
            {formatDate(scan.date)}
            <br />
            <span class="text-xs opacity-70">{formatTime(scan.date)}</span>
          </td>
          <td>{scan.pb.toFixed(2)}</td>
          <td>{scan.zn.toFixed(2)}</td>
          <td>{scan.cu.toFixed(2)}</td>
          <td>{scan.sn.toFixed(2)}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>