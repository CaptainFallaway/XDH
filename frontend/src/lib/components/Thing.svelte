<script lang="ts">
  import type { internal } from "../wailsjs/go/models";
  let { id, grouping }: { id: string | null | undefined, grouping: internal.Grouping } = $props();

  function formatDate(timestamp: number) {
    return new Date(timestamp * 1000).toLocaleString();
  }
</script>

<div class="container mx-auto p-4 bg-base-200 min-h-screen">
  <h1 class="text-3xl font-bold mb-4">Boat Scan Dashboard</h1>
  
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
    <div class="card bg-primary text-primary-content">
      <div class="card-body">
        <h2 class="card-title">Boat ID</h2>
        <p>{grouping.boatID}</p>
      </div>
    </div>
    <div class="card bg-secondary text-secondary-content">
      <div class="card-body">
        <h2 class="card-title">Scan Period</h2>
        <p>{formatDate(grouping.firstDate)} - {formatDate(grouping.lastDate)}</p>
      </div>
    </div>
    <div class="card bg-accent text-accent-content">
      <div class="card-body">
        <h2 class="card-title">Total Scans</h2>
        <p>{grouping.scans.length}</p>
      </div>
    </div>
    <div class="card bg-neutral text-neutral-content">
      <div class="card-body">
        <h2 class="card-title">Violations</h2>
        <ul>
          {#each Object.entries(grouping.violations) as [element, count]}
            <li>{element}: {count}</li>
          {/each}
        </ul>
      </div>
    </div>
  </div>

  <div class="card bg-base-100 shadow-xl">
    <div class="card-body">
      <h2 class="card-title mb-4">Detailed Scan grouping</h2>
      <div class="overflow-x-auto">
        <table class="table table-zebra w-full">
          <thead>
            <tr>
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
            {#each grouping.scans as scan}
              <tr>
                <td>{scan.reading}</td>
                <td>{formatDate(scan.date)}</td>
                <td>{scan.duration.toFixed(2)}s</td>
                <td>{scan.pb.toFixed(3)}</td>
                <td>{scan.zn.toFixed(3)}</td>
                <td>{scan.cu.toFixed(3)}</td>
                <td>{scan.sn.toFixed(3)}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>