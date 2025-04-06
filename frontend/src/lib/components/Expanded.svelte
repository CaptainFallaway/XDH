<script lang="ts">
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

function isViolation(element: string, scan: internal.Scan): boolean {
    console.log(scan)
    return scan.violations[element];
}
</script>

<div {id} class="container mx-auto p-4 bg-base-200 min-h-screen max-w-full">
  <h1 class="text-3xl font-bold mb-4">Enhanced Boat Scan Dashboard</h1>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
    <div class="card bg-primary text-primary-content">
      <div class="card-body">
        <h2 class="card-title">Boat ID</h2>
        <p>{data.boatID}</p>
      </div>
    </div>
    <div class="card bg-secondary text-secondary-content">
      <div class="card-body">
        <h2 class="card-title">Scan Period</h2>
        <p>{formatDate(data.firstDate)} - {formatDate(data.lastDate)}</p>
      </div>
    </div>
    <div class="card bg-accent text-accent-content">
      <div class="card-body">
        <h2 class="card-title">Total Scans</h2>
        <p>{data.scans.length}</p>
      </div>
    </div>
    <div class="card bg-neutral text-neutral-content">
      <div class="card-body">
        <h2 class="card-title">Violations</h2>
        <ul>
          {#each Object.entries(data.violations) as [element, count]}
            <li>{element}: {count}</li>
          {/each}
        </ul>
      </div>
    </div>
  </div>

  <div class="card bg-base-100 shadow-xl mb-4">
    <div class="card-body">
      <h2 class="card-title mb-2">Average Readings</h2>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div class="stat">
          <div class="stat-title">Pb</div>
          <div class="stat-value">{averages.pb}</div>
        </div>
        <div class="stat">
          <div class="stat-title">Zn</div>
          <div class="stat-value">{averages.zn}</div>
        </div>
        <div class="stat">
          <div class="stat-title">Cu</div>
          <div class="stat-value">{averages.cu}</div>
        </div>
        <div class="stat">
          <div class="stat-title">Sn</div>
          <div class="stat-value">{averages.sn}</div>
        </div>
      </div>
    </div>
  </div>

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
