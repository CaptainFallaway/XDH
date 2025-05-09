<script lang="ts">
	import MetalSummary from './MetalSummary.svelte';
import type {
    internal
} from "../wailsjs/go/models";
let {
    id,
    data
}: {
    id: string | null | undefined,
    data: internal.Grouping
} = $props();

function formatDate(timestamp: number) {
    return new Date(timestamp * 1000).toLocaleString();
}

const averages = {
    pb: (data.scans.reduce((sum, scan) => sum + scan.pb, 0) / data.scans.length).toFixed(2),
    sn: (data.scans.reduce((sum, scan) => sum + scan.sn, 0) / data.scans.length).toFixed(2),
    zn: (data.scans.reduce((sum, scan) => sum + scan.zn, 0) / data.scans.length).toFixed(2),
    cu: (data.scans.reduce((sum, scan) => sum + scan.cu, 0) / data.scans.length).toFixed(2),
};

const violations = data.violations;

function isViolation(element: string, scan: internal.Scan): boolean {
    console.log(scan)
    return scan.violations[element];
}
</script>

<div {id} class="container mx-auto py-4 bg-base-200 min-h-screen max-w-full">
  <MetalSummary violationCounts={violations}/>

  {#if data.errorNotes && data.errorNotes.length > 0}
    <div class="alert alert-error shadow-lg mb-4">
      <div class="flex flex-row items-center space-x-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>Error Notes:</span>
      </div>
      <ul>
        {#each data.errorNotes as note}
          <li>{note}</li>
        {/each}
      </ul>
    </div>
  {/if}

  <div class="card bg-base-100 shadow-xl">
    <div class="card-body">
      <h2 class="card-title mb-4">Detailed Scan Data</h2>
      <div class="overflow-x-auto">
        <table class="table table-zebra w-full">
          <thead>
            <tr>
              <th>#</th>
              <th>Reading</th>
              <th>Date</th>
              <th>Duration</th>
              <th>Pb</th>
              <th>Zn</th>
              <th>Cu</th>
              <th>Sn</th>
            </tr>
          </thead>
          <tbody>
            {#each data.scans as scan, i}
              <tr>
                <td>{i + 1}</td>
                <td>{scan.reading}</td>
                <td>{formatDate(scan.date)}</td>
                <td>{scan.duration.toFixed(2)}s</td>
                <td class:text-error={isViolation('pb', scan)}>{scan.pb.toFixed(3)}</td>
                <td class:text-error={isViolation('zn', scan)}>{scan.zn.toFixed(3)}</td>
                <td class:text-error={isViolation('cu', scan)}>{scan.cu.toFixed(3)}</td>
                <td class:text-error={isViolation('sn', scan)}>{scan.sn.toFixed(3)}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>

  <!-- <style lang="postcss">
.text-error {
    @apply text-red-500 font-bold;
}
</style>

-->
